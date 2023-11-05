package db

import (
	"bytes"
	"context"
	"html/template"
	"os"

	"healthcare/logger"

	"github.com/jackc/pgx/v4/pgxpool"
)

const default_DATABASE_URL = "postgres://healthcare:healthcare@127.0.0.1:5432/healthcare_service"

var dbPool *pgxpool.Pool

// InitDB - initialize a connection pool to DB.
func InitDB(ctx context.Context) {
	if dbPool != nil {
		return
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		dbUrl = default_DATABASE_URL
	}

	// Connect to DB
	dbConfig, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		logger.Fatalf("error configuring the database:%v", err)
	}
	dbConfig.MinConns = 1
	dbConfig.MaxConns = 10
	dbPool, err = pgxpool.ConnectConfig(ctx, dbConfig)
	if err != nil {
		logger.Errorf("Unable to create connection pool: %v", err)
		os.Exit(1)
	}

	if err := dbPool.Ping(ctx); err != nil {
		logger.Fatalf("%v", err)
	}

	logger.Infof("Connected to DB ...")
}

func Close() {
	if dbPool != nil {
		dbPool.Close()
	}
}

func GetSQLString(templ *template.Template, param any) (*bytes.Buffer, error) {
	sql := new(bytes.Buffer)

	err := templ.Execute(sql, param)
	if err != nil {
		return nil, err
	}

	return sql, nil
}

// DoInsert -
func DoInsert(templ *template.Template, param any) (int64, error) {
	sql, err := GetSQLString(templ, param)
	if err != nil {
		return 0, err
	}

	// Get a db connection from the Pool
	conn, err := dbPool.Acquire(context.Background())
	if err != nil {
		return 0, err
	}
	defer conn.Release()

	logger.Debugf("====sql string==== for template:%s ====\n%v", templ.Name(), sql.String())
	//
	cmd, err := conn.Exec(context.Background(), sql.String())
	if err != nil {
		return 0, err
	}
	cmd.RowsAffected()

	return cmd.RowsAffected(), nil
}

// GetSingleRow -
func GetSingleRow(templ *template.Template, param any, dest ...interface{}) error {
	sql, err := GetSQLString(templ, param)
	if err != nil {
		return err
	}

	// Get a db connection from the Pool
	conn, err := dbPool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()

	logger.Debugf("====sql string==== for template:%s ====\n%v", templ.Name(), sql.String())
	//
	err = conn.QueryRow(context.Background(), sql.String()).Scan(dest...)
	if err != nil {
		return err
	}

	return nil
}

// GetRows -
func GetRows(templ *template.Template, param *map[string]interface{}) (*[]map[string]interface{}, error) {
	sql, err := GetSQLString(templ, param)
	if err != nil {
		return nil, err
	}

	// Get a db connection from the Pool
	conn, err := dbPool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	logger.Debugf("====sql string==== for template:%s ====\n%v", templ.Name(), sql.String())
	//
	rows, err := conn.Query(context.Background(), sql.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns := rows.FieldDescriptions()
	var results []map[string]interface{}

	// get the sensor details from db
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			logger.Errorf(err.Error())
			continue
		}

		row := make(map[string]interface{})
		for i, v := range values {
			row[string(columns[i].Name)] = v
		}

		results = append(results, row)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	logger.Infof("results:%v", results)

	return &results, nil
}

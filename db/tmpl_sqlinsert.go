package db

const CreateSensorSQLText = `
	INSERT INTO  sensors (group_name, code_name,idx, x_3d, y_3d, z_3d, output_rate_sec) 
	VALUES ('{{.GROUP_NAME}}', '{{.CODE_NAME}}', '{{.IDX}}',{{.X_3D}}, {{.Y_3D}}, {{.Z_3D}}, {{.RATE}})
`

const SensorIdMaxSQLText = `
	SELECT
		max(idx)
	FROM
		sensors s
	WHERE
		s.group_name = '{{.GROUP_NAME}}'
`

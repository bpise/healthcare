package db

const TransparencyAverageSQLText = `
	SELECT
		AVG(f.transparency)
	FROM
		sensors s
		left join fish_specie_data f on s.id = f.sensor_id
	where
		s.group_name = '{{.GROUP_NAME}}'
`

const TemperatureAverageSQLText = `
	SELECT
		AVG(f.temperature)
	FROM
		sensors s
		left join fish_specie_data f on s.id = f.sensor_id
	where
		s.group_name = '{{.GROUP_NAME}}'
`

const FishSpeciesSQLText = `
	SELECT
		f.fish_specie_name, sum(f.fish_specie_count) as total
	FROM
		sensors s
		left join fish_specie_data f on s.id = f.sensor_id
	where
		s.group_name = '{{.GROUP_NAME}}'
		group by f.fish_specie_name
`

const FishSpeciesTopNSQLText = `
	SELECT
		f.fish_specie_name, sum(f.fish_specie_count) as total
	FROM
		sensors s
		left join fish_specie_data f on s.id = f.sensor_id
	where
		s.group_name = '{{.GROUP_NAME}}'
		{{if .isValidFromTill}}
		and f.created_time BETWEEN to_timestamp({{.FROM}}) and to_timestamp({{.TILL}})
		{{end}}
	group by f.fish_specie_name
	order by total desc
	limit {{.TOP_N}} offset 0
`

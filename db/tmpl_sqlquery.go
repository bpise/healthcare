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

package db

// Calculate the average transparency in a sensor group.
const TransparencyAverageSQLText = `
	SELECT
		AVG(f.transparency)
	FROM
		sensors s
		left join fish_specie_data f on s.id = f.sensor_id
	where
		s.group_name = '{{.GROUP_NAME}}'
`

// Calculate the average temperature in a sensor group.
const TemperatureAverageSQLText = `
	SELECT
		AVG(f.temperature)
	FROM
		sensors s
		left join fish_specie_data f on s.id = f.sensor_id
	where
		s.group_name = '{{.GROUP_NAME}}'
`

// Retrieve a list of fish species with their total counts in a sensor group.
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

// Retrieve the top N fish species with their total counts in a sensor group.
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

// Find the minimum temperature in a specified 3D region.
const TemperatureInRegionMinSQLText = `
	SELECT
		min(f.temperature)
	FROM
		sensors s
		left join fish_specie_data f on s.id = f.sensor_id
	WHERE
		s.x_3d BETWEEN {{.xMin}} and {{.xMax}}
		and s.y_3d BETWEEN {{.yMin}} and {{.yMax}}
		and s.z_3d BETWEEN {{.zMin}} and {{.zMax}}
`

// Find the maximum temperature in a specified 3D region.
const TemperatureInRegionMaxSQLText = `
	SELECT
		max(f.temperature)
	FROM
		sensors s
		left join fish_specie_data f on s.id = f.sensor_id
	WHERE
		s.x_3d BETWEEN {{.xMin}} and {{.xMax}}
		and s.y_3d BETWEEN {{.yMin}} and {{.yMax}}
		and s.z_3d BETWEEN {{.zMin}} and {{.zMax}}
`

// Calculate the average temperature detected by a specific sensor within a specified time range.
const TemperatureAverageBySensorSQLText = `
	SELECT
		avg(f.temperature)
	FROM
		sensors s
		left join fish_specie_data f on s.id = f.sensor_id
	WHERE
		s.code_name = '{{.CODE_NAME}}'
		and f.created_time BETWEEN to_timestamp({{.FROM}}) and to_timestamp({{.TILL}})
`

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

const CreateFishSpecieDataSQLText = `
	INSERT INTO  fish_specie_data (fish_specie_name, fish_specie_count, temperature, transparency, sensor_id) 
	VALUES ('{{.NAME}}', {{.COUNT}}, {{.TEMP}}, {{.TRAN}}, '{{.ID}}')
`

const ActivatedSensorsSQLText = `
	SELECT
		s.id::Text, s.group_name, s.code_name, s.idx, s.x_3d, s.y_3d, s.z_3d, s.output_rate_sec
	FROM
		sensors s
	WHERE
		s.deactivation_time is NULL 
`

const NearbySensorTransparencySQLText = `
	select 
		f.transparency, (|/({{.X_3D}}- s.x_3d)^2 + ({{.Y_3D}}-s.y_3d)^2 + ({{.Z_3D}}-s.z_3d)^2) as dist
	from sensors s
			left join fish_specie_data f on s.id = f.sensor_id
	where s.id != '{{.ID}}'
	order by dist asc
	limit 1 offset 0
`

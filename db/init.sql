CREATE ROLE healthcare WITH ENCRYPTED PASSWORD 'healthcare' LOGIN;

REVOKE ALL ON SCHEMA public FROM healthcare;

CREATE DATABASE healthcare_service;

\c healthcare_service;
REVOKE ALL ON DATABASE healthcare_service FROM public;

GRANT ALL ON DATABASE healthcare_service TO healthcare;

-- Extension to use gen_random_uuid()
CREATE EXTENSION pgcrypto;

SET ROLE healthcare;

CREATE TABLE IF NOT EXISTS sensors(
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  group_name text NOT NULL,
  code_name text NOT NULL,
  idx int NOT NULL,
  x_3d bigint,
  y_3d bigint,
  z_3d bigint,
  output_rate_sec int,
  activation_time timestamp DEFAULT NOW() NOT NULL,
  deactivation_time timestamp
);

CREATE UNIQUE INDEX sensors_uidx1 ON sensors(group_name, idx, code_name)
WHERE
  activation_time IS NOT NULL AND deactivation_time IS NULL;

GRANT SELECT, INSERT, UPDATE, DELETE ON sensors TO healthcare;

CREATE TABLE IF NOT EXISTS fish_specie_data(
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  fish_specie_name text NOT NULL,
  fish_specie_count int,
  temperature real,
  transparency int,
  sensor_id uuid REFERENCES sensors(id),
  created_time timestamp DEFAULT NOW() NOT NULL,
  deleted_time timestamp
);

CREATE UNIQUE INDEX fish_specie_data_uidx1 ON fish_specie_data(fish_specie_name, fish_specie_count, created_time)
WHERE
  created_time IS NOT NULL;

GRANT SELECT, INSERT, UPDATE, DELETE ON fish_specie_data TO healthcare;


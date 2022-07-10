/* just a template
CREATE TABLE hello (
	id SERIAL PRIMARY KEY,
	msg VARCHAR(250) NOT NULL,
	temperature DOUBLE PRECISION NOT NULL,
	ts TIMESTAMP NOT NULL,
	tstz TIMESTAMPTZ NOT NULL
	);
*/

DROP TABLE IF EXISTS devicestatus;
CREATE TABLE devicestatus (
    rowid SERIAL PRIMARY KEY,
	device_id VARCHAR(250) NOT NULL,
	created_on TIMESTAMPTZ NOT NULL,
	status INTEGER,
	status_desc VARCHAR
	);

DROP TABLE IF EXISTS metrics;
CREATE TABLE metrics (
    rowid SERIAL PRIMARY KEY,
	device_id VARCHAR(250),
	created_on TIMESTAMPTZ NOT NULL,
	temperature DOUBLE PRECISION,
	humidity DOUBLE PRECISION,
	co2 DOUBLE PRECISION,
	voc DOUBLE PRECISION,
	pm25 DOUBLE PRECISION,
	lighting DOUBLE PRECISION,
	noise DOUBLE PRECISION
	);

/* no FK to rawmetrics yet */
DROP TABLE IF EXISTS metricscores;
CREATE TABLE metricscores (
    rowid SERIAL PRIMARY KEY,
	device_id VARCHAR(250),
	created_on TIMESTAMPTZ NOT NULL,
	temperature DOUBLE PRECISION,
	humidity DOUBLE PRECISION,
	co2 DOUBLE PRECISION,
	voc DOUBLE PRECISION,
	pm25 DOUBLE PRECISION,
	lighting DOUBLE PRECISION,
	noise DOUBLE PRECISION
	);

/* no FK to metricscores yet */
DROP TABLE IF EXISTS ieqscores;
CREATE TABLE ieqscores (
    rowid SERIAL PRIMARY KEY,
	device_id VARCHAR(250),
	created_on TIMESTAMPTZ NOT NULL,
	scheme VARCHAR(250),
	thermal DOUBLE PRECISION,
	iaq DOUBLE PRECISION,
	lighting DOUBLE PRECISION,
	noise DOUBLE PRECISION,
	overall DOUBLE PRECISION,
    thermal_weighting DOUBLE PRECISION,
    iaq_weighting DOUBLE PRECISION,
    lighting_weighting DOUBLE PRECISION,
    noise_weighting DOUBLE PRECISION
    );

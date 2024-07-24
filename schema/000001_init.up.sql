CREATE TABLE anomalies
(
    id serial not null unique,
    session_id varchar(255) not null,
    frequency  float,
    tmstp      timestamp with time zone
);
CREATE DATABASE IF NOT EXISTS exchanger;

CREATE TABLE IF NOT EXISTS exchanger.course
(
    exchanger String,
    currency String,
    course Float64,
    time timestamp
)
ENGINE = MergeTree()
ORDER BY (time);

--BASIC

CREATE DATABASE IF NOT EXISTS exchanger;
CREATE TABLE IF NOT EXISTS exchanger.course
(
    exchanger UInt8,
    exchangers_condition_id UInt16,
    course Float64,
    time timestamp
)
ENGINE = MergeTree()
ORDER BY (time);


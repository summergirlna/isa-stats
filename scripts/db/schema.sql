CREATE TABLE border_passenger_counts
(
    id           SERIAL PRIMARY KEY,
    direction CHAR(1) NOT NULL,
    year         INTEGER     NOT NULL,
    month        INTEGER     NOT NULL,
    airport_name VARCHAR(30) NOT NULL,
    japan_num    INTEGER     NOT NULL,
    foreign_num  INTEGER     NOT NULL,
    treaty_num   INTEGER     NOT NULL,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (direction, year, month, airport_name)
);

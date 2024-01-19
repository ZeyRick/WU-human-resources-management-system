ALTER TABLE
    schedules DROP COLUMN hour_work;

ALTER TABLE
    schedules
ADD
    COLUMN clock_in_time TIMESTAMP NOT NULL;

ALTER TABLE
    schedules
ADD
    COLUMN clock_out_time TIMESTAMP NOT NULL;
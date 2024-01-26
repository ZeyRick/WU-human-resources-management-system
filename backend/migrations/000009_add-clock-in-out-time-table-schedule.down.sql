ALTER TABLE
    schedules
ADD
    COLUMN hour_work INT NOT NULL;

ALTER TABLE
    schedules DROP COLUMN clock_in_time;

ALTER TABLE
    schedules DROP COLUMN clock_out_time;
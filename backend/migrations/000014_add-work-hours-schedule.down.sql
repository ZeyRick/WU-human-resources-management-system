ALTER TABLE
    schedules
ADD
    COLUMN minute_work_per_day INT NOT NULL,
ADD
    COLUMN minute_break_per_day INT NOT NULL,
ADD
    COLUMN minute_allow_time INT DEFAULT 0 NOT NULL;

ALTER TABLE
    clocks RENAME COLUMN clock_out_hour TO clock_out_minute;
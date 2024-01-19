ALTER TABLE
    clocks
ADD
COLUMN clock_in_id INT REFERENCES clocks;

ALTER TABLE
    clocks
ADD
COLUMN clock_out_hour INT;
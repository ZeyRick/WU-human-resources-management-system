ALTER TABLE
    clocks
ADD
    COLUMN check_in_id INT REFERENCES clocks;

ALTER TABLE
    clocks
ADD
    COLUMN check_out_hour INT NOT NULL;
ALTER TABLE
    clocks
ADD
    COLUMN schedule_id INT,
ADD
    CONSTRAINT FOREIGN KEY (schedule_id) REFERENCES schedules(id) ON DELETE RESTRICT;
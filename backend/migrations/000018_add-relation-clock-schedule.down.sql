ALTER TABLE
    clocks DROP FOREIGN KEY IF EXISTS fk_schedule_id,
    DROP COLUMN IF EXISTS schedule_id;
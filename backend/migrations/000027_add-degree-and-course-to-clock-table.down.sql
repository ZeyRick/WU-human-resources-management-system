ALTER TABLE
    clocks DROP FOREIGN KEY course;

ALTER TABLE
    clocks DROP FOREIGN KEY degree;

ALTER TABLE
    clocks DROP COLUMN IF EXISTS course;

ALTER TABLE
    clocks DROP COLUMN IF EXISTS degree;

ALTER TABLE
    clocks DROP COLUMN IF EXISTS clockTime;
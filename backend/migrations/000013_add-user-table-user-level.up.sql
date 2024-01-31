ALTER TABLE
    employees DROP COLUMN bind_status;

ALTER TABLE
    employees DROP COLUMN profile_pic;

ALTER TABLE
    users
ADD
    COLUMN user_level INT NOT NULL;

UPDATE
    users
SET
    user_level = 99
WHERE
    username = 'root';
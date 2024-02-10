ALTER TABLE
    employees DROP COLUMN bind_status;

ALTER TABLE
    employees DROP COLUMN profile_pic;

ALTER TABLE
    users
ADD
COLUMN user_level VARCHAR(25) NOT NULL;

UPDATE
    users
SET
user_level = 'root'
WHERE
    username = 'root';
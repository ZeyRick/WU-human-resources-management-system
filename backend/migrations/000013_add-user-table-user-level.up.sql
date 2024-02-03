ALTER TABLE
    employees DROP COLUMN bind_status;

ALTER TABLE
    employees DROP COLUMN profile_pic;

ALTER TABLE
    users
ADD
COLUMN user_level VARCHAR(25) NOT NULL;

ALTER TABLE
    users
ADD
    COLUMN profile_pic VARCHAR(255);

UPDATE
    users
SET
user_level = 'root'
WHERE
    username = 'root';
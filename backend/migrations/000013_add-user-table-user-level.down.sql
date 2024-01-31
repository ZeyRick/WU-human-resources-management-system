ALTER TABLE
    employees
ADD
    COLUMN bind_status ENUM('pending', 'approved');

ALTER TABLE
    employees
ADD
    COLUMN profile_pic TEXT(255);

ALTER TABLE
    users DROP COLUMN profile_pic;

ALTER TABLE
    users DROP COLUMN user_level;
-- Step 1: Drop the new tables
DROP TABLE IF EXISTS employee_courses;

DROP TABLE IF EXISTS employee_degrees;

DROP TABLE IF EXISTS degrees;

DROP TABLE IF EXISTS courses;

-- Step 2: Recreate the courses table
CREATE TABLE courses (
    id INT NOT NULL AUTO_INCREMENT,
    alias VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY fk_alias (alias),
    PRIMARY KEY(id)
);

-- Step 3: Add the course_id column back to the employees table
ALTER TABLE
    employees
ADD
    COLUMN course_id INT;
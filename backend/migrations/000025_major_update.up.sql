CREATE TABLE courses (
    id INT NOT NULL AUTO_INCREMENT,
    alias VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY fk_alias (alias),
    PRIMARY KEY(id)
);

ALTER TABLE
    employees DROP COLUMN department_id;

DROP TABLE IF EXISTS departments;

CREATE TABLE degrees (
    id INT NOT NULL AUTO_INCREMENT,
    alias VARCHAR(255) NOT NULL,
    rate DOUBLE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY fk_alias (alias),
    PRIMARY KEY(id)
);

CREATE TABLE employee_degrees (
    employee_id INT,
    degree_id INT,
    PRIMARY KEY (employee_id, degree_id),
    FOREIGN KEY (employee_id) REFERENCES employees(id) ON DELETE CASCADE,
    FOREIGN KEY (degree_id) REFERENCES degrees(id) ON DELETE CASCADE
);

CREATE TABLE employee_courses (
    employee_id INT,
    course_id INT,
    PRIMARY KEY (employee_id, course_id),
    FOREIGN KEY (employee_id) REFERENCES employees(id) ON DELETE CASCADE,
    FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE CASCADE
);
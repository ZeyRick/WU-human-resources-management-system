ALTER TABLE
    employees
ADD
    COLUMN employee_type ENUM('Fulltime', 'Parttime'),
ADD
    COLUMN salary DOUBLE;
ALTER TABLE
    employee_requests
MODIFY
    COLUMN telegram_id INT NOT NULL;

ALTER TABLE
    employees
MODIFY
    COLUMN telegram_id INT NULL;
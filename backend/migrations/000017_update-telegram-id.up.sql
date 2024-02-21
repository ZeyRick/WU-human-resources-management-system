ALTER TABLE
    employee_requests
MODIFY
    COLUMN telegram_id BIGINT NOT NULL;

ALTER TABLE
    employees
MODIFY
    COLUMN telegram_id BIGINT NULL;
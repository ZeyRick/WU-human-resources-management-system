CREATE TABLE employee_requests (
    id INT NOT NULL AUTO_INCREMENT,
    employee_id INT NOT NULL,
    telegram_id INT NOT NULL,
    telegram_username VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (employee_id) REFERENCES employees(id),
    PRIMARY KEY(id),
    UNIQUE (telegram_id)
);
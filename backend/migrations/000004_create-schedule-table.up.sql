CREATE TABLE schedules (
    id INT NOT NULL AUTO_INCREMENT,
    employeeId INT NOT NULL,
    scope VARCHAR(7) NOT NULL,
    dates VARCHAR(80) DEFAULT '[]',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY employee_id_scope (employeeId, scope),
    FOREIGN KEY (employeeId) REFERENCES employees(id),
    PRIMARY KEY(id)
);
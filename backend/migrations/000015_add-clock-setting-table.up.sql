CREATE TABLE clock_settings (
    coordinate varchar(255) NOT NULL DEFAULT '11.571667,104.889259',
    clock_range INT NOT NULL DEFAULT 10,
    allow_time INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
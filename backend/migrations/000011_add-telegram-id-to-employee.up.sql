ALTER TABLE
    employees
ADD
    telegram_id varchar(255);

ALTER TABLE
    employees
ADD
    telegram_username varchar(255);

ALTER TABLE
    employees
ADD
    bind_status ENUM('pending', 'approved');
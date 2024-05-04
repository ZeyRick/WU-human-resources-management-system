ALTER TABLE
    employees
ADD
    `not_deleted` tinyint AS (if(`deleted_at` is null, 1, NULL)) STORED;

ALTER TABLE
    employees DROP INDEX `name`;

CREATE UNIQUE INDEX `IDX_name_not_deleted` ON employees (`name`, `not_deleted`);
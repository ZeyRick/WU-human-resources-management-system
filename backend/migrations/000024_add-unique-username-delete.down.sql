ALTER TABLE
    employees DROP INDEX `IDX_name_not_deleted`;

ALTER TABLE
    employees DROP COLUMN `not_deleted`;

ALTER TABLE
    employees
ADD
    UNIQUE INDEX `name` (`name`);
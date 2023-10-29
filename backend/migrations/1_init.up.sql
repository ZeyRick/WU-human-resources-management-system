CREATE TABLE Users (
    id int NOT NULL AUTO_INCREMENT,
    username varchar(255) NOT NULL,
    name varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    profile_pic TEXT(255),
    PRIMARY KEY (id)
);
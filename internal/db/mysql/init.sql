USE hitss;
-- --------------- --
-- --- TABLES  --- --
-- --------------- --
CREATE TABLE Users (
    id        INT(6)       NOT NULL AUTO_INCREMENT,
    name      VARCHAR(50)  NOT NULL,
    email     VARCHAR(80)  NOT NULL,
    password  VARCHAR(100) NOT NULL,
    tag       VARCHAR(50)  NOT NULL,
    active    TINYINT(1),
    PRIMARY KEY (id)
);
-- --------------- --
-- STORE PROCEDURE --
-- --------------- --

-- LOGIN
-- -----------------------------------------
DELIMITER //
CREATE PROCEDURE spUserLogin(IN _email VARCHAR(80))
BEGIN
 SELECT id, name, password FROM Users WHERE email = _email AND active = 1;
END 
//
-- GET All
-- -----------------------------------------
DELIMITER //
CREATE PROCEDURE spUsersGetAll()
BEGIN
 SELECT id, name, email, password, tag, active FROM Users;
END 
//
-- CREATE
-- -----------------------------------------
DELIMITER //
CREATE PROCEDURE spUserCreate(
 IN _name      VARCHAR(50),
 IN _email     VARCHAR(80),
 IN _password  VARCHAR(100),
 IN _tag       VARCHAR(50),
 IN _active    TINYINT(1)
)
BEGIN
   INSERT INTO Users(name, email, password, tag, active)VALUES(_name, _email, _password, _tag, _active);
   SELECT LAST_INSERT_ID();
END 
//
-- UPDATE
-- -----------------------------------------
DELIMITER //
CREATE PROCEDURE spUserUpdate(
 IN _id        INT(6),
 IN _name      VARCHAR(50),
 IN _email     VARCHAR(80),
 IN _password  VARCHAR(100),
 IN _tag       VARCHAR(50),
 IN _active    TINYINT(1)
)
BEGIN
   UPDATE Users SET name=_name, email=_email, password=_password, tag=_tag, active=_active WHERE id=_id;
END 
//
-- DELETE
-- -----------------------------------------
DELIMITER //
CREATE PROCEDURE spUserDelete(IN _id INT(6))
BEGIN
   DELETE FROM Users WHERE id=_id;
END 
//
-- --------------- --
-- ---- CALLS ---- --
-- --------------- --
CALL spUserCreate('Jorge Luis Alonso','alonso12@gmail.com','$2a$10$8vSiub/rQTBwrqcwju3.V.qGjAPGzpeAx4Ui2irjYuJLP83ZK2kF2','Alonso12',1);

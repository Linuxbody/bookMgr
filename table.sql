--
-- bookMgs Sql
--
CREATE TABLE book(
    id bigint(100) AUTO_INCREMENT PRIMARY KEY,
    title varchar(50) NOT NULL,
    price double(16,2) NOT NULL
)engine=InnoDB default charset=utf8mb4;
# Online Shopping

## Data Bases

- hotwords

  ```MYSQL
  CREATE TABLE hotwords (
  	`id` INT(10) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  	`word` varchar(30) NOT NULL UNIQUE,
      `num` INT(6) NOT NULL DEFAULT 1
  ) CHARACTER SET = utf8
  ```
  

  
- Users

  ```mysql
  CREATE TABLE `users` (
  	`id` INT(10) NOT NULL AUTO_INCREMENT PRIMARY KEY,
      `username` varchar(64) NOT NULL,
      `nickname` varchar(64) NOT NULL,
      `password` varchar(64) NOT NULL,
      `mobile` varchar(15) NOT NULL,
      `email` varchar(128) NOT NULL
  ) CHARACTER SET = utf8
  ```

- Commodities

  ```mysql
  CREATE TABLE `commodities` (
  	`id` INT(10) NOT NULL AUTO_INCREMENT PRIMARY KEY,
      `title` varchar(64) NOT NULL,
      `desc` TEXT NOT NULL,
      `category` int(2) NOT NULL,
      `price` INT(10) NOT NULL,
      `picture` varchar(256) NOT NULL,
      `pub_user` INT(10) NOT NULL,
      `view_count` INT(10) NOT NULL DEFAULT 0,
      `collect_count` INT(10) NOT NULL DEFAULT 0,
      CONSTRAINT fk_pub_user_id FOREIGN KEY (`pub_user`) REFERENCES users(id)
  ) CHARACTER SET = utf8
  ```

- collections

  ```MYSQL
  CREATE TABLE collections (
  	id INT(10) AUTO_INCREMENT PRIAMRY KEY,
      user_id INT(10) NOT NULL,
      commodity_id INT(10) NOT NULL,
      CONSTRAINT fk_user_id FOREIGN KEY (`user_id`) REFERENCES users(id),
      CONSTRAINT fk_commodity_id FOREIGN KEY (`commodity_id`) REFERENCES commodity(id)
  ) CHARACTER SET = utf8
  ```
  
  
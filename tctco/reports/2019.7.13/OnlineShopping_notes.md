# Online Shopping

## Data Bases

- Tags

  ```SQL
  CREATEM TABLE `tags` (
  	`id` INT(5) NOT NULL AUTO_INCREMENT PRIMARY KEY,
      `tagname` varchar(10) NOT NULL,
      PRIMARY KEY(id),
      FOREIGN KEY(`good_id`) REFERENCES goods(id),
  )
  ```

  

- Users

  ```sql
  CREATE TABLE `users` (
  	`id` INT(10) NOT NULL AUTO_INCREMENT PRIMARY KEY,
      `username` varchar(64) NOT NULL,
      `password` varchar(64) NOT NULL,
      `telephone` varchar(15) NOT NULL,
  )
  ```

- Goods

  ```SQL
  CREATE TABLE `goods` (
  	`id` INT(10) NOT NULL AUTO_INCREMENT PRIMARY KEY,
      `goodname` varchar(64) NOT NULL,
      `price` INT(8) NOT NULL,
      `pic_path` varchar(128) NOT NULL,
      `seller_id` INT(10) NOT NULL,
      CONSTRAINT fk_ FOREIGN KEY(`seller_id`) REFERENCES users(id),
  )
  ```

- buyer2good ?

  ```MYSQL
  CREATE TABLE buyer2good (
  	id INT(10) AUTO_INCREMENT PRIAMRY KEY,
      buyer_id INT(10),
      FOERIGN KEY(buyer_id) REFERENCES user(id),
      ON DELETE CASCADE ON UPDATE CASCADE,
      good_id INT(10),
      FOREIGN KEY(good_id) REFERENCES good(id),
      ON DELETE CASCADE ON UPDATE CASCADE,
  )
  ```

  
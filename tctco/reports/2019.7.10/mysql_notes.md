# MySQL

## Basic commands

1. Create, delete and basic query:

   ```mysql
   show databases;
   create database db;
   drop database db;
   use db;
   show tables;
   create table mytable(name varchar(20), sex char(1), birth date);
   drop table mytable;
   show tables;
   describe mytable;
   ```

2. manipulating columns:

   ```mysql
   alter table yourtable add name varchar(20) not null;
   alter table yourtable drop name;
   select name from mytable;
   select name, birth from mytable;
   ```

3. manipulating rows:

   ```mysql
   insert into mytable values('summer', 'm', '1983-08-24');
   delete from mytable where name='summer';
   update mytable set sex='vm' where name='summer';
   insert into mytable select *from yourtable;
   ```

4. show column name in the result:

   ```mysql
   select name as '姓名' from students order by age;
   select name '姓名' from students order by age;
   ```

5. precise lookup:

   ```mysql
   select * from students where native in ('湖南'， '四川');
   select * from students where age between 20 and 30;
   select * from students where name = '李山';
   select * from students where name like '李%'; ('%李%'，'_李','_李_'， RE)
   ```

6. ```mysql
   select count(*) from students;(总人数)
   select avg(mark) from grades where cno='B2';
   min(col) and min(col)
   ```




### My database

```mysql
create table `userinfo` (
	`id` int(10) not null auto_increment,
    `username` varchar(64) not null,
    `password` varchar(128) not null,
    `phonenumber` varchar(15) not null,
    `email` varchar(128) not null,
    `authority` int(1),
    primary key(`id`),
    unique key(`username`)
);
```


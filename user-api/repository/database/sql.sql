sudo su -

service mysql start

mysql -u root -p 

create database store;

use store;

create table user(
  id varchar(150) not null,
  name varchar(50) not null,
  cpf varchar(15) not null unique,
  email varchar(50) not null unique,
  phone_number varchar(20) not null,
  created_at timestamp default current_timestamp(),
  updated_at timestamp default current_timestamp() on update current_timestamp()
)ENGINE=INNODB;


create user 'adm'@'localhost' identified by 'Pass123!';

grant all privileges on store.* to 'adm'@'localhost';

exit;


mysql -u adm -p → senha: Pass123!

show databases;

use store;

show tables;

exit;

//ENGINE=INNODB é adequado para aplicativos que exigem suporte para transações e integridade de dados.

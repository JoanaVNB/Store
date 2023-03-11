create table `order`(
  id varchar(150) not null,
  user_id varchar(99) not null,
  item_description varchar(20) not null,
  item_quantity int not null,
  item_price int not null,
  total_value int not null,
  created_at timestamp default current_timestamp(),
  updated_at timestamp default current_timestamp() on update current_timestamp()
)ENGINE=INNODB;
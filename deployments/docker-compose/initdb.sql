create database appdb;
use appdb;

create table product (
    id int not null auto_increment,
    name varchar(100) not null,
    description varchar(255) not null,
    price decimal(10,2) not null,
    quantity int not null,
    created_at timestamp default current_timestamp,
    primary key (id)
);
create database appdb;
use appdb;

create table product (
    id int not null auto_increment,
    name varchar(255) not null,
    price decimal(10,2) not null,
    quantity int not null,
    primary key (id)
);
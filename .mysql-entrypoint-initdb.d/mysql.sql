create database if not exists fydb;
USE mysql;
update user set host = '%' where user = 'root';
flush privileges;

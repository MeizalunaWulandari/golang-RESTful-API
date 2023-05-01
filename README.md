# Golang RESTful API

## Aplikasi CRUD Sederhana
## Depedensi
Driver MySQL  : https://github.com/go-sql-driver/mysql<br>
HTTP Router: https://github.com/julienschmidt/httprouter<br>
Validation: https://github.com/go-playground/validator<br>

## Membuat Open API

## Membuat Database
```SQL
	CREATE DATABASE golang_restful_api;
	CREATE TABLE category(
			id integer primary key auto_increment,
			name varchar(200) not null
		)engine = InnoDB;
```
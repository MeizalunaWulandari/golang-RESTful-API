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

## Ragam Error

```
T2023/05/05 12:50:15 http: panic serving 127.0.0.1:55530: dial tcp: lookup locahost on 127.0.0.53:53: server misbehaving
goroutine 8 [running]:
net/http.(*conn).serve.func1()
	/usr/lib/go-1.18/src/net/http/server.go:1825 +0xbf
panic({0x6c7480, 0xc0002ba140})
	/usr/lib/go-1.18/src/runtime/panic.go:844 +0x258
golang-resful-api/helper.PanicIfError(...)
	/home/bismillah/GoLang/src/golang-resful-api/helper/error.go:5
golang-resful-api/service.(*CategoryServiceImpl).FindAll(0xc000027f80, {0x7801b0, 0xc000214100})
	/home/bismillah/GoLang/src/golang-resful-api/service/category_service_impl.go:84 +0x23c
golang-resful-api/controller.(*CategoryControllerImpl).FindAll(0x0?, {0x780000, 0xc0002b0000}, 0x40?, {0xc00007ef00?, 0x0?, 0xc00021c004?})
	/home/bismillah/GoLang/src/golang-resful-api/controller/category_controller_impl.go:86 +0x5d
github.com/julienschmidt/httprouter.(*Router).ServeHTTP(0xc000074300, {0x780000, 0xc0002b0000}, 0xc00021a000)
	/home/bismillah/GoLang/pkg/mod/github.com/julienschmidt/httprouter@v1.3.0/router.go:387 +0x82b
net/http.serverHandler.ServeHTTP({0xc000212030?}, {0x780000, 0xc0002b0000}, 0xc00021a000)
	/usr/lib/go-1.18/src/net/http/server.go:2916 +0x43b
net/http.(*conn).serve(0xc00024a280, {0x780258, 0xc0001f6db0})
	/usr/lib/go-1.18/src/net/http/server.go:1966 +0x5d7
created by net/http.(*Server).Serve
	/usr/lib/go-1.18/src/net/http/server.go:3071 +0x4db
^Csignal: interrupt

```
terjadi karena typo pada file pada hostname
```
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

## Manual Testing
### Get All Categories
```curl
curl -X GET "http://localhost:8000/api/categories" \
     -H "Accept: application/json" \
     -H "Content-Type: application/json"

```
### Create New Category
```curl
curl -X POST "http://localhost:8000/api/categories" \
     -H "Accept: application/json" \
     -H "Content-Type: application/json" \
     -d '{"name": "Gadget"}'

```
### Get Category By Id
```curl
curl -X GET "http://localhost:8000/api/categories/1" \
     -H "Accept: application/json"

```
### Update Category By Id
```curl
curl -X PUT "http://localhost:8000/api/categories/1" \
     -H "Accept: application/json" \
     -H "Content-Type: application/json" \
     -d '{"name": "Fashion"}'
```

### Delete Category By Id
```
curl -X DELETE "http://localhost:8000/api/categories/2" \
     -H "Accept: application/json"

```
### Create Category Blank
```
curl -X POST "http://localhost:8000/api/categories" \
     -H "Accept: application/json" \
     -H "Content-Type: application/json" \
     -d '{"name": ""}'
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
Terjadi karena typo pada file pada hostname

```
023/05/05 14:10:31 http: panic serving 127.0.0.1:55558: Error 1146 (42S02): Table 'golang_resful_api.customer' doesn't exist
goroutine 21 [running]:
net/http.(*conn).serve.func1()
	/usr/lib/go-1.18/src/net/http/server.go:1825 +0xbf
panic({0x6ba900, 0xc0002a0018})
	/usr/lib/go-1.18/src/runtime/panic.go:844 +0x258
golang-resful-api/helper.CommitOrRollback(0xc0000e5e40?)
	/home/bismillah/GoLang/src/golang-resful-api/helper/tx.go:10 +0x8b
panic({0x6ba900, 0xc0002a0018})
	/usr/lib/go-1.18/src/runtime/panic.go:838 +0x207
golang-resful-api/helper.PanicIfError(...)
	/home/bismillah/GoLang/src/golang-resful-api/helper/error.go:5
golang-resful-api/repository.(*CategoryRepositoryImpl).Save(0xc0000b02a0?, {0x780230, 0xc0000e5e40}, 0x6c3280?, {0xc0001bb920?, {0xc0002209f8?, 0xc0001b3a68?}})
	/home/bismillah/GoLang/src/golang-resful-api/repository/category_repository_impl.go:21 +0xe5
golang-resful-api/service.(*CategoryServiceImpl).Create(0xc0000cff40, {0x780230, 0xc0000e5e40}, {{0xc0002209f8?, 0x203000?}})
	/home/bismillah/GoLang/src/golang-resful-api/service/category_service_impl.go:40 +0x14a
golang-resful-api/controller.(*CategoryControllerImpl).Create(0xc0001bb680, {0x780080, 0xc00022a0e0}, 0xc00013cb00, {0xc0000d0ff0?, 0x0?, 0xc0001f3065?})
	/home/bismillah/GoLang/src/golang-resful-api/controller/category_controller_impl.go:27 +0xc4
github.com/julienschmidt/httprouter.(*Router).ServeHTTP(0xc0000b0300, {0x780080, 0xc00022a0e0}, 0xc00013cb00)
	/home/bismillah/GoLang/pkg/mod/github.com/julienschmidt/httprouter@v1.3.0/router.go:387 +0x82b
net/http.serverHandler.ServeHTTP({0x77f430?}, {0x780080, 0xc00022a0e0}, 0xc00013cb00)
	/usr/lib/go-1.18/src/net/http/server.go:2916 +0x43b
net/http.(*conn).serve(0xc0002281e0, {0x7802d8, 0xc0001cadb0})
	/usr/lib/go-1.18/src/net/http/server.go:1966 +0x5d7
created by net/http.(*Server).Serve
	/usr/lib/go-1.18/src/net/http/server.go:3071 +0x4db

```
Terjadi karena salah nama table

```
2023/05/05 14:38:25 http: panic serving 127.0.0.1:55568: Error 1054 (42S22): Unknown column 'id' in 'field list'
goroutine 23 [running]:
net/http.(*conn).serve.func1()
	/usr/lib/go-1.18/src/net/http/server.go:1825 +0xbf
panic({0x6ba900, 0xc0000b02e8})
	/usr/lib/go-1.18/src/runtime/panic.go:844 +0x258
golang-resful-api/helper.CommitOrRollback(0xc00006c100?)
	/home/bismillah/GoLang/src/golang-resful-api/helper/tx.go:10 +0x8b
panic({0x6ba900, 0xc0000b02e8})
	/usr/lib/go-1.18/src/runtime/panic.go:838 +0x207
golang-resful-api/helper.PanicIfError(...)
	/home/bismillah/GoLang/src/golang-resful-api/helper/error.go:5
golang-resful-api/repository.(*CategoryRepositoryImpl).FindById(0x203000?, {0x780230, 0xc00006c100}, 0xc000026060?, 0x7f0cdd517e50?)
	/home/bismillah/GoLang/src/golang-resful-api/repository/category_repository_impl.go:46 +0x1ab
golang-resful-api/service.(*CategoryServiceImpl).FindById(0xc0000d3f40, {0x780230, 0xc00006c100}, 0x18?)
	/home/bismillah/GoLang/src/golang-resful-api/service/category_service_impl.go:76 +0xd9
golang-resful-api/controller.(*CategoryControllerImpl).FindById(0xc0001bb6a0, {0x780080, 0xc000282000}, 0xc00007c000, {0xc000026060, 0x1, 0xc000020113?})
	/home/bismillah/GoLang/src/golang-resful-api/controller/category_controller_impl.go:75 +0x15b
github.com/julienschmidt/httprouter.(*Router).ServeHTTP(0xc000092360, {0x780080, 0xc000282000}, 0xc00007c000)
	/home/bismillah/GoLang/pkg/mod/github.com/julienschmidt/httprouter@v1.3.0/router.go:387 +0x82b
net/http.serverHandler.ServeHTTP({0xc00006a030?}, {0x780080, 0xc000282000}, 0xc00007c000)
	/usr/lib/go-1.18/src/net/http/server.go:2916 +0x43b
net/http.(*conn).serve(0xc000228140, {0x7802d8, 0xc0001cadb0})
	/usr/lib/go-1.18/src/net/http/server.go:1966 +0x5d7
created by net/http.(*Server).Serve
	/usr/lib/go-1.18/src/net/http/server.go:3071 +0x4db
```
Terjadi karena salah pada query ke database, solusinya rubah
```sql
SELECT id, name WHERE id = ?
```
jadi
```sql
SELECT id, name FROM category WHERE id = ?
```

```
mysql] 2023/05/05 14:57:23 connection.go:299: invalid connection
2023/05/05 14:57:23 http: panic serving 127.0.0.1:55576: invalid connection
goroutine 34 [running]:
net/http.(*conn).serve.func1()
	/usr/lib/go-1.18/src/net/http/server.go:1825 +0xbf
panic({0x6b2120, 0xc00016ef40})
	/usr/lib/go-1.18/src/runtime/panic.go:844 +0x258
golang-resful-api/helper.PanicIfError(...)
	/home/bismillah/GoLang/src/golang-resful-api/helper/error.go:5
golang-resful-api/helper.CommitOrRollback(0xc000099660?)
	/home/bismillah/GoLang/src/golang-resful-api/helper/tx.go:9 +0x7c
panic({0x6b2120, 0xc000012270})
	/usr/lib/go-1.18/src/runtime/panic.go:838 +0x207
golang-resful-api/helper.PanicIfError(...)
	/home/bismillah/GoLang/src/golang-resful-api/helper/error.go:5
golang-resful-api/repository.(*CategoryRepositoryImpl).Update(0xc00006a300?, {0x780230, 0xc00029e0c0}, 0x6cb560?, {0xc0002ac090?, {0xc00029c0d0?, 0x0?}})
	/home/bismillah/GoLang/src/golang-resful-api/repository/category_repository_impl.go:33 +0x105
golang-resful-api/service.(*CategoryServiceImpl).Update(0xc000027f80, {0x780230, 0xc00029e0c0}, {0xc00029e100?, {0xc00029c0d0?, 0x7f0195b5d0e8?}})
	/home/bismillah/GoLang/src/golang-resful-api/service/category_service_impl.go:57 +0x1c4
golang-resful-api/controller.(*CategoryControllerImpl).Update(0xc00016f6c0, {0x780080, 0xc0002b2000}, 0xc0002a2000, {0xc000292020, 0x1, 0xc000288044?})
	/home/bismillah/GoLang/src/golang-resful-api/controller/category_controller_impl.go:46 +0x185
github.com/julienschmidt/httprouter.(*Router).ServeHTTP(0xc00006a360, {0x780080, 0xc0002b2000}, 0xc0002a2000)
	/home/bismillah/GoLang/pkg/mod/github.com/julienschmidt/httprouter@v1.3.0/router.go:387 +0x82b
net/http.serverHandler.ServeHTTP({0x77f430?}, {0x780080, 0xc0002b2000}, 0xc0002a2000)
	/usr/lib/go-1.18/src/net/http/server.go:2916 +0x43b
net/http.(*conn).serve(0xc000290000, {0x7802d8, 0xc00017edb0})
	/usr/lib/go-1.18/src/net/http/server.go:1966 +0x5d7
created by net/http.(*Server).Serve
	/usr/lib/go-1.18/src/net/http/server.go:3071 +0x4db
clera
```
Terjadi karena mengirim beberapa perintah, namun perintah salah satu perintahnya belum ditutup
Solusinya tutup koneksi ke database setelah selesai melakukan query
```go
defer rows.Close()
```


## Implementation middleware with jwt and gorilla sessions

For database, better using https://railway.app/ :)

Configure database on models/database.go

```bash
host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai
```

#### How to run

```bash
  go run main.go
```

#### API 

```bash
http:localhost:3000/login
http:localhost:3000/register
http:localhost:3000/change-password
```
```bash
  {
	"username": "",
	"password": ""
}
```

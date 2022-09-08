# Simple fullnode checker

### Comands<br/>
POST `/start/`<br/>
```
{ 
  Ip
  Seconds
}
```
GET `/stop/`<br/>
GET `/jobs/`<br/>

### Packages<br/>
go get github.com/antchfx/htmlquery - XPath query package for the HTML document<br/>
// go get github.com/jinzhu/gorm - db connector<br/>
// go get github.com/jinzhu/gorm/dialects/mysql - dialect<br/>
go get github.com/gorilla/mux - router<br/>
go get github.com/go-co-op/gocron - планировщик задач<br/>

### MySQL<br/>
```
docker run --detach --name=test --env="MYSQL_ROOT_PASSWORD=password" --publish 6603:3306 mysql
mysql -u root -p
CREATE DATABASE;
```
### Docker<br/>
```
docker
docker build -t sch/dev:latest .
docker run -p 8080:8080 -d sch/dev:latest
```

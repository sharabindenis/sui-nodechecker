# Simple sui fullnode checker

### Telegram bot
apiendpoint const
```
{
    chatid string
    key	string
    message string
}
```
### Web commands<br/>
//deprecated<br/>
POST `/start/`<br/>
```
{ 
  Ip string
  Seconds int
}
```
POST `/stop/`<br/>
```
{
  Ip string
}
```
GET `/tasks/`<br/>

### Packages<br/>
go get github.com/antchfx/htmlquery - XPath query package for the HTML document<br/>
// go get github.com/jinzhu/gorm - db connector<br/>
// go get github.com/jinzhu/gorm/dialects/mysql - dialect<br/>
go get github.com/gorilla/mux - router<br/>
go get github.com/go-co-op/gocron - crone RemoveByRef<br/>
github.com/Syfaro/telegram-bot-api - tg connector

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
### Test data</br
```
/ip 5.23.48.96
/ip 45.8.147.68
/ip 159.223.119.225
```
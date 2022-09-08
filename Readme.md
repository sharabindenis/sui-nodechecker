Simple fullnode checker

/startsch
Ip
Seconds



packages
go get github.com/antchfx/htmlquery - XPath query package for the HTML document
go get github.com/jinzhu/gorm - db connector
go get github.com/jinzhu/gorm/dialects/mysql - dialect
go get github.com/gorilla/mux - router
go get github.com/go-co-op/gocron - планировщик задач

MySQL
docker run --detach --name=test --env="MYSQL_ROOT_PASSWORD=password" --publish 6603:3306 mysql
mysql -u root -p
CREATE DATABASE;


fullnode, err := post sui_getTotalTransactionNumber node_ip
devnet, err :=post sui_getTotalTransactionNumber https://fullnode.devnet.sui.io/

docker
docker build -t sch/dev:latest .
docker run -p 8080:8080 -d sch/dev:latest
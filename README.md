# gin-gorm-ex

How to use gin gorm with mysql

## pkg

go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/mysql

## Access to mysql

sudo docker exec -it fc3ccaf40db2 /bin/bash
mysql -u root -h localhost -P 3306 hackernews -p
DROP TABLE IF EXISTS user;
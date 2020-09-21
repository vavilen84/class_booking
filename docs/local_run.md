### Run project with local MySQL server

##### Environment

Create .env file
```
cat .env.dist >> .env
```
Set PROJECT_ROOT env var according to your project path. It should forward to 'server' Golang application. Ex.:
- Root path is /var/www/class_booking
- PROJECT_ROOT=/var/www/class_booking/server

##### MySQL

Install [MySQL](https://dev.mysql.com/doc/mysql-installation-excerpt/5.7/en/) on your local machine. 
To make your work with project easier you can use [PHPMyAdmin](https://www.phpmyadmin.net/) MySQL client. 


##### Run tests

```
cd PROJECT_ROOT
go test ./... -p 1 -count=1 -v
```

##### Run server

```
cd PROJECT_ROOT
go run main.go
```

##### Run migrations 

```
cd PROJECT_ROOT
go run cli/db/migrate/up/up.go
```

##### Postman

Export [Postman collection](../postman/Class%20Booking.postman_collection.json) to your Postman

Now you can run Postman queries!

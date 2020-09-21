### Run project with local MySQL server

##### Environment

Create .env file
```
cd path/to/prject
cat .env.dist >> .env
```
Set PROJECT_ROOT env var according to your project path. It should forward to 'server' Golang application. Ex.:
- path/to/project is /var/www/class_booking
- PROJECT_ROOT=/var/www/class_booking/server

##### Make all .sh files executable
```
cd path/to/project
find . -type f -iname "*.sh" -exec chmod +x {} \;
```

##### MySQL

Install [MySQL](https://dev.mysql.com/doc/mysql-installation-excerpt/5.7/en/) on your local machine. 
To make your work with project easier you can use [PHPMyAdmin](https://www.phpmyadmin.net/) MySQL client. 

##### Run migrations 

```
cd PROJECT_ROOT (path to 'server' Golang APP)
go run cli/db/migrate/up/up.go
```

##### Run tests

```
cd PROJECT_ROOT (path to 'server' Golang APP)
go test ./... -p 1 -count=1 -v
```

##### Run server

```
cd PROJECT_ROOT (path to 'server' Golang APP)
go run main.go
```

##### Postman

Export [Postman collection](../postman/Class%20Booking.postman_collection.json) to your Postman

Now you can run Postman queries!

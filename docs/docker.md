### Run project with Docker (MySQL server is a docker container)

##### Install Docker

In order to use this workflow for development, you must first [install Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/).

##### Environment

Create .env file
```
cat .env.dist >> .env
```
Set PROJECT_ROOT env var according to your project path. It should forward to 'server' Golang application. Ex.:
- path/to/project is /var/www/class_booking
- PROJECT_ROOT=/var/www/class_booking/server

##### Run project

```
cd path/to/project
shell/docker_sql/docker_up.sh
```

##### Run migrations

```
cd path/to/project
shell/docker_sql/migrate_apply_all.sh
```

##### PMA

Enter this data on login:
- Server: db 
- User: root 
- Password: 123456 

Databases:
- Database is 'class_booking' 
- Test database is 'class_booking_test'

Database will be created automatically.


##### Run tests
```
cd path/to/project
shell/docker_sql/run_all_server_tests.sh
```

##### Postman

Export [Postman collection](../postman/Class%20Booking.postman_collection.json) to your Postman

Now you can run Postman queries!


##### Stop project

```
cd path/to/project
shell/docker_sql/docker_down.sh
```

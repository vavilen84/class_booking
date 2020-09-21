### Run project with Docker (MySQL server is a docker container)

##### Install Docker

In order to use this workflow for development, you must first [install Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/).

##### Environment

Create .env file
```
cat .env.dist >> .env
```
Set PROJECT_ROOT env var according to your project path. It should forward to 'server' Golang application. Ex.:
- Root path is /var/www/class_booking
- PROJECT_ROOT=/var/www/class_booking/server

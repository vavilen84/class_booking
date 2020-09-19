# Class booking project

[Server application design](docs/server-design.md)

## Run project

### Install Docker

In order to use this workflow for development, you must first [install Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/).

### Environment

Create .env file (just copy env.dist content will be enough in order to run project)
```
cat .env.dist >> .env
```

### Console commands

#### Docker

Docker-compose up
```
shell/docker_up.sh
```

Docker-compose up (with --build flag)
```
shell/docker_up_build.sh
```

Docker-compose down
```
shell/docker_down.sh
```

#### Migrations

Create migration
```
shell/migrate_create.sh migration_name
```

Apply all migrations
```
shell/server_build.sh
shell/migrate_apply_all.sh
```

### PHPMyAdmin

Enter this data on login:
- Server: db (according to docker-compose.yaml)
- User: root (according to .env.dist)
- Password: 123456 (according to .env.dist)

Database is 'class_booking' (according to .env.dist)

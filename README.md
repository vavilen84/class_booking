# Class booking project

There are 3 possible ways to run project locally:
- [Run project without Docker](docs/local_run.md)
- [Run project with Docker (MySQL server is on host machine)](docs/docker_mysql_on_host_machine.md)
- [Run project with Docker (MySQL server is a docker container](docs/docker.md)

[Database design](docs/db-design.md)












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

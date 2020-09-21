docker-compose -f docker/dev/docker-compose.yaml --env-file=.env run server go run cli/db/migrate/up/up.go

docker-compose -f docker/dev/docker-compose.host-mysql.yaml --env-file=.env run server go test ./... -p 1 -count=1 -v

docker-compose -f docker/dev/docker-compose.yaml --env-file=.env run server cd models; go test . -v

.PHONY=run
run:
	go run src/main.go

.PHONY=build
build:
	go build -o ./bin/main ./src/main.go

.PHONY=docker
docker:
	docker build -t ejosephodom/go-open:latest . && docker push ejosephodom/go-open:latest

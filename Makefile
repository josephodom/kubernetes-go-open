DOCKER_HUB_USER:=ejosephodom
DOCKER_HUB_TAG:=latest
APP:=go-open
APP_PORT_EXTERNAL:=8080
APP_PORT_INTERNAL:=8080

.PHONY=run
run:
	go run src/main.go

.PHONY=build
build:
	go build -o ./bin/main ./src/main.go

.PHONY=docker
docker:
	docker build -t $(DOCKER_HUB_USER)/$(APP):$(DOCKER_HUB_TAG) . && docker push $(DOCKER_HUB_USER)/$(APP):$(DOCKER_HUB_TAG)

.PHONY=kube
kube:
	kubectl apply -f deployment.yaml
	kubectl apply -f deployment.redis.yaml

.PHONY=kube-restart
kube-restart:
	kubectl apply -f deployment.yaml
	kubectl apply -f deployment.redis.yaml
	kubectl rollout restart deployment/$(APP)

.PHONY=kube-fwd
kube-fwd:
	kubectl port-forward deployment/$(APP) $(APP_PORT_EXTERNAL):$(APP_PORT_INTERNAL)

.PHONY=stress
stress:
	while true; do curl http://localhost:$(APP_PORT_EXTERNAL); sleep 0.25; done

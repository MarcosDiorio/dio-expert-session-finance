build-image:
	docker build -t MarcosDiorio/finance .

run-app:
	docker-compose -f .devops/app.yml up -d 

lint: 
	golint ./...
	go fmt ./...


unit_test:
	go test ./...
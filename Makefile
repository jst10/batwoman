run:
	cd cmd; go run *.go

build:
	cd cmd; go build -o ../bin/batwoman

stop-local-db:
	docker rm -f mysql

start-local-db:
	docker run --name mysql  -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=batwoman -p 3306:3306 -d mysql:latest

restart-local-db:stop-local-db start-local-db


integration-test:
	cd integration-tests; go run *.go
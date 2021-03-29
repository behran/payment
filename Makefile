build-composer:
	docker-compose -f docker-compose.yml up -d
migration:
	docker exec -it payment_app_1 ./migration
app-start:
	docker exec -it payment_app_1 ./app

test:
	docker-compose -f docker-compose-test.yml up -d
	docker exec -it payment_apptest_1 ./migration
	docker exec -it payment_apptest_1 go test -v ./...
	docker-compose -f docker-compose-test.yml down





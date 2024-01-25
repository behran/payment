build-composer:
	docker-compose -f docker-compose.yml up -d
migration-up:
	docker exec -it payment_app_1 ./cli m postgres up
migration-down:
	docker exec -it payment_app_1 ./cli m postgres down
app-start:
	docker exec -it payment_app_1 ./app
test:
	go test -v ./...





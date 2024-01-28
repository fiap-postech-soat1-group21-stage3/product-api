
run-product-api:
	cd product-api/cmd && go run main.go

run-project:
	@echo "-------------Initializing container-------------"
	docker compose -f "docker-compose.yml" up -d --build
	@echo
	@echo "-------------Initializing scheduler-------------"
	make run-product-api



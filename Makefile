.PHONY: dev prod

dev:
	@echo "Running in development environment"
	docker compose --env-file .env.development up
dev-web-shell:
	@echo "Entering shell..."
	 docker compose --env-file .env.development run --service-ports web bash

prod:
	@echo "Running in production environment"
	docker compose --env-file .env.production up
prod-web-shell:
	@echo "Entering shell..."
	 docker compose --env-file .env.production run --service-ports web bash
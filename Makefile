
## up: start containerized app in the background
up:
	"$(CURDIR)/scripts/start_application.sh"

## down: stop docker compose
down:
	"$(CURDIR)/scripts/stop_application.sh"

## restart: restart containerized app
restart:
	"$(CURDIR)/scripts/restart_application.sh"

## up_rebuild: rebuild application and start it
up_rebuild:
	"$(CURDIR)/scripts/rebuild_application.sh"


#	@echo "Stopping application..."
#	docker compose down
#	@echo "Building and starting application..."
#	docker compose up --build -d
#	@echo "Application has been built and started!"

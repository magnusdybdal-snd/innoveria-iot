SERVICES := services/device-service services/collection-service services/auth-service

.PHONY: swagger-deps swagger-gen

## Install swaggo dependencies in all services (run once per new service)
swagger-deps:
	@for svc in $(SERVICES); do \
		echo "Installing swagger deps in $$svc"; \
		cd $(CURDIR)/$$svc && go get github.com/swaggo/http-swagger@latest && go get -tool github.com/swaggo/swag/cmd/swag && go mod tidy; \
	done

## Regenerate swagger docs in all services (run after changing annotations)
swagger-gen:
	@for svc in $(SERVICES); do \
		echo "Generating swagger docs in $$svc"; \
		cd $(CURDIR)/$$svc && go tool swag init -g cmd/main.go; \
	done

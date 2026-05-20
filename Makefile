# Services that need swagger deps (includes api-gateway which serves the merged UI)
SERVICES := services/device-service services/collection-service services/auth-service services/onboarding-service services/context-service services/api-gateway

# Services that have swagger annotations and need doc generation (excludes api-gateway)
ANNOTATED_SERVICES := services/device-service services/collection-service services/auth-service services/onboarding-service services/context-service

# Pinned versions — update here when upgrading, keep all services in sync
SWAGGO_HTTP_SWAGGER_VERSION := v1.3.4
SWAGGO_CLI_VERSION           := v1.16.6

.PHONY: swagger-deps swagger-gen

## Install swaggo dependencies in all services (run once per new service)
swagger-deps:
	@for svc in $(SERVICES); do \
		echo "Installing swagger deps in $$svc"; \
		cd $(CURDIR)/$$svc && go get github.com/swaggo/http-swagger@$(SWAGGO_HTTP_SWAGGER_VERSION) && go get -tool github.com/swaggo/swag/cmd/swag@$(SWAGGO_CLI_VERSION) && go mod tidy; \
	done

## Regenerate swagger docs in annotated services (run after changing annotations)
swagger-gen:
	@for svc in $(ANNOTATED_SERVICES); do \
		echo "Generating swagger docs in $$svc"; \
		cd $(CURDIR)/$$svc && go tool swag init -g cmd/main.go; \
	done

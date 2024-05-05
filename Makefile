############################# HELP MESSAGE #############################
# Make sure the help command stays first, so that it's printed by default when `make` is called without arguments
.PHONY: help tests
help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


-----------------------------: ## 
___CONTRACTS___: ## 
gen-contracts-bindings: ## generates contracts bindings
	cd contracts && ./gen_bindings.sh

______API______:
gen-api-bindings: ## generate protobuf binding for grpc server
	cd api && ./gen_pb.sh


-----------------------------: ## 
# TODO: WIP
___DOCKER___: ## 
docker-build-and-publish-images: ## builds and publishes operator and aggregator docker images using Ko
	KO_DOCKER_REPO=ghcr.io/layr-labs/incredible-squaring ko build aggregator/cmd/main.go --preserve-import-paths
	KO_DOCKER_REPO=ghcr.io/layr-labs/incredible-squaring ko build operator/cmd/main.go --preserve-import-paths


-----------------------------: ## 
____OFFCHAIN_SOFTWARE___: ## 
start-operators:
	cd operator && ./start.sh

start-relayer:
	cd relayer && ./start.sh

-----------------------------: ## 
_____HELPER_____: ## 
mocks: ## generates mocks for tests
	go install go.uber.org/mock/mockgen@v0.4.0
	go generate ./...

run-unit-tests: ## runs all unit tests
	go test $$(go list ./... | grep -v /bindings | grep -v /pb) -coverprofile=coverage.out -covermode=atomic --timeout 15s
	go tool cover -html=coverage.out -o coverage.html

test-integration: ## runs all integration tests
	go test ./tests/integration/... -v -count=1

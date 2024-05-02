############################# HELP MESSAGE #############################
# Make sure the help command stays first, so that it's printed by default when `make` is called without arguments
.PHONY: help tests
help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


-----------------------------: ## 

___CONTRACTS___: ## 

build-contracts: ## builds all contracts
	cd contracts && forge build

bindings: ## generates contracts bindings
	cd contracts && ./gen_bindings.sh

______API______:

gen-api-pb: ## generate api pbs
	cd api && ./gen_pb.sh


-----------------------------: ## 
# TODO: WIP
___DOCKER___: ## 
docker-build-and-publish-images: ## builds and publishes operator and aggregator docker images using Ko
	KO_DOCKER_REPO=ghcr.io/layr-labs/incredible-squaring ko build aggregator/cmd/main.go --preserve-import-paths
	KO_DOCKER_REPO=ghcr.io/layr-labs/incredible-squaring ko build operator/cmd/main.go --preserve-import-paths


-----------------------------: ## 
# TODO: WIP
____OFFCHAIN_SOFTWARE___: ## 
start-aggregator: ## 
	go run aggregator/cmd/main.go --config config-files/aggregator.yaml \
		--skate-deployment ${DEPLOYMENT_FILES_DIR}/skate_avs_deployment_output.json \
		--ecdsa-private-key ${AGGREGATOR_ECDSA_PRIV_KEY}

start-operator: ## 
	go run operator/cmd/main.go --config config-files/operator.anvil.yaml


-----------------------------: ## 
_____HELPER_____: ## 
mocks: ## generates mocks for tests
	go install go.uber.org/mock/mockgen@v0.4.0
	go generate ./...

test-unit: ## runs all unit tests
	go test $$(go list ./... | grep -v /bindings | grep -v /logging) -coverprofile=coverage.out -covermode=atomic --timeout 15s
	go tool cover -html=coverage.out -o coverage.html

test-contracts: ## runs all forge tests
	cd contracts && forge test

test-integration: ## runs all integration tests
	go test ./tests/integration/... -v -count=1

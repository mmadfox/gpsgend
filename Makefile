
.PHONY: docker-mock
docker-mock:
	@docker build -t github.com/mmadfox/gpsgend/mock:latest -f    \
           ./docker/mockgen.dockerfile .

.PHONY: mock
mocks: docker-mock
	@bash ./scripts/mockgen.sh

.PHONY: cover
cover:
	go test ./... -cover -coverprofile=cover.out
	go tool cover -func=cover.out
	go tool cover -html=cover.out	

.PHONY: test
test:
	go test -v ./...

.PHONY: test/integration	
test/integration:
	GPSGEND_INTEGRATION_TESTS_ENABLED=true go test -v -run TestIntegration ./tests/integration/...

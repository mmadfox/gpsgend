
.PHONY: docker-gpsgend
docker-gpsgend:
	@docker build -t github.com/mmadfox/gpsgend/mock:latest -f \
	    ./docker/gpsgend.dockerfile  .

.PHONY: docker-mock
docker-mock:
	@docker build -t github.com/mmadfox/gpsgend/mock:latest -f    \
           ./docker/mockgen.dockerfile .

.PHONY app:
app:
	@docker-compose up --build 

clean:
	@docker-compose down	

.PHONY: mock
mocks: docker-mock
	@bash ./scripts/mockgen.sh

.PHONY: docker-protoc
docker-protoc:
	@docker build -t github.com/mmadfox/gpsgend/protoc:latest -f    \
           ./docker/protoc.dockerfile .

.PHONY: proto
proto: docker-protoc
	@bash ./scripts/proto.sh

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

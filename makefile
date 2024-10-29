APP_NAME=go-pizza
BUILD_FILES=$(shell find . -name '*.go')
BUILD_PATH=build/$(APP_NAME)

# Application
build: clean
	mkdir -p build
	go build -o $(BUILD_PATH) $(BUILD_FILES)

run: build
	$(BUILD_PATH)

dev:
	CompileDaemon -command="./$(APP_NAME)"

# Database
clean-docker:
	docker-compose down && docker rmi -f `docker images -a -q`

db-dev:
	docker compose -f database-dev.yml up -d

clean:
	rm -f $(APP_NAME)
	rm -rf build

migrate:
	go run migrate/migrate.go
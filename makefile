APP_NAME=go-pizza
BUILD_FILES=$(bash find . -name '*.go')
BUILD_PATH="build/$(APP_NAME)"

# Application
build: clean
	go build -o $(BUILD_PATH) $(BUILD_FILES)

run: build
	$(BUILD_PATH)

dev: build
	CompileDaemon -command="$(BUILD_PATH)"

# Database
clean-docker:
	docker-compose down && docker rmi -f `docker images -a -q`

db-dev:
	docker compose -f database-dev.yml up -d

clean:
	rm -rf build
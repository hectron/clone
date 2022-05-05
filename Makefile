.PHONY: clean build

build: clean update-remote
	go build -o bin/clone -ldflags="-X 'main.DefaultOwner=hectron'"

update-remote:
	git remote update

clean:
	rm -rf ./bin

test:
	go test ./...

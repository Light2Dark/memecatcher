run:
	go run ./cmd/api
hot_reload_run:
	wgo run ./cmd/api
update_packages:
	go get -d -u ./...
	go mod tidy
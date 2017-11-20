vendor: Gopkg.lock Gopkg.toml
	rm -rf vendor/
	dep ensure -v
	dep prune -v

.PHONE: banking
banking: vendor
	docker build -t banking .

.PHONE: run
run: banking
	docker-compose down --volumes
	-docker-compose up
	docker-compose down --volumes

.PHONE: shell
shell: banking
	docker-compose down --volumes
	-docker-compose run --rm shell
	docker-compose down --volumes

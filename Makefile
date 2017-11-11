PROJECT_PREFIX := $(PWD)/src/banking

vendor: $(PROJECT_PREFIX)/Gopkg.lock $(PROJECT_PREFIX)/Gopkg.toml
	rm -rf $(PROJECT_PREFIX)/vendor
	-cd $(PROJECT_PREFIX); \
		dep ensure -v && \
		dep prune -v && \
	cd -;

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

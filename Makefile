UIDEPS := $(shell find ui/src/ -name "*")

all: banking

.PHONE: banking
banking: ui vendor
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

.PHONE: ui
ui: static/flag

static/flag: $(UIDEPS)
	rm -rf static/
	cd ui/ && ng build
	mkdir -p static/
	cp -r ui/dist/* static/
	touch static/flag

vendor: Gopkg.lock Gopkg.toml
	rm -rf vendor/
	dep ensure -v
	dep prune -v

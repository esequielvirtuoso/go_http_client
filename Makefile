# include help.mk

# tell Make the following targets are not files, but targets within Makefile
.PHONY: version clean audit lint install tag release check-env check-used-ports
.DEFAULT_GOAL := help

HUB_HOST     = hub.docker.com/repository/docker
HUB_USER 	 = esequielvirtuoso
HUB_REPO     = go_http_client

BUILD         	= $(shell git rev-parse --short HEAD)
VERSION  	  	= $(shell git describe --always --tags)
NAME           	= $(shell basename $(CURDIR))
IMAGE          	= $(HUB_USER)/$(HUB_REPO):$(BUILD)

check-used-ports:
	sudo netstat -tulpn | grep LISTEN

git-config:
	git config --replace-all core.hooksPath .githooks

print-env-%:
	echo "Env '$*' set as: ${${*}}";
		

check-env-%:
	@ if [ "${${*}}" = ""  ]; then \
		echo "Variable '$*' not set"; \
		exit 1; \
	fi

version: ##@other Check version.
	@echo $(VERSION)

clean: ##@dev Remove folder vendor, public and coverage.
	rm -rf vendor public coverage

install: clean ##@dev Download dependencies via go mod.
	GO111MODULE=on go mod download
	GO111MODULE=on go mod vendor

audit: ##@check Run vulnerability check in Go dependencies.
	DOCKER_BUILDKIT=1 docker build --progress=plain --target=audit --file=Dockerfile .

lint: ##@check Run lint on docker.
	DOCKER_BUILDKIT=1 \
	docker build --progress=plain \
		--target=lint \
		--file=Dockerfile .

test: ##@check Run tests and coverage.
	docker build --progress=plain \
		--tag $(IMAGE) \
		--target=test \
		--file=Dockerfile .

	-mkdir coverage
	docker create --name $(NAME)-$(BUILD) $(IMAGE)
	docker cp $(NAME)-$(BUILD):/index.html ./coverage/.
	docker rm -vf $(NAME)-$(BUILD)

tag: check-env-VERSION ##@build Add docker tag.
	docker tag $(IMAGE) \
		$(HUB_USER)/$(HUB_REPO):$(VERSION)

release: check-env-TAG ##@build Create and push git tag.
	git tag -a $(TAG) -m "Generated release "$(TAG)
	git push origin $(TAG)
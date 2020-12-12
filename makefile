VERSION := $(shell gogitver)

.PHONY: build unit-tests end-to-end-test test

build: clean
	docker-compose \
		-f docker-compose.yml \
		pull \
		--ignore-pull-failures

	docker-compose \
		-f docker-compose.yml \
		-f docker-compose.ci.yml \
		build \
		--pull

unit-tests: build
	docker-compose \
		-f docker-compose.yml \
		-f docker-compose.ci.yml \
		up \
		--exit-code-from unit --abort-on-container-exit \
		unit
	cd artifacts && curl -s https://codecov.io/bash | bash

end-to-end-test: clean build
	docker-compose \
		-f docker-compose.yml \
		-f docker-compose.ci.yml \
		up \
		--exit-code-from endtoend --abort-on-container-exit \
		endtoend

test: unit-tests end-to-end-test

clean:
	docker-compose \
	-f docker-compose.yml \
	-f docker-compose.ci.yml \
	down

generate-end-to-end:
	cd testing && godog --strict --format=pretty ../docs/features/

generate:
	statik -f -src=./docs/migrations -dest=./internal/migrations
	rm -rf ./internal/definitions
	../kafmesh/artifacts/kafmesh-gen docs/kafmesh/definition.yml

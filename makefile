VERSION := $(shell gogitver)

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

end-to-end-test: clean build
	docker-compose \
		-f docker-compose.yml \
		-f docker-compose.ci.yml \
		up \
		--exit-code-from endtoend --abort-on-container-exit \
		endtoend

test: end-to-end-test

clean:
	docker-compose \
	-f docker-compose.yml \
	-f docker-compose.ci.yml \
	down

generate-end-to-end:
	cd testing && godog --strict --format=pretty ../docs/features/

generate:
	statik -f -src=./docs/migrations -dest=./internal/migrations
	../kafmesh/artifacts/kafmesh-gen docs/kafmesh/definition.yml

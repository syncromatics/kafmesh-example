build: generate
	go build ./...

generate:
	../kafmesh/artifacts/kafmesh-gen docs/kafmesh/definition.yml
	
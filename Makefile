VERSION = v2.5.3

all:
	# Clean up orphan files.
	rm -rf config
	# Generate the actual code.
	docker run --rm -ti \
		--user $$(id -u):$$(id -g) \
		--mount type=bind,source=$$(pwd),destination=/mnt/src \
		openapitools/openapi-generator-cli \
		generate \
		--git-host github.com \
		--git-user-id eschercloudai \
		--git-repo-id argocd-client-go/client \
		-g go \
		-p packageName=client \
		-i https://raw.githubusercontent.com/argoproj/argo-cd/$(VERSION)/assets/swagger.json  \
		-o /mnt/src/client
	# Pull in any deps missing from auto-generation (e.g. test files).
	cd client; go mod tidy

.PHONY: test
test:
	cd client; go test ./...

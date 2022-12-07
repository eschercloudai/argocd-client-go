VERSION = v2.5.3

all:
	wget -q https://raw.githubusercontent.com/argoproj/argo-cd/$(VERSION)/assets/swagger.json -O swagger.json.orig
	# We need to fix the broken...
	cat swagger.json.orig | go run github.com/evanphx/json-patch/v5/cmd/json-patch@v5.6.0 -p patches.json > swagger.json
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
		-i /mnt/src/swagger.json \
		-o /mnt/src/client
	# Do some fixes...
	rm -rf go.* client/go.*
	go mod init github.com/eschercloudai/argocd-client-go
	go mod tidy
	# Cleanup
	rm swagger.*

.PHONY: test
test:
	go test ./...

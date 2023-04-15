
FUNCTION_NAME := whatsapp
GO_FILES := $(shell find . -name "*.go" -type f)
GO_LDFLAGS := -ldflags="-s -w"

.PHONY: build
build: clean
    GOOS=linux GOARCH=amd64 go build $(GO_LDFLAGS) -o bin/$(FUNCTION_NAME) ./cmd/$(FUNCTION_NAME)

.PHONY: clean
clean:
    rm -rf ./bin

.PHONY: test
test:
    go test -v -cover ./...

.PHONY: deploy
deploy: build
    aws cloudformation package \
        --template-file template.yaml \
        --s3-bucket my-bucket \
        --output-template-file packaged.yaml
    aws cloudformation deploy \
        --template-file packaged.yaml \
        --stack-name $(FUNCTION_NAME)-stack \
        --capabilities CAPABILITY_IAM \
        --parameter-overrides FunctionName=$(FUNCTION_NAME) \
        --no-fail-on-empty-changeset

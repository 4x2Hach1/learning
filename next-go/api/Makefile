.PHONY: help init gen example run
.DEFAULT_GOAL := help

help: ## ヘルプを表示する
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

init: ## 環境構築
	go install goa.design/goa/v3/cmd/goa@v3
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/cosmtrek/air@latest

gen: ## designからの生成
	goa gen github.com/4x2Hach1/learning/next-go/api/design

example: ## exampleの生成
	goa example github.com/4x2Hach1/learning/next-go/api/design

run: ## サーバーの起動
	go run ./cmd/server
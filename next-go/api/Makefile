.PHONY: help init db-up db-down db-ver db-create gen example run test
.DEFAULT_GOAL := help

help: ## ヘルプを表示する
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

init: ## 環境構築
	go install goa.design/goa/v3/cmd/goa@v3
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/cosmtrek/air@latest
	go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest

db-up: ## migrate
	export $$(cat ./.env | grep -v ^# | xargs); migrate -database="$${DB_DRIVER}://$${DB_USER}:$${DB_PASS}@tcp($${DB_HOST}:$${DB_PORT})/$${DB_NAME}?multiStatements=true" -path=database/migrations/ up
db-down: ## migrate
	export $$(cat ./.env | grep -v ^# | xargs); migrate -database="$${DB_DRIVER}://$${DB_USER}:$${DB_PASS}@tcp($${DB_HOST}:$${DB_PORT})/$${DB_NAME}?multiStatements=true" -path=database/migrations/ down
db-down-%: ## migrate
	export $$(cat ./.env | grep -v ^# | xargs); migrate -database="$${DB_DRIVER}://$${DB_USER}:$${DB_PASS}@tcp($${DB_HOST}:$${DB_PORT})/$${DB_NAME}?multiStatements=true" -path=database/migrations/ down ${@:db-down-%=%}
db-ver: ## migrate
	export $$(cat ./.env | grep -v ^# | xargs); migrate -database="$${DB_DRIVER}://$${DB_USER}:$${DB_PASS}@tcp($${DB_HOST}:$${DB_PORT})/$${DB_NAME}?multiStatements=true" -path=database/migrations/ version
db-create-%: ## migrate
	migrate create -ext sql -dir database/migrations -seq ${@:db-create-%=%}
db-force-%: ## migrate
	export $$(cat ./.env | grep -v ^# | xargs); migrate -database="$${DB_DRIVER}://$${DB_USER}:$${DB_PASS}@tcp($${DB_HOST}:$${DB_PORT})/$${DB_NAME}?multiStatements=true" -path=database/migrations/ force ${@:db-force-%=%}

gen: ## designからの生成
	goa gen github.com/4x2Hach1/learning/next-go/api/design

example: ## exampleの生成
	goa example github.com/4x2Hach1/learning/next-go/api/design

run: ## サーバーの起動
	export $$(cat ./.env | grep -v ^# | xargs); air

test: ## すべてのテスト
	export $$(cat ./.env | grep -v ^# | xargs); go test -cover -v -bench . -benchmem ./...

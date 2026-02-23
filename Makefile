# ===== Docker / DB =====
.PHONY: db-up db-down db-reset migrate

db-up: ## PostgreSQL コンテナ起動
	docker compose up -d

db-down: ## PostgreSQL コンテナ停止
	docker compose down

db-reset: ## PostgreSQL コンテナ停止 + データ削除
	docker compose down -v

migrate: ## DBマイグレーション実行
	cd backend && go run cmd/migrate/main.go

# ===== バックエンド =====
.PHONY: generate server test-repo test-backend

generate: ## oapi-codegen によるコード生成
	cd backend && go generate ./gen/...

server: ## バックエンド開発サーバー起動
	cd backend && go run cmd/server/main.go

test-repo: ## リポジトリ層の結合テスト
	cd backend && go test ./infrastructure/persistence/ -v

test-backend: ## バックエンド全テスト
	cd backend && go test ./... -v

# ===== フロントエンド =====
.PHONY: front-dev front-build front-lint

front-dev: ## フロントエンド開発サーバー起動
	cd frontend && npm run dev

front-build: ## フロントエンドビルド
	cd frontend && npm run build

front-lint: ## フロントエンド Lint
	cd frontend && npm run lint

# ===== ヘルプ =====
.PHONY: help
help: ## コマンド一覧を表示
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

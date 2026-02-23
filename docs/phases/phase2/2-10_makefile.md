# 2-10. Phase 2 のコマンドを Makefile に登録

## 実施内容
- プロジェクトルートに Makefile を作成
- 定常的に使うコマンドを集約

## コマンド一覧
| コマンド | 説明 |
|---------|------|
| `make db-up` | PostgreSQL コンテナ起動 |
| `make db-down` | PostgreSQL コンテナ停止 |
| `make db-reset` | コンテナ停止 + データ削除 |
| `make migrate` | DB マイグレーション実行 |
| `make generate` | oapi-codegen コード生成 |
| `make server` | バックエンド開発サーバー起動 |
| `make test-repo` | リポジトリ層の結合テスト |
| `make test-backend` | バックエンド全テスト |
| `make front-dev` | フロントエンド開発サーバー起動 |
| `make front-build` | フロントエンドビルド |
| `make front-lint` | フロントエンド Lint |
| `make help` | コマンド一覧表示 |

## 備考
- Windows 環境では `make` のインストールが必要 (`choco install make` 等)

# 2-5. GORM モデル・マイグレーション実装

## 実施内容
- `infrastructure/persistence/model.go` に GORM モデルを定義
- `infrastructure/config/database.go` に DB接続・マイグレーション関数を実装
- `cmd/migrate/main.go` にマイグレーション実行用コマンドを作成
- ローカル PostgreSQL に対してマイグレーションを実行し、テーブル作成を確認

## GORM モデル
### UserModel
- `users` テーブルに対応
- `email` に UNIQUE インデックス
- `deleted_at` のデフォルトは `0001-01-01 00:00:00`

### TaskModel
- `tasks` テーブルに対応
- `user_id` に外部キー (→ users.id) + インデックス
- `description` のデフォルトは空文字
- `status` のデフォルトは `todo`
- `deleted_at` のデフォルトは `0001-01-01 00:00:00`

## DB接続情報
環境変数で設定可能（デフォルト値あり）:
- `DB_HOST` (localhost), `DB_PORT` (5432), `DB_USER` (app), `DB_PASSWORD` (password), `DB_NAME` (claude_task_app)

## コマンド
- `go run cmd/migrate/main.go` — マイグレーション実行

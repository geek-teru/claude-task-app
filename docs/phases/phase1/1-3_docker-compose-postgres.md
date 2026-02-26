# 1-3. Docker Compose で PostgreSQL 環境構築

## 実施内容
- `docker-compose.yml` を作成し、PostgreSQL 16 のコンテナを定義

## 接続情報
| 項目 | 値 |
|------|-----|
| ホスト | localhost |
| ポート | 5432 |
| ユーザー | app |
| パスワード | password |
| データベース | claude_task_app |

## コマンド
- `docker compose up -d` — コンテナ起動
- `docker compose down` — コンテナ停止
- `docker compose down -v` — コンテナ停止 + データ削除

# 1-4. PostgreSQL 接続確認

## 実施内容
- `docker compose up -d` でコンテナを起動
- `docker exec` 経由で psql 接続し、バージョンを確認

## 確認結果
- PostgreSQL 16.12 が正常に起動・接続できることを確認

## 備考
- ポート 5432 が既存プロセスと競合したため、既存を停止してから起動した

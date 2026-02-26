# 4-4. フロント ↔ バックエンド結合動作確認

## 概要

フロントエンドのビルドと Lint の通過確認を行った。実 DB を使った E2E 動作確認手順を記録する。

## 確認結果

| チェック項目 | 結果 |
|-------------|------|
| `npm run build` | ✅ 成功 |
| `npm run lint` | ✅ 成功 |
| TypeScript 型チェック | ✅ ビルド内で実行・パス |

## 動作確認手順 (手動)

Docker Desktop の WSL 統合が有効な環境で以下を実行する。

```bash
# 1. PostgreSQL 起動
docker compose up -d

# 2. バックエンド起動
cd backend
go run cmd/server/main.go

# 3. フロントエンド起動
cd frontend
npm run dev
```

### 確認シナリオ

1. `http://localhost:3000` にアクセス → ホームページ表示
2. ユーザー登録 (`/users/new`) → 名前・メールアドレスを入力して登録
3. タスク登録 (`/tasks/new`) → タイトル・担当ユーザー ID を入力して登録
4. タスク一覧 (`/tasks`) → 登録したタスクが表示される
5. タスク詳細 (`/tasks/:id`) → 詳細情報の確認
6. タスク編集 (`/tasks/:id/edit`) → 情報を更新して保存
7. タスク削除 → 確認ダイアログを経て削除・一覧から消える

## 注意事項

- WSL2 環境では Docker Desktop の「WSL Integration」を有効にする必要がある
- フロントエンドからバックエンドへのリクエストには CORS 設定が必要
  - バックエンドの Echo ルーターに CORS ミドルウェアが設定済みであることを確認する

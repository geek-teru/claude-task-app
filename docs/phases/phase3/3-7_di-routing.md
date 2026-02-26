# 3-7. DI・ルーティング結合 (main.go)

## 実施内容
- `backend/adapter/handler/handler.go` を作成（`gen.ServerInterface` を実装する統合 `Handler` struct）
- `backend/infrastructure/router/router.go` を作成（Echo インスタンス生成・ミドルウェア設定・ルート登録）
- `backend/cmd/server/main.go` を作成（DB接続→リポジトリ→ユースケース→ハンドラーの DI 組み立て、サーバー起動）
- `go.sum` に `golang.org/x/time` を追加（Echo middleware の依存）
- `go build ./cmd/server/...` でビルド確認済み
- `go test ./usecase/... ./adapter/handler/...` で既存テスト全件 PASS 確認済み

## 追加・変更ファイル
| ファイル | 変更内容 |
|---------|---------|
| `backend/adapter/handler/handler.go` | TaskHandler + UserHandler を統合する Handler struct を新規作成 |
| `backend/infrastructure/router/router.go` | Echo ルーター設定（Logger/Recover ミドルウェア + ルート登録）を新規作成 |
| `backend/cmd/server/main.go` | DI 組み立てとサーバー起動エントリポイントを新規作成 |
| `backend/go.mod` / `backend/go.sum` | `golang.org/x/time` 依存を追加 |

## 設計ポイント
- `TaskHandler` と `UserHandler` を struct 埋め込みで `Handler` に集約し、`gen.ServerInterface` を満たす
- DI は `cmd/server/main.go` で手動で行い、各層の構築順序は infrastructure → usecase → adapter
- ルーター設定は `infrastructure/router/` に分離し、`gen.RegisterHandlers` を呼ぶことでルート定義を OpenAPI 仕様と一致させる

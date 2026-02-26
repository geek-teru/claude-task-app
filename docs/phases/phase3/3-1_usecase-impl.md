# 3-1. usecase 層実装

## 実施内容
- `backend/usecase/task/usecase.go` を作成（タスク CRUD の usecase 実装）
- `backend/usecase/user/usecase.go` を作成（ユーザー作成・更新の usecase 実装）
- 各 usecase はインターフェース（`Usecase`）と実装（`usecase` struct）を同一パッケージに定義
- `go build ./usecase/...` でビルド確認済み

## 追加・変更ファイル
| ファイル | 変更内容 |
|---------|---------|
| `backend/usecase/task/usecase.go` | タスク usecase インターフェースと実装を新規作成 |
| `backend/usecase/user/usecase.go` | ユーザー usecase インターフェースと実装を新規作成 |

## 設計ポイント
- usecase 層は `domain/repository` インターフェースにのみ依存し、具体的な DB 実装には依存しない
- `Update` は事前に `FindByID` で存在確認を行い、存在しない場合はエラーを返す
- usecase インターフェースを同パッケージで定義することで、handler 層がモックしやすくなる
- CLAUDE.md の構造に従い `usecase/task/` と `usecase/user/` を別パッケージとした

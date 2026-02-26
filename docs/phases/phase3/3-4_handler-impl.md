# 3-4. adapter 層実装 (handler)

## 実施内容
- `backend/domain/entity/errors.go` を作成（センチネルエラー `ErrNotFound` を定義）
- `backend/usecase/task/usecase.go` を更新（FindByID 失敗時に `entity.ErrNotFound` をラップ）
- `backend/usecase/user/usecase.go` を更新（FindByID 失敗時に `entity.ErrNotFound` をラップ）
- `backend/adapter/handler/task_handler.go` を作成（タスク CRUD の Echo ハンドラー実装）
- `backend/adapter/handler/user_handler.go` を作成（ユーザー作成・更新の Echo ハンドラー実装）
- `go build ./adapter/...` でビルド確認済み
- `go test ./usecase/...` で既存のユニットテストが全 PASS であることを確認済み

## 追加・変更ファイル
| ファイル | 変更内容 |
|---------|---------|
| `backend/domain/entity/errors.go` | `ErrNotFound` センチネルエラーを新規追加 |
| `backend/usecase/task/usecase.go` | Get/Update で `entity.ErrNotFound` をラップするよう修正 |
| `backend/usecase/user/usecase.go` | Update で `entity.ErrNotFound` をラップするよう修正 |
| `backend/adapter/handler/task_handler.go` | TaskHandler (ListTasks/CreateTask/GetTask/UpdateTask/DeleteTask) を新規作成 |
| `backend/adapter/handler/user_handler.go` | UserHandler (CreateUser/UpdateUser) を新規作成 |

## 設計ポイント
- `TaskHandler` / `UserHandler` はそれぞれ対応する usecase インターフェースに依存
- エラー種別に応じて HTTP ステータスを変える: 400 (bind error), 404 (ErrNotFound), 500 (other)
- 404 判定は `errors.Is(err, entity.ErrNotFound)` で行い、センチネルエラーでラップ済みのエラーを正確に検出
- `gen.ServerInterface` の実装は `TaskHandler` と `UserHandler` に分割し、DI で組み合わせる（3-7 で実装）
- エンティティ → レスポンスDTO変換は `toTaskResponse` / `toUserResponse` ヘルパーで行う

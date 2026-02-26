# 3-2. usecase 層のテストデータ・モック作成

## 実施内容
- `backend/domain/repository/mock/task_repository.go` を作成（TaskRepository のモック実装）
- `backend/domain/repository/mock/user_repository.go` を作成（UserRepository のモック実装）
- `go build ./domain/repository/mock/...` でビルド確認済み

## 追加・変更ファイル
| ファイル | 変更内容 |
|---------|---------|
| `backend/domain/repository/mock/task_repository.go` | TaskRepository モック実装を新規作成 |
| `backend/domain/repository/mock/user_repository.go` | UserRepository モック実装を新規作成 |

## 設計ポイント
- モックは各メソッドに対応する `Fn` フィールド（関数型）を持つ struct として実装
- テストケースごとに `Fn` フィールドにクロージャを設定することで、柔軟に振る舞いを制御できる
- `backend/domain/repository/mock/` パッケージに配置することで、usecase 層の任意のテストから import して利用可能
- mockgen などの外部ツールへの依存を避けるため、手書きモックを採用

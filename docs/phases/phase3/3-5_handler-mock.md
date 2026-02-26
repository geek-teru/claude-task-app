# 3-5. handler 層のテストデータ・モック作成

## 実施内容
- `backend/usecase/task/mock/usecase.go` を作成（task.Usecase のモック実装）
- `backend/usecase/user/mock/usecase.go` を作成（user.Usecase のモック実装）
- `go build ./usecase/task/mock/... ./usecase/user/mock/...` でビルド確認済み

## 追加・変更ファイル
| ファイル | 変更内容 |
|---------|---------|
| `backend/usecase/task/mock/usecase.go` | TaskUsecase モック実装を新規作成 |
| `backend/usecase/user/mock/usecase.go` | UserUsecase モック実装を新規作成 |

## 設計ポイント
- リポジトリモックと同じパターン（`Fn` フィールドによる関数差し替え）を採用
- `usecase/task/mock/` に配置することで handler テストから import しやすい構造
- 手書きモックにより外部ツール依存なし

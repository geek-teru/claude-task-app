# 2-6. infrastructure 層実装 (GORM リポジトリ, DB接続)

## 実施内容
- `infrastructure/persistence/user_repository.go` に UserRepository の GORM 実装
- `infrastructure/persistence/task_repository.go` に TaskRepository の GORM 実装

## 実装メソッド

### UserRepository
| メソッド | 説明 |
|---------|------|
| Create | ユーザー作成 |
| FindByID | ID で検索 (論理削除済みは除外) |
| Update | ユーザー更新 |

### TaskRepository
| メソッド | 説明 |
|---------|------|
| Create | タスク作成 |
| FindAll | 全件取得 (論理削除済みは除外) |
| FindByID | ID で検索 (論理削除済みは除外) |
| Update | タスク更新 |
| Delete | 論理削除 (deleted_at を現在時刻に更新) |

## 設計ポイント
- entity ↔ model の変換関数 (`toUserModel`/`toUserEntity` 等) で domain と infrastructure を分離
- 論理削除: `deleted_at` がゼロ値のレコードのみを対象にクエリ
- Delete は `deleted_at` を `time.Now()` に更新する論理削除方式

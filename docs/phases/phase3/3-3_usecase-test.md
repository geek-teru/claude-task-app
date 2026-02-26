# 3-3. usecase 層のユニットテスト

## 実施内容
- `backend/usecase/task/usecase_test.go` を作成（タスク usecase の全メソッドをテーブルドリブンテストで網羅）
- `backend/usecase/user/usecase_test.go` を作成（ユーザー usecase の全メソッドをテーブルドリブンテストで網羅）
- `go test ./usecase/...` で全件 PASS 確認済み（12 テスト）

## 追加・変更ファイル
| ファイル | 変更内容 |
|---------|---------|
| `backend/usecase/task/usecase_test.go` | Create/List/Get/Update/Delete の正常系・異常系テストを新規作成 |
| `backend/usecase/user/usecase_test.go` | Create/Update の正常系・異常系テストを新規作成 |

## テストケース一覧

### task usecase
| テスト | ケース |
|--------|--------|
| Create | 正常系: タスクを作成できる |
| Create | 異常系: リポジトリエラー |
| List   | 正常系: タスク一覧を取得できる |
| List   | 正常系: タスクが0件でも空スライスを返す |
| List   | 異常系: リポジトリエラー |
| Get    | 正常系: タスクを取得できる |
| Get    | 異常系: タスクが存在しない |
| Update | 正常系: タスクを更新できる |
| Update | 異常系: 対象タスクが存在しない |
| Update | 異常系: リポジトリの更新でエラー |
| Delete | 正常系: タスクを削除できる |
| Delete | 異常系: リポジトリエラー |

### user usecase
| テスト | ケース |
|--------|--------|
| Create | 正常系: ユーザーを作成できる |
| Create | 異常系: リポジトリエラー |
| Update | 正常系: ユーザーを更新できる |
| Update | 異常系: 対象ユーザーが存在しない |
| Update | 異常系: リポジトリの更新でエラー |

## 設計ポイント
- 各テストケースの `setup` フィールドでモックのクロージャを定義し、テストケースごとに振る舞いを制御
- 動的フィールド（CreatedAt / UpdatedAt / DeletedAt）は `cmpopts.IgnoreFields` で比較から除外
- DBへの依存ゼロのため、純粋なビジネスロジック検証が可能

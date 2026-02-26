# 3-8. E2E テスト

## 実施内容
- `backend/e2e/testdata/users.yml` / `tasks.yml` を作成（E2E テスト用フィクスチャー）
- `backend/e2e/e2e_test.go` を作成（全エンドポイントを通しで検証する E2E テスト）
- `backend/adapter/handler/task_handler.go` を修正（CreateTask の title 必須バリデーション追加）
- `backend/adapter/handler/user_handler.go` を修正（CreateUser の name 必須バリデーション追加）
- `go test ./e2e/...` で全件 PASS 確認済み（13 テスト）

## 追加・変更ファイル
| ファイル | 変更内容 |
|---------|---------|
| `backend/e2e/testdata/users.yml` | E2E テスト用ユーザーフィクスチャーを新規作成 |
| `backend/e2e/testdata/tasks.yml` | E2E テスト用タスクフィクスチャーを新規作成 |
| `backend/e2e/e2e_test.go` | 全エンドポイントの E2E テストを新規作成 |
| `backend/adapter/handler/task_handler.go` | CreateTask の title 空値バリデーション追加 |
| `backend/adapter/handler/user_handler.go` | CreateUser の name 空値バリデーション追加 |

## テストケース一覧
| テスト | ケース |
|--------|--------|
| CreateUser | 正常系: ユーザーを作成できる (201) |
| CreateUser | 異常系: name が空 (400) |
| UpdateUser | 正常系: ユーザーを更新できる (200) |
| UpdateUser | 異常系: 存在しないユーザー (404) |
| CreateTask | 正常系: タスクを作成できる (201) |
| CreateTask | 異常系: title が空 (400) |
| ListTasks  | 正常系: タスク一覧を取得できる (200、件数確認) |
| GetTask    | 正常系: タスクを取得できる (200) |
| GetTask    | 異常系: 存在しないタスク (404) |
| UpdateTask | 正常系: タスクを更新できる (200) |
| UpdateTask | 異常系: 存在しないタスク (404) |
| DeleteTask | 正常系: タスクを削除できる (204) |
| DeleteTask | 異常系: 存在しないタスク ID (204、Delete は存在確認なし) |

## 設計ポイント
- `TestMain` でDB接続・マイグレーション・DI組み立て・`httptest.NewServer` を一度だけ実行
- `loadFixture` で YAML フィクスチャーを投入し、テスト間のデータ状態を制御
- 書き込みテスト（POST/PUT/DELETE）は `t.Cleanup` でテーブルをリセットし、テスト間の干渉を防ぐ
- 実 DB + 全層通過でのテストにより、DI 結合・SQL・HTTPルーティングをまとめて検証

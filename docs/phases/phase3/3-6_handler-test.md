# 3-6. handler 層のテスト

## 実施内容
- `backend/adapter/handler/task_handler_test.go` を作成（タスクハンドラーの全メソッドをテーブルドリブンテストで網羅）
- `backend/adapter/handler/user_handler_test.go` を作成（ユーザーハンドラーの全メソッドをテーブルドリブンテストで網羅）
- `go test ./adapter/handler/...` で全件 PASS 確認済み（22 テスト）

## 追加・変更ファイル
| ファイル | 変更内容 |
|---------|---------|
| `backend/adapter/handler/task_handler_test.go` | CreateTask/ListTasks/GetTask/UpdateTask/DeleteTask のハンドラーテストを新規作成 |
| `backend/adapter/handler/user_handler_test.go` | CreateUser/UpdateUser のハンドラーテストを新規作成 |

## テストケース一覧

### task handler
| テスト | ケース |
|--------|--------|
| CreateTask | 正常系: タスクを作成できる (201) |
| CreateTask | 異常系: 不正なJSON (400) |
| CreateTask | 異常系: usecase エラー (500) |
| ListTasks  | 正常系: タスク一覧を取得できる (200) |
| ListTasks  | 正常系: タスクが0件 (200) |
| ListTasks  | 異常系: usecase エラー (500) |
| GetTask    | 正常系: タスクを取得できる (200) |
| GetTask    | 異常系: タスクが存在しない (404) |
| GetTask    | 異常系: usecase エラー (500) |
| UpdateTask | 正常系: タスクを更新できる (200) |
| UpdateTask | 異常系: タスクが存在しない (404) |
| UpdateTask | 異常系: 不正なJSON (400) |
| DeleteTask | 正常系: タスクを削除できる (204) |
| DeleteTask | 異常系: タスクが存在しない (404) |
| DeleteTask | 異常系: usecase エラー (500) |

### user handler
| テスト | ケース |
|--------|--------|
| CreateUser | 正常系: ユーザーを作成できる (201) |
| CreateUser | 異常系: 不正なJSON (400) |
| CreateUser | 異常系: usecase エラー (500) |
| UpdateUser | 正常系: ユーザーを更新できる (200) |
| UpdateUser | 異常系: ユーザーが存在しない (404) |
| UpdateUser | 異常系: 不正なJSON (400) |

## 設計ポイント
- `httptest.NewRecorder()` と `echo.New().NewContext()` で HTTP コンテキストをインメモリ構築
- usecase はモック (`taskmock.TaskUsecase` / `usermock.UserUsecase`) で差し替え
- DBへの依存ゼロのため高速かつ安定したテスト
- 各テストケースの `setup` でモックの振る舞いを柔軟に定義

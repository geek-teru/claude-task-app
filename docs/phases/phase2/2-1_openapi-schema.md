# 2-1. OpenAPI スキーマ定義

## 実施内容
- `api/openapi.yaml` に OpenAPI 3.0.3 スキーマを定義

## エンドポイント一覧
| メソッド | パス | operationId | 説明 |
|---------|------|-------------|------|
| POST | /api/v1/users | createUser | ユーザー登録 |
| PUT | /api/v1/users/{userId} | updateUser | ユーザー変更 |
| POST | /api/v1/tasks | createTask | タスク登録 |
| GET | /api/v1/tasks | listTasks | タスク一覧 |
| GET | /api/v1/tasks/{taskId} | getTask | タスク詳細 |
| PUT | /api/v1/tasks/{taskId} | updateTask | タスク変更 |
| DELETE | /api/v1/tasks/{taskId} | deleteTask | タスク削除 |

## スキーマ
- User: id, name, email, createdAt, updatedAt
- Task: id, title, description, status (todo/in_progress/done), userId, createdAt, updatedAt
- ErrorResponse: message

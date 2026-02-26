# 2-3. domain 層実装

## 実施内容
- `domain/entity/` にエンティティを定義
- `domain/repository/` にリポジトリインターフェースを定義

## エンティティ
### User
| フィールド | 型 |
|-----------|-----|
| ID | int64 |
| Name | string |
| Email | string |
| CreatedAt | time.Time |
| UpdatedAt | time.Time |

### Task
| フィールド | 型 |
|-----------|-----|
| ID | int64 |
| Title | string |
| Description | string |
| Status | TaskStatus (todo / in_progress / done) |
| UserID | int64 |
| CreatedAt | time.Time |
| UpdatedAt | time.Time |

## リポジトリインターフェース
### UserRepository
- `Create`, `FindByID`, `Update`

### TaskRepository
- `Create`, `FindAll`, `FindByID`, `Update`, `Delete`

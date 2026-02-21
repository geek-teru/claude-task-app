# 2-4. DB スキーマ定義 (Mermaid ER図)

## ER図

```mermaid
erDiagram
    users {
        bigint id PK "AUTO INCREMENT"
        varchar name "NOT NULL"
        varchar email "NOT NULL, UNIQUE"
        timestamp created_at "NOT NULL, DEFAULT CURRENT_TIMESTAMP"
        timestamp updated_at "NOT NULL, DEFAULT CURRENT_TIMESTAMP"
    }

    tasks {
        bigint id PK "AUTO INCREMENT"
        varchar title "NOT NULL"
        text description "NULL"
        varchar status "NOT NULL, DEFAULT 'todo'"
        bigint user_id FK "NOT NULL"
        timestamp created_at "NOT NULL, DEFAULT CURRENT_TIMESTAMP"
        timestamp updated_at "NOT NULL, DEFAULT CURRENT_TIMESTAMP"
    }

    users ||--o{ tasks : "has many"
```

## テーブル定義

### users
| カラム | 型 | 制約 |
|-------|-----|------|
| id | BIGINT | PK, AUTO INCREMENT |
| name | VARCHAR | NOT NULL |
| email | VARCHAR | NOT NULL, UNIQUE |
| created_at | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP |
| updated_at | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP |

### tasks
| カラム | 型 | 制約 |
|-------|-----|------|
| id | BIGINT | PK, AUTO INCREMENT |
| title | VARCHAR | NOT NULL |
| description | TEXT | NULL |
| status | VARCHAR | NOT NULL, DEFAULT 'todo' (todo / in_progress / done) |
| user_id | BIGINT | FK → users.id, NOT NULL |
| created_at | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP |
| updated_at | TIMESTAMP | NOT NULL, DEFAULT CURRENT_TIMESTAMP |

### インデックス
- `tasks.user_id` に外部キーインデックス

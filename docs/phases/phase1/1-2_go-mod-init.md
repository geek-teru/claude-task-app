# 1-2. Go モジュール初期化・依存パッケージ導入

## 実施内容
- `backend/` ディレクトリに Go モジュールを初期化
- モジュール名: `github.com/nanch/claude-task-app/backend`
- 依存パッケージを導入

## 導入パッケージ
| パッケージ | バージョン | 用途 |
|-----------|-----------|------|
| github.com/labstack/echo/v4 | v4.15.0 | Web フレームワーク |
| gorm.io/gorm | v1.31.1 | ORM |
| gorm.io/driver/postgres | v1.6.0 | PostgreSQL ドライバ |
| github.com/oapi-codegen/oapi-codegen/v2 | v2.5.1 | OpenAPI コード生成 |
| github.com/oapi-codegen/runtime | v1.1.2 | 生成コードのランタイム |

## 備考
- Go 1.24.0 を使用

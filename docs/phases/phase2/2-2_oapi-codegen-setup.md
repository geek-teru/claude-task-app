# 2-2. oapi-codegen セットアップ・コード生成

## 実施内容
- `backend/gen/oapi-codegen.cfg.yaml` に設定ファイルを作成
- `backend/gen/generate.go` に `go:generate` ディレクティブを記述
- `go generate ./gen/...` でコード生成を実行

## 生成設定
| 項目 | 値 |
|------|-----|
| パッケージ名 | gen |
| 出力ファイル | server.gen.go |
| Echo サーバーIF | 生成する |
| モデル型 | 生成する |
| 埋め込みスペック | 生成する |

## 生成された ServerInterface
- `ListTasks`, `CreateTask`, `GetTask`, `UpdateTask`, `DeleteTask`
- `CreateUser`, `UpdateUser`

## コマンド
- `cd backend && go generate ./gen/...` — コード再生成

# 1-1. Git リポジトリ初期化・.gitignore 作成

## 実施内容
- `git init` でリポジトリを初期化
- `.gitignore` を作成

## .gitignore の対象
| カテゴリ | 対象 |
|---------|------|
| Go | ビルド成果物 (`*.exe`, `*.out`), `backend/bin/`, `backend/gen/` |
| Next.js | `frontend/node_modules/`, `frontend/.next/`, `frontend/out/` |
| 環境変数 | `.env`, `.env.local`, `.env.*.local` |
| IDE/OS | `.vscode/`, `.idea/`, `.DS_Store`, `Thumbs.db` |
| Docker | `docker-compose.override.yml` |

## 備考
- `backend/gen/` は oapi-codegen の自動生成コードのため Git 管理外とした

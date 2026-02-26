# 2-8. テストデータ作成 (YAML Fixture)

## 実施内容
- `testdata/users/` と `testdata/tasks/` に YAML Fixture ファイルを作成
- 正常系・異常系（論理削除済み等）のテストデータを定義

## 追加ファイル
| ファイル | 内容 |
|---------|------|
| `testdata/users/users.yml` | ユーザー Fixture (3件) |
| `testdata/tasks/users.yml` | タスクテスト用ユーザー Fixture (1件) |
| `testdata/tasks/tasks.yml` | タスク Fixture (3件) |

## Fixture データ

### testdata/users/users.yml
| id | name | email | 用途 |
|----|------|-------|------|
| 1 | テストユーザー | test@example.com | 正常系テスト |
| 2 | 検索テストユーザー | search@example.com | 検索テスト |
| 3 | 削除済みユーザー | deleted@example.com | 論理削除テスト (deleted_at 設定済み) |

### testdata/tasks/tasks.yml
| id | title | status | user_id | 用途 |
|----|-------|--------|---------|------|
| 1 | タスク1 | todo | 1 | 正常系テスト |
| 2 | タスク2 | in_progress | 1 | 一覧取得テスト |
| 3 | 削除済みタスク | done | 1 | 論理削除テスト (deleted_at 設定済み) |

## 設計ポイント
- tasks ディレクトリには users.yml も配置し、外部キー制約を満たす
- 論理削除済みデータを含めることで、FindByID / FindAll の除外ロジックをテスト可能

# 2-7. テスト基盤構築

## 実施内容
- go-testfixtures、go-cmp を導入
- `repository_test.go` に TestMain、loadFixture、truncateTable を実装
- テスト用 DB 接続はパッケージ変数 `testDB` で保持し、TestMain で1回だけ接続

## 追加・変更ファイル
| ファイル | 変更内容 |
|---------|---------|
| `repository_test.go` | 新規作成: TestMain, setUp, loadFixture, truncateTable, getTestEnv |
| `go.mod` / `go.sum` | go-testfixtures/v3, go-cmp 追加 |

## 関数一覧
| 関数 | 説明 |
|------|------|
| `TestMain` | テスト開始前に DB 接続を1回だけ実行 |
| `setUp` | DSN を組み立てて GORM で PostgreSQL に接続 |
| `loadFixture` | go-testfixtures で YAML Fixture を DB に投入 |
| `truncateTable` | 指定テーブルを TRUNCATE (RESTART IDENTITY CASCADE) |
| `getTestEnv` | 環境変数を取得 (デフォルト値付き) |

## 導入パッケージ
| パッケージ | 用途 |
|-----------|------|
| `github.com/go-testfixtures/testfixtures/v3` | YAML ファイルからテストデータを DB に投入 |
| `github.com/google/go-cmp` | 構造体の差分比較 (cmpopts.IgnoreFields で動的フィールド除外) |

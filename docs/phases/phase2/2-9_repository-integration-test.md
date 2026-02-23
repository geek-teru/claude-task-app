# 2-9. repository 層の結合テスト実装

## 実施内容
- UserRepository / TaskRepository の結合テストをテーブルドリブンテストで実装
- 正常系・異常系を網羅し、全 13 サブテスト PASS を確認

## テストケース

### UserRepository (6サブテスト)

#### TestUserRepository_Create
| ケース | 内容 | 種別 |
|--------|------|------|
| OK | ユーザー作成、ID 自動採番・名前・メールの検証 | 正常系 |
| ERROR: email 重複 | 既存メールで作成しエラーを確認 | 異常系 |

#### TestUserRepository_FindByID
| ケース | 内容 | 種別 |
|--------|------|------|
| OK | Fixture ユーザーを ID で取得 | 正常系 |
| ERROR: 存在しないID | 存在しない ID でエラーを確認 | 異常系 |
| ERROR: 論理削除済み | 論理削除済みユーザーでエラーを確認 | 異常系 |

#### TestUserRepository_Update
| ケース | 内容 | 種別 |
|--------|------|------|
| OK | 名前を更新し、DB 再取得で反映を確認 | 正常系 |

### TaskRepository (7サブテスト)

#### TestTaskRepository_Create
| ケース | 内容 | 種別 |
|--------|------|------|
| OK | タスク作成、ID 自動採番・タイトル・ステータスの検証 | 正常系 |
| ERROR: 存在しないユーザーID | 外部キー制約違反でエラーを確認 | 異常系 |

#### TestTaskRepository_FindAll
| 検証項目 | 内容 |
|---------|------|
| 件数 | 論理削除済みを除外して 2件 |
| 内容 | タイトルが「タスク1」「タスク2」であること |

#### TestTaskRepository_FindByID
| ケース | 内容 | 種別 |
|--------|------|------|
| OK | Fixture タスクを ID で取得 | 正常系 |
| ERROR: 存在しないID | 存在しない ID でエラーを確認 | 異常系 |
| ERROR: 論理削除済み | 論理削除済みタスクでエラーを確認 | 異常系 |

#### TestTaskRepository_Update
| ケース | 内容 | 種別 |
|--------|------|------|
| OK | タイトル・ステータスを更新し、DB 再取得で反映を確認 | 正常系 |

#### TestTaskRepository_Delete
| ケース | 内容 | 種別 |
|--------|------|------|
| OK | 論理削除後に FindByID でエラーが返ることを確認 | 正常系 |

## テスト設計パターン
- **テーブルドリブンテスト**: `tests := []struct{...}` + `t.Run` で正常系・異常系を列挙
- **インラインデータ**: テストデータは各テストケースに直接定義（グローバル変数不使用）
- **setup per subtest**: DB を変更するテスト (Create/Update/Delete) は各サブテストの先頭で Fixture 投入
- **single loadFixture**: 読み取り専用テスト (FindAll/FindByID) はテスト関数の先頭で1回だけ投入
- **go-cmp 比較**: `cmpopts.IgnoreFields` で CreatedAt, UpdatedAt, DeletedAt を除外

## テスト実行コマンド
```
cd backend && go test ./infrastructure/persistence/ -v
```

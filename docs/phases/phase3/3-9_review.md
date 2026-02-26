# 3-9. コード・CLAUDE.md レビュー

## レビュー結果サマリ

| 分類 | 件数 | 重要度 |
|------|------|--------|
| 修正済み | 3件 | 高 |
| 将来対応 | 5件 | 中/低 |
| 問題なし | - | - |

## 修正済み（重要度「高」）

### 1. usecase 層エラーハンドリングの一貫性
- **問題**: `Get` / `Update` で全エラーを `entity.ErrNotFound` で無差別ラップ → 接続エラー等も 404 になるバグ
- **修正**:
  - `persistence/task_repository.go`, `user_repository.go` で `gorm.ErrRecordNotFound` の場合のみ `entity.ErrNotFound` をラップして返すよう修正
  - `usecase/task/usecase.go`, `usecase/user/usecase.go` で `entity.ErrNotFound` のハードコードを廃止し、リポジトリエラーをそのまま伝播させるよう修正

### 2. CLAUDE.md プロジェクト構造の更新
- **問題**: `adapter/presenter/` が未実装なのに記載あり、`domain/repository/mock/`, `usecase/task/mock/` 等の新規ディレクトリが未記載
- **修正**: プロジェクト構造を実際のディレクトリ構成に合わせて更新

### 3. DeleteTask ハンドラーの無効コード
- **問題**: `usecase.Delete` は存在確認を行わないため `ErrNotFound` が返ることがなく、handler の `ErrNotFound` チェックは無効コード
- **修正**: 削除操作は冪等と定義し、`ErrNotFound` 判定を削除。CLAUDE.md にも DELETE の冪等性ポリシーを追記

## 将来対応（重要度「中」/「低」）

### 1. ユーザーID バリデーション（中）
- タスクの `userId` に 0 や負数を受け付ける
- 外部キー制約がない設計なので、存在しない userId でもタスク作成が可能
- 対応: usecase 層での userId バリデーション追加 または UserRepository.Exists() の追加

### 2. E2E テストの追加ケース（中）
- メールアドレス重複エラーのテストがない
- タスク更新のバリデーション（title 空）テストがない
- これらは E2E テスト拡充時に対応

### 3. 入力値の最大長バリデーション（中）
- `title` / `name` / `description` に最大長の制限がない
- Phase 4 (フロントエンド実装) 前後で対応を検討

### 4. エラーメッセージの情報開示（低）
- 500 エラー時に `err.Error()` をそのままクライアントに返している
- GORM のエラーメッセージ等が漏洩する可能性
- 将来的にエラーをログに記録し、レスポンスは汎用メッセージに変更

### 5. DRY: エラーハンドリングの重複（低）
- handler の `ErrNotFound` チェックが各ハンドラーに重複
- ヘルパー関数に抽出することで保守性向上（ただし現状の規模では許容範囲）

# 4-5. コード・CLAUDE.md レビュー

## レビュー実施内容

Phase 4 で実装したフロントエンドコードのレビューを実施し、指摘事項を修正した。

## 指摘・対応

| # | 指摘内容 | 重要度 | 対応 |
|---|---------|--------|------|
| 1 | Server Component の fetch エラーハンドリング欠如 | MAJOR | `try/catch` でエラーを捕捉してユーザーフレンドリーなメッセージを表示 |
| 2 | ステータスラベル・カラーマップがページ間で重複 | MINOR | `src/lib/constants.ts` に共通定数として抽出 |
| 3 | `Task` 型の `description` がオプショナルだがバックエンドは NOT NULL | MINOR | `description: string` (必須) に修正 |
| 4 | `UserEditForm` という誤解を招くインポートエイリアス | MINOR | `UserForm` に統一 |

## 既知の制限事項

- ユーザー編集ページ (`/users/:id/edit`) は既存データのプリフィルが不可
  - バックエンドに `GET /api/v1/users/:id` エンドポイントが存在しないため
  - Phase 5 以降でのバックエンド拡張時に対応を検討

## CLAUDE.md 規約準拠確認

- ✅ 関数コンポーネントを使用
- ✅ Server Components をデフォルト、必要時のみ `"use client"`
- ✅ 変数名・関数名は camelCase、コンポーネント名は PascalCase
- ✅ `npm run build` 成功 (TypeScript チェック通過)
- ✅ `npm run lint` 成功

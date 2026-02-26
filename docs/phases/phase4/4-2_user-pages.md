# 4-2. ユーザー管理画面 (登録・編集)

## 概要

ユーザー登録フォームと編集フォームの画面を実装した。

## 成果物

### `frontend/src/components/UserForm.tsx`

- 登録 / 編集を兼ねる Client Component フォーム
- `userId` props の有無で登録 / 編集を切り替え
- 送信成功時はトップページへ遷移
- エラー時はフォーム上にメッセージ表示

### `frontend/src/app/users/new/page.tsx`

- ユーザー登録ページ (`/users/new`)
- `UserForm` をレンダリング

### `frontend/src/app/users/[id]/edit/page.tsx`

- ユーザー編集ページ (`/users/:id/edit`)
- バックエンドに `GET /api/v1/users/:id` が存在しないためプリフィルなし

## 制限事項

- バックエンドの API 仕様にユーザー取得 (`GET /api/v1/users/:id`) がないため、編集フォームは既存値のプリフィルができない。ユーザーが再入力する必要がある。

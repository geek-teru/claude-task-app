# 4-3. タスク管理画面 (一覧・登録・詳細・編集・削除)

## 概要

タスク管理に必要な全画面と共通レイアウト・ホームページを実装した。

## 成果物

### コンポーネント

| ファイル | 説明 |
|---------|------|
| `src/components/TaskForm.tsx` | タスク登録/編集フォーム (Client Component) |
| `src/components/DeleteTaskButton.tsx` | タスク削除ボタン (確認ステップ付き、Client Component) |

### ページ

| ファイル | URL | 説明 |
|---------|-----|------|
| `src/app/page.tsx` | `/` | ホームページ (ナビゲーションリンク) |
| `src/app/layout.tsx` | 全ページ共通 | ナビゲーションバー付きレイアウト |
| `src/app/tasks/page.tsx` | `/tasks` | タスク一覧 (Server Component) |
| `src/app/tasks/new/page.tsx` | `/tasks/new` | タスク登録 |
| `src/app/tasks/[id]/page.tsx` | `/tasks/:id` | タスク詳細 (Server Component) |
| `src/app/tasks/[id]/edit/page.tsx` | `/tasks/:id/edit` | タスク編集 (Server Component でデータ取得 → TaskForm へ渡す) |

## 設計ポイント

- データ取得は Server Component で行い、Client Component は最小限に留めた
- 削除ボタンは確認ステップを設けて誤操作を防止
- ステータスは日本語ラベルとカラーバッジで視覚的に表示
- タスク編集はバックエンドの `GET /api/v1/tasks/:id` でデータを取得してプリフィル

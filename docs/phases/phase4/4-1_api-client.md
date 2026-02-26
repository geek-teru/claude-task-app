# 4-1. API クライアント作成

## 概要

フロントエンドからバックエンド REST API を呼び出すための型定義と fetch ラッパーを実装した。

## 成果物

### `frontend/src/types/index.ts`

OpenAPI スキーマに対応する TypeScript 型を定義。

- `User`, `CreateUserRequest`, `UpdateUserRequest`
- `Task`, `TaskStatus`, `CreateTaskRequest`, `UpdateTaskRequest`
- `ErrorResponse`

### `frontend/src/lib/api.ts`

共通 fetch ラッパー (`request<T>`) と各エンドポイントの呼び出し関数。

- `ApiError` クラス: HTTP エラー時にステータスコードとメッセージを保持
- `userApi.create`, `userApi.update`
- `taskApi.list`, `taskApi.get`, `taskApi.create`, `taskApi.update`, `taskApi.delete`

## 設計ポイント

- `NEXT_PUBLIC_API_BASE_URL` 環境変数でベース URL を切り替え可能 (デフォルト: `http://localhost:8080`)
- 204 No Content レスポンスはボディ解析をスキップ
- エラーレスポンスは `ApiError` としてスローし、呼び出し元でハンドリング

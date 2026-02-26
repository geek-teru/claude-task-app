import Link from "next/link";

export default function Home() {
  return (
    <div className="flex min-h-screen items-center justify-center bg-gray-50">
      <div className="text-center space-y-6">
        <h1 className="text-3xl font-bold text-gray-900">Claude Task App</h1>
        <p className="text-gray-500">シンプルなタスク管理アプリ</p>

        <div className="flex flex-col sm:flex-row gap-4 justify-center">
          <Link
            href="/tasks"
            className="rounded-md bg-blue-600 px-6 py-3 text-sm font-medium text-white hover:bg-blue-700"
          >
            タスク管理
          </Link>
          <Link
            href="/users/new"
            className="rounded-md border border-gray-300 px-6 py-3 text-sm font-medium text-gray-700 hover:bg-gray-50"
          >
            ユーザー登録
          </Link>
        </div>
      </div>
    </div>
  );
}

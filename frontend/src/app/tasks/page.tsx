import Link from "next/link";
import type { Task } from "@/types";
import { TASK_STATUS_LABEL, TASK_STATUS_COLOR } from "@/lib/constants";
import DeleteTaskButton from "@/components/DeleteTaskButton";

async function fetchTasks(): Promise<Task[]> {
  const baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL ?? "http://localhost:8080";
  const res = await fetch(`${baseUrl}/api/v1/tasks`, { cache: "no-store" });
  if (!res.ok) throw new Error("タスクの取得に失敗しました");
  return res.json() as Promise<Task[]>;
}

export default async function TaskListPage() {
  let tasks: Task[];
  try {
    tasks = await fetchTasks();
  } catch {
    return (
      <div className="p-6">
        <h1 className="text-2xl font-bold mb-6">タスク一覧</h1>
        <p className="text-red-600">タスクの取得に失敗しました。バックエンドサーバーが起動しているか確認してください。</p>
      </div>
    );
  }

  return (
    <div className="p-6">
      <div className="flex items-center justify-between mb-6">
        <h1 className="text-2xl font-bold">タスク一覧</h1>
        <Link
          href="/tasks/new"
          className="rounded-md bg-blue-600 px-4 py-2 text-sm font-medium text-white hover:bg-blue-700"
        >
          + 新規タスク
        </Link>
      </div>

      {tasks.length === 0 ? (
        <p className="text-gray-500">タスクがありません</p>
      ) : (
        <div className="space-y-3">
          {tasks.map((task) => (
            <div
              key={task.id}
              className="flex items-center justify-between rounded-lg border border-gray-200 bg-white p-4 shadow-sm"
            >
              <div className="flex items-center gap-3">
                <span
                  className={`rounded-full px-2 py-0.5 text-xs font-medium ${TASK_STATUS_COLOR[task.status] ?? "bg-gray-100 text-gray-700"}`}
                >
                  {TASK_STATUS_LABEL[task.status] ?? task.status}
                </span>
                <Link
                  href={`/tasks/${task.id}`}
                  className="font-medium text-gray-900 hover:text-blue-600"
                >
                  {task.title}
                </Link>
              </div>

              <div className="flex items-center gap-2">
                <Link
                  href={`/tasks/${task.id}/edit`}
                  className="rounded-md border border-gray-300 px-3 py-1 text-xs font-medium text-gray-700 hover:bg-gray-50"
                >
                  編集
                </Link>
                <DeleteTaskButton taskId={task.id} />
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

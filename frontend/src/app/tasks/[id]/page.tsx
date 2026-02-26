import { notFound } from "next/navigation";
import Link from "next/link";
import type { Task } from "@/types";
import DeleteTaskButton from "@/components/DeleteTaskButton";

type Props = {
  params: Promise<{ id: string }>;
};

async function fetchTask(id: string): Promise<Task | null> {
  const baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL ?? "http://localhost:8080";
  const res = await fetch(`${baseUrl}/api/v1/tasks/${id}`, { cache: "no-store" });
  if (res.status === 404) return null;
  if (!res.ok) throw new Error("タスクの取得に失敗しました");
  return res.json() as Promise<Task>;
}

const statusLabel: Record<string, string> = {
  todo: "未着手",
  in_progress: "進行中",
  done: "完了",
};

const statusColor: Record<string, string> = {
  todo: "bg-gray-100 text-gray-700",
  in_progress: "bg-blue-100 text-blue-700",
  done: "bg-green-100 text-green-700",
};

export default async function TaskDetailPage({ params }: Props) {
  const { id } = await params;
  const task = await fetchTask(id);

  if (!task) notFound();

  return (
    <div className="p-6 max-w-2xl">
      <div className="flex items-center justify-between mb-6">
        <h1 className="text-2xl font-bold">タスク詳細</h1>
        <div className="flex gap-2">
          <Link
            href={`/tasks/${task.id}/edit`}
            className="rounded-md border border-gray-300 px-3 py-1.5 text-sm font-medium text-gray-700 hover:bg-gray-50"
          >
            編集
          </Link>
          <DeleteTaskButton taskId={task.id} redirectTo="/tasks" />
        </div>
      </div>

      <div className="rounded-lg border border-gray-200 bg-white p-6 space-y-4 shadow-sm">
        <div>
          <p className="text-xs font-medium text-gray-500 uppercase tracking-wide">タイトル</p>
          <p className="mt-1 text-lg font-semibold text-gray-900">{task.title}</p>
        </div>

        {task.description && (
          <div>
            <p className="text-xs font-medium text-gray-500 uppercase tracking-wide">説明</p>
            <p className="mt-1 text-gray-700 whitespace-pre-wrap">{task.description}</p>
          </div>
        )}

        <div className="flex gap-8">
          <div>
            <p className="text-xs font-medium text-gray-500 uppercase tracking-wide">ステータス</p>
            <span
              className={`mt-1 inline-block rounded-full px-2 py-0.5 text-xs font-medium ${statusColor[task.status] ?? "bg-gray-100 text-gray-700"}`}
            >
              {statusLabel[task.status] ?? task.status}
            </span>
          </div>

          <div>
            <p className="text-xs font-medium text-gray-500 uppercase tracking-wide">担当ユーザー ID</p>
            <p className="mt-1 text-gray-700">{task.userId}</p>
          </div>
        </div>

        <div className="flex gap-8 text-sm text-gray-500">
          <div>
            <span className="font-medium">作成日時: </span>
            {new Date(task.createdAt).toLocaleString("ja-JP")}
          </div>
          <div>
            <span className="font-medium">更新日時: </span>
            {new Date(task.updatedAt).toLocaleString("ja-JP")}
          </div>
        </div>
      </div>

      <div className="mt-4">
        <Link href="/tasks" className="text-sm text-blue-600 hover:underline">
          ← タスク一覧に戻る
        </Link>
      </div>
    </div>
  );
}

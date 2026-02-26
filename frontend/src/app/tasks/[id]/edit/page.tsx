import { notFound } from "next/navigation";
import Link from "next/link";
import TaskForm from "@/components/TaskForm";
import type { Task } from "@/types";

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

export default async function EditTaskPage({ params }: Props) {
  const { id } = await params;

  let task: Task | null;
  try {
    task = await fetchTask(id);
  } catch {
    return (
      <div className="p-6">
        <p className="text-red-600">タスクの取得に失敗しました。</p>
        <Link href="/tasks" className="mt-2 inline-block text-sm text-blue-600 hover:underline">
          ← タスク一覧に戻る
        </Link>
      </div>
    );
  }

  if (!task) notFound();

  return (
    <div className="p-6">
      <h1 className="text-2xl font-bold mb-6">タスク編集</h1>
      <TaskForm task={task} />
    </div>
  );
}

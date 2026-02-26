import TaskForm from "@/components/TaskForm";

export default function NewTaskPage() {
  return (
    <div className="p-6">
      <h1 className="text-2xl font-bold mb-6">タスク登録</h1>
      <TaskForm />
    </div>
  );
}

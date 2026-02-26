import UserForm from "@/components/UserForm";

export default function NewUserPage() {
  return (
    <div className="p-6">
      <h1 className="text-2xl font-bold mb-6">ユーザー登録</h1>
      <UserForm />
    </div>
  );
}

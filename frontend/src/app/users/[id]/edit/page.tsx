import UserEditForm from "@/components/UserForm";

type Props = {
  params: Promise<{ id: string }>;
};

export default async function EditUserPage({ params }: Props) {
  const { id } = await params;

  return (
    <div className="p-6">
      <h1 className="text-2xl font-bold mb-6">ユーザー編集 (ID: {id})</h1>
      {/* バックエンドに GET /users/:id エンドポイントがないためプリフィルなし */}
      <UserEditForm userId={Number(id)} />
    </div>
  );
}

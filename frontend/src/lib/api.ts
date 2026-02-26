import type {
  User,
  Task,
  CreateUserRequest,
  UpdateUserRequest,
  CreateTaskRequest,
  UpdateTaskRequest,
} from "@/types";

const BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL ?? "http://localhost:8080";

async function request<T>(path: string, init?: RequestInit): Promise<T> {
  const res = await fetch(`${BASE_URL}${path}`, {
    headers: { "Content-Type": "application/json" },
    ...init,
  });

  if (!res.ok) {
    const body = await res.json().catch(() => ({ message: res.statusText }));
    throw new ApiError(res.status, body.message ?? res.statusText);
  }

  // 204 No Content はボディなし
  if (res.status === 204) {
    return undefined as T;
  }

  return res.json() as Promise<T>;
}

export class ApiError extends Error {
  constructor(
    public readonly status: number,
    message: string,
  ) {
    super(message);
    this.name = "ApiError";
  }
}

// ========== Users ==========

export const userApi = {
  create: (body: CreateUserRequest): Promise<User> =>
    request<User>("/api/v1/users", {
      method: "POST",
      body: JSON.stringify(body),
    }),

  update: (id: number, body: UpdateUserRequest): Promise<User> =>
    request<User>(`/api/v1/users/${id}`, {
      method: "PUT",
      body: JSON.stringify(body),
    }),
};

// ========== Tasks ==========

export const taskApi = {
  list: (): Promise<Task[]> => request<Task[]>("/api/v1/tasks"),

  get: (id: number): Promise<Task> => request<Task>(`/api/v1/tasks/${id}`),

  create: (body: CreateTaskRequest): Promise<Task> =>
    request<Task>("/api/v1/tasks", {
      method: "POST",
      body: JSON.stringify(body),
    }),

  update: (id: number, body: UpdateTaskRequest): Promise<Task> =>
    request<Task>(`/api/v1/tasks/${id}`, {
      method: "PUT",
      body: JSON.stringify(body),
    }),

  delete: (id: number): Promise<void> =>
    request<void>(`/api/v1/tasks/${id}`, { method: "DELETE" }),
};

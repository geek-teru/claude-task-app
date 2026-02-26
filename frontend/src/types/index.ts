// ========== User ==========

export type User = {
  id: number;
  name: string;
  email: string;
  createdAt: string;
  updatedAt: string;
};

export type CreateUserRequest = {
  name: string;
  email: string;
};

export type UpdateUserRequest = {
  name: string;
  email: string;
};

// ========== Task ==========

export type TaskStatus = "todo" | "in_progress" | "done";

export type Task = {
  id: number;
  title: string;
  description?: string;
  status: TaskStatus;
  userId: number;
  createdAt: string;
  updatedAt: string;
};

export type CreateTaskRequest = {
  title: string;
  description?: string;
  status?: TaskStatus;
  userId: number;
};

export type UpdateTaskRequest = {
  title: string;
  description?: string;
  status?: TaskStatus;
  userId: number;
};

// ========== Error ==========

export type ErrorResponse = {
  message: string;
};

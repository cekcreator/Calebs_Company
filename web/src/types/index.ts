export interface User {
  id: string;
  email: string;
  name: string;
  role: "admin" | "user";
  createdAt: string;
}

export interface ApiResponse<T> {
  data: T;
  message?: string;
}

export interface WsMessage {
  type: string;
  payload: unknown;
  timestamp: string;
}

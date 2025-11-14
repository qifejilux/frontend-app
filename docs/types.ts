// types.ts

export interface User {
  id: string;
  name: string;
  email: string;
  role: 'admin' | 'user' | 'guest';
  createdAt: Date;
  updatedAt: Date;
}

export type ApiResponse<T> = {
  data: T;
  success: boolean;
  message?: string;
  timestamp: string;
};

export type PaginatedResponse<T> = ApiResponse<T[]> & {
  total: number;
  page: number;
  limit: number;
};

export interface ErrorResponse {
  error: string;
  statusCode: number;
  message: string;
  timestamp: string;
}

export type Theme = 'light' | 'dark' | 'system';

export interface AppConfig {
  apiBaseUrl: string;
  theme: Theme;
  featureFlags: {
    enableAnalytics: boolean;
    enableNotifications: boolean;
  };
}

export type Nullable<T> = T | null;

export type Optional<T> = T | undefined;
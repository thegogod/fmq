export interface Toast {
  readonly id: string;
  readonly type: 'info' | 'success' | 'warn' | 'error';
  readonly message: string;
  readonly timeout?: number;
}

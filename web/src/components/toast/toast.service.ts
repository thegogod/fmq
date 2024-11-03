import { Overlay } from '@angular/cdk/overlay';
import { ComponentPortal } from '@angular/cdk/portal';
import { ComponentRef, Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import * as uuid from 'uuid';

import { ToastsComponent } from './toasts.component';
import { Toast } from './toast';

export interface ToastOptions {
  readonly id?: string;
  readonly timeout?: number;
}

@Injectable()
export class ToastService {
  readonly $stack = new BehaviorSubject<Toast[]>([ ]);

  private _ref?: ComponentRef<ToastsComponent>;

  constructor(private readonly _overlay: Overlay) { }

  info(message: string, options?: ToastOptions) {
    return this.toast('info', message, options);
  }

  warn(message: string, options?: ToastOptions) {
    return this.toast('warn', message, options);
  }

  success(message: string, options?: ToastOptions) {
    return this.toast('success', message, options);
  }

  error(message: string, options?: ToastOptions) {
    return this.toast('error', message, options);
  }

  toast(type: 'info' | 'warn' | 'success' | 'error', message: string, options?: ToastOptions) {
    if (!this._ref) {
      const overlayRef = this._overlay.create({
        panelClass: 'justify-end',
        positionStrategy: this._overlay.position()
          .global()
          .bottom('15px')
          .right('15px')
      });

      const ref = overlayRef.attach(new ComponentPortal(ToastsComponent));
      this._ref = ref;
    }

    const stack = this.$stack.value;

    if (stack.length === 5) {
      stack.shift();
    }

    stack.push({
      id: uuid.v4(),
      type,
      message,
      ...options
    });

    this.$stack.next(stack);
  }

  dismiss(id: string) {
    const stack = this.$stack.value;
    const i = stack.findIndex(v => v.id === id);

    if (i === -1) return;

    stack.splice(i, 1);
    this.$stack.next(stack);
  }
}

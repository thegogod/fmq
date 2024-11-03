import { HttpErrorResponse, HttpHandler, HttpInterceptor, HttpRequest } from '@angular/common/http';
import { Injectable, isDevMode } from '@angular/core';
import { Router } from '@angular/router';
import { catchError, throwError } from 'rxjs';

import { ToastService } from '../components/toast';

@Injectable()
export class ErrorHttpInterceptor implements HttpInterceptor {
  constructor(
    private readonly _router: Router,
    private readonly _toasts: ToastService
  ) { }

  intercept(req: HttpRequest<any>, next: HttpHandler) {
    return next.handle(req).pipe(
      catchError((err: HttpErrorResponse) => {
        if (err.status === 401) {
          localStorage.removeItem('social-ai--token');
          this._router.navigate(['/']);
        }

        if (isDevMode()) {
          this._toasts.error(err.error, {
            timeout: 3000
          });
        }

        return throwError(() => err);
      })
    );
  }
}

import { HttpHandler, HttpInterceptor, HttpRequest } from '@angular/common/http';
import { Injectable } from '@angular/core';

import { environment } from '../environments/environment';

@Injectable()
export class AuthHttpInterceptor implements HttpInterceptor {
  intercept(req: HttpRequest<any>, next: HttpHandler) {
    const token = localStorage.getItem('fmq--token');

    req = req.clone({
      url: `${environment.api.baseUrl}${req.url}`
    });

    if (token) {
      req = req.clone({
        setHeaders: {
          Authorization: `Basic ${token}`
        }
      });
    }

    return next.handle(req);
  }
}

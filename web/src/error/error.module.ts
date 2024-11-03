import { HTTP_INTERCEPTORS } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';

import { ToastModule } from '../components/toast';

import { ErrorHttpInterceptor } from './error-http.interceptor';

@NgModule({
  imports: [
    RouterModule,
    ToastModule
  ],
  providers: [
    {
      provide: HTTP_INTERCEPTORS,
      useClass: ErrorHttpInterceptor,
      multi: true
    }
  ]
})
export class ErrorModule { }

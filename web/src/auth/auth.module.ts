import { HTTP_INTERCEPTORS } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';

import { AuthHttpInterceptor } from './auth-http.interceptor';
import { PublicGuard } from './public.guard';
import { PrivateGuard } from './private.guard';

@NgModule({
  imports: [RouterModule],
  providers: [
    PublicGuard,
    PrivateGuard,
    {
      provide: HTTP_INTERCEPTORS,
      useClass: AuthHttpInterceptor,
      multi: true
    }
  ]
})
export class AuthModule { }

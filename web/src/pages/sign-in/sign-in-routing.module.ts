import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { AuthModule, PublicGuard } from '../../auth';

import { SignInComponent } from './sign-in.component';

const routes: Routes = [
  {
    path: '',
    canActivate: [PublicGuard],
    component: SignInComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule, AuthModule]
})
export class SignInRoutingModule { }

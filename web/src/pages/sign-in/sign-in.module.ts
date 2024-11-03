import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';

import { IconModule } from '../../components/icon';
import { ToastModule } from '../../components/toast';

import { SignInRoutingModule } from './sign-in-routing.module';
import { SignInComponent } from './sign-in.component';

@NgModule({
  declarations: [SignInComponent],
  imports: [
    CommonModule,
    ReactiveFormsModule,

    SignInRoutingModule,
    IconModule,
    ToastModule
  ]
})
export class SignInModule { }

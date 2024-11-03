import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { OverlayModule } from '@angular/cdk/overlay';

import { IconModule } from '../icon';

import { ToastComponent } from './toast.component';
import { ToastService } from './toast.service';
import { ToastsComponent } from './toasts.component';

@NgModule({
  declarations: [ToastComponent, ToastsComponent],
  providers: [ToastService],
  exports: [ToastComponent, ToastsComponent],
  imports: [CommonModule, OverlayModule, IconModule]
})
export class ToastModule { }

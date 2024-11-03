import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { OverlayModule } from '@angular/cdk/overlay';

import { PopoverComponent } from './popover.component';
import { PopoverDirective } from './popover.directive';

@NgModule({
  declarations: [PopoverComponent, PopoverDirective],
  exports: [PopoverDirective],
  imports: [CommonModule, OverlayModule]
})
export class PopoverModule { }

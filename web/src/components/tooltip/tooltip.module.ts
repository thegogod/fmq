import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { OverlayModule } from '@angular/cdk/overlay';

import { TooltipComponent } from './tooltip.component';
import { TooltipDirective } from './tooltip.directive';

@NgModule({
  declarations: [TooltipComponent, TooltipDirective],
  exports: [TooltipDirective],
  imports: [CommonModule, OverlayModule]
})
export class TooltipModule { }

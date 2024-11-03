import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';

import { IconModule } from '../../components/icon';
import { TooltipModule } from '../../components/tooltip';
import { DevtoolsModule } from '../../devtools';
import { AvatarModule } from '../../components/avatar';

import { RootRoutingModule } from './root-routing.module';
import { RootComponent } from './root.component';

@NgModule({
  declarations: [RootComponent],
  imports: [
    CommonModule,

    RootRoutingModule,
    IconModule,
    TooltipModule,
    DevtoolsModule,
    AvatarModule
  ]
})
export class RootModule { }

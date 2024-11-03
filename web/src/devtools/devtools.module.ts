import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';

import { PortalModule } from '@angular/cdk/portal';
import { DialogModule } from '@angular/cdk/dialog';

import { IconModule } from '../components/icon';
import { JsonModule } from '../components/json';
import { ToastModule } from '../components/toast';

import { DevtoolsComponent } from './devtools.component';
import { DevtoolsService } from './devtools.service';
import { DevtoolsStateComponent } from './devtools-state';
import { DevtoolsInfoComponent } from './devtools-info';

@NgModule({
  declarations: [
    DevtoolsComponent,
    DevtoolsStateComponent,
    DevtoolsInfoComponent
  ],
  imports: [
    CommonModule,
    RouterModule,

    DialogModule,
    PortalModule,

    IconModule,
    JsonModule,
    ToastModule
  ],
  exports: [DevtoolsComponent],
  providers: [DevtoolsService]
})
export class DevtoolsModule { }

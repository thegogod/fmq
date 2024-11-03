import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ClipboardModule } from '@angular/cdk/clipboard';

import { JsonComponent } from './json.component';
import { JsonObjectComponent } from './json-object';
import { JsonStringComponent } from './json-string';

@NgModule({
  declarations: [
    JsonComponent,
    JsonObjectComponent,
    JsonStringComponent
  ],
  exports: [
    JsonComponent,
    JsonObjectComponent,
    JsonStringComponent
  ],
  imports: [
    CommonModule,
    ClipboardModule
  ]
})
export class JsonModule { }

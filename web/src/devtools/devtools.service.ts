import { Dialog } from '@angular/cdk/dialog';
import { Injectable } from '@angular/core';

import { DevtoolsComponent, DevtoolsTabKey } from './devtools.component';

@Injectable()
export class DevtoolsService {
  constructor(private readonly _dialog: Dialog) { }

  open(tab: DevtoolsTabKey = 'info') {
    if (this._dialog.openDialogs.length > 0) return;
    return this._dialog.open(DevtoolsComponent, {
      data: { tab },
      hasBackdrop: true,
      maxWidth: '600px',
      maxHeight: '80%',
      minHeight: '0px',
      backdropClass: ['bg-gray-500/30', 'backdrop-blur-sm'],
      panelClass: [
        'shadow-xl',
        'bg-slate-900',
        'rounded-lg',
        'p-5',
        'dark:text-white'
      ]
    });
  }
}

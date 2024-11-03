import { ChangeDetectionStrategy, Component, ViewEncapsulation } from '@angular/core';

import { State } from '../../state';
import { ToastService } from '../../components/toast';

@Component({
  selector: 'app-devtools-state',
  templateUrl: './devtools-state.component.html',
  styleUrl: './devtools-state.component.scss',
  host: { class: 'app-devtools-state' },
  encapsulation: ViewEncapsulation.None,
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class DevtoolsStateComponent {
  constructor(
    readonly state: State,
    private readonly _toast: ToastService
  ) { }

  onCopy() {
    this._toast.info('Copied To Clipboard!', {
      timeout: 3000
    });
  }
}

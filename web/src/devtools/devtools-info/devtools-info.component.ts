import { ChangeDetectionStrategy, Component, ViewEncapsulation } from '@angular/core';

import pkg from '../../../package.json';

@Component({
  selector: 'app-devtools-info',
  templateUrl: './devtools-info.component.html',
  styleUrl: './devtools-info.component.scss',
  host: { class: 'app-devtools-info' },
  encapsulation: ViewEncapsulation.None,
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class DevtoolsInfoComponent {
  readonly pkg = pkg;
}

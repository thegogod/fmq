import { ChangeDetectionStrategy, Component, ViewEncapsulation } from '@angular/core';

import { ToastService } from './toast.service';

@Component({
  selector: 'app-toasts',
  templateUrl: './toasts.component.html',
  styleUrl: './toasts.component.scss',
  host: { class: 'app-toasts' },
  encapsulation: ViewEncapsulation.None,
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class ToastsComponent {
  constructor(readonly toasts: ToastService) { }
}

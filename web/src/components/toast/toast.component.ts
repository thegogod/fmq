import { ChangeDetectionStrategy, Component, EventEmitter, Input, OnInit, Output, ViewEncapsulation } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import * as feathericons from 'feather-icons';

import { Toast } from './toast';

@Component({
  selector: 'app-toast',
  templateUrl: './toast.component.html',
  styleUrl: './toast.component.scss',
  host: {
    class: 'app-toast',
    '[class]': 'value.type'
  },
  encapsulation: ViewEncapsulation.None,
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class ToastComponent implements OnInit {
  @Input() value!: Toast;

  @Output() dismiss = new EventEmitter<void>();

  readonly progress$ = new BehaviorSubject<number>(100);

  get icon(): keyof (typeof feathericons.icons) {
    switch (this.value.type) {
    case 'info':
      return 'info';
    case 'warn':
      return 'alert-octagon';
    case 'success':
      return 'check';
    case 'error':
      return 'alert-triangle';
    }
  }

  ngOnInit() {
    if (this.value.timeout) {
      const end = Date.now() + this.value.timeout;

      const timeout = setInterval(() => {
        const diff = end - Date.now();
        const perc = (diff / this.value.timeout!) * 100;
        this.progress$.next(perc);
      }, 1);

      setTimeout(() => {
        this.dismiss.emit();
        clearInterval(timeout);
      }, this.value.timeout);
    }
  }
}

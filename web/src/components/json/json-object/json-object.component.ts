import { ChangeDetectionStrategy, Component, EventEmitter, Input, Output, ViewEncapsulation } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Component({
  selector: 'app-json-object',
  templateUrl: './json-object.component.html',
  styleUrl: './json-object.component.scss',
  host: { class: 'app-json-object' },
  encapsulation: ViewEncapsulation.None,
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class JsonObjectComponent {
  @Input() value: any;

  @Output() copy = new EventEmitter<void>();

  readonly expanded$ = new BehaviorSubject<{ [key: string]: boolean }>({ });

  get keys() {
    return Object.keys(this.value).filter(key =>
      typeof this.value[key] !== 'function'
    );
  }

  toggle(key: string) {
    const expanded = this.expanded$.value;
    expanded[key] = !expanded[key];
    this.expanded$.next(expanded);
  }

  isObject(key: string) {
    return typeof this.value[key] === 'object';
  }
}

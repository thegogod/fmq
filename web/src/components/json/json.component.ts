import { ChangeDetectionStrategy, Component, EventEmitter, Input, Output, ViewEncapsulation } from '@angular/core';

@Component({
  selector: 'app-json',
  templateUrl: './json.component.html',
  styleUrl: './json.component.scss',
  host: { class: 'app-json' },
  encapsulation: ViewEncapsulation.None,
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class JsonComponent {
  @Input() value?: any;

  @Output() copy = new EventEmitter<void>();

  get type() { return typeof this.value; }
}

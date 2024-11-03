import { Clipboard } from '@angular/cdk/clipboard';
import { ChangeDetectionStrategy, Component, EventEmitter, HostListener, Input, Output, ViewEncapsulation } from '@angular/core';

@Component({
  selector: 'app-json-string',
  templateUrl: './json-string.component.html',
  styleUrl: './json-string.component.scss',
  host: { class: 'app-json-string' },
  encapsulation: ViewEncapsulation.None,
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class JsonStringComponent {
  @Input() value!: string;

  @Output() copy = new EventEmitter<void>();

  constructor(private readonly _clipboard: Clipboard) { }

  @HostListener('click', ['$event'])
  click(e: Event) {
    e.preventDefault();
    e.stopImmediatePropagation();
    this._clipboard.copy(this.value);
    this.copy.emit();
  }
}

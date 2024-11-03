import { ChangeDetectionStrategy, Component, Inject, TemplateRef, ViewEncapsulation } from '@angular/core';

@Component({
  selector: 'app-popover',
  templateUrl: './popover.component.html',
  styleUrl: './popover.component.scss',
  host: { class: 'app-popover' },
  encapsulation: ViewEncapsulation.None,
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class PopoverComponent {
  get string() {
    return typeof this.data === 'string' ? this.data : undefined;
  }

  get template() {
    return this.data instanceof TemplateRef ? this.data : undefined;
  }

  constructor(
    @Inject('POPOVER_DATA')
    readonly data: string | TemplateRef<void>
  ) { }
}

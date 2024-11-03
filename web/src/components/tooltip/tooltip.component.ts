import { ChangeDetectionStrategy, Component, Inject, TemplateRef, ViewEncapsulation } from '@angular/core';

@Component({
  selector: 'app-tooltip',
  templateUrl: './tooltip.component.html',
  styleUrl: './tooltip.component.scss',
  host: { class: 'app-tooltip' },
  encapsulation: ViewEncapsulation.None,
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class TooltipComponent {
  get string() { return typeof this.data === 'string' ? this.data : undefined; }
  get template() { return this.data instanceof TemplateRef ? this.data : undefined; }

  constructor(
    @Inject('TOOLTIP_DATA')
    readonly data: string | TemplateRef<void>
  ) { }
}

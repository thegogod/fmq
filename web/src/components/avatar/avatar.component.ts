import { ChangeDetectionStrategy, Component, Input, ViewEncapsulation } from '@angular/core';

@Component({
  selector: 'app-avatar',
  templateUrl: './avatar.component.html',
  styleUrl: './avatar.component.scss',
  host: { class: 'app-avatar' },
  encapsulation: ViewEncapsulation.None,
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class AvatarComponent {
  @Input() url?: string;
  @Input()
  get name() { return this._name; }
  set name(v) {
    if (v) {
      const parts = v.split(' ');
      v = parts[0].charAt(0).toUpperCase();

      if (parts.length === 1) {
        v += parts[0].charAt(parts[0].length - 1).toUpperCase();
      } else {
        v += parts[parts.length - 1].charAt(0).toUpperCase();
      }
    }

    this._name = v;
  }
  private _name: string = '??';
}

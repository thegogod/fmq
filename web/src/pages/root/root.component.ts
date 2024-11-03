import { ChangeDetectionStrategy, Component, ViewEncapsulation, isDevMode } from '@angular/core';
import { Router } from '@angular/router';
import { BehaviorSubject } from 'rxjs';
import * as feathericons from 'feather-icons';

import { State } from '../../state';

@Component({
  selector: 'app-root',
  templateUrl: './root.component.html',
  styleUrl: './root.component.scss',
  host: { class: 'app-root' },
  encapsulation: ViewEncapsulation.None,
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class RootComponent {
  readonly isDev = isDevMode();
  readonly expanded$ = new BehaviorSubject(localStorage.getItem('fmq--expanded') === 'true');
  readonly routes: {
    readonly name: string;
    readonly path: string;
    readonly icon: keyof (typeof feathericons.icons);
  }[] = [
    {
      name: 'Dashboard',
      path: 'dashboard',
      icon: 'coffee'
    }
  ];

  constructor(
    readonly state: State,
    private readonly _router: Router,
  ) { }

  openDevtools() {
    this._router.navigateByUrl(this._router.createUrlTree([], {
      queryParams: { devtools: 'info' },
      queryParamsHandling: 'merge',
      preserveFragment: true
    }));
  }

  toggleExpanded() {
    this.expanded$.next(!this.expanded$.value);

    if (this.expanded$.value) {
      localStorage.setItem('fmq--expanded', 'true');
    } else {
      localStorage.removeItem('fmq--expanded');
    }
  }
}

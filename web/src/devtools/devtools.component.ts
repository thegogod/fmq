import { ChangeDetectionStrategy, Component, Inject, OnInit, ViewEncapsulation } from '@angular/core';
import { ComponentPortal, Portal } from '@angular/cdk/portal';
import { DIALOG_DATA } from '@angular/cdk/dialog';
import { BehaviorSubject } from 'rxjs';

import { DevtoolsStateComponent } from './devtools-state';
import { DevtoolsInfoComponent } from './devtools-info';
import { ActivatedRoute, Router } from '@angular/router';

export type DevtoolsTabKey = 'info' | 'state';

interface DevtoolsTab {
  readonly key: DevtoolsTabKey;
  readonly name: string;
  readonly component: Portal<any>;
}

@Component({
  selector: 'app-devtools',
  templateUrl: './devtools.component.html',
  styleUrl: './devtools.component.scss',
  host: { class: 'app-devtools' },
  encapsulation: ViewEncapsulation.None,
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class DevtoolsComponent implements OnInit {
  readonly active$ = new BehaviorSubject<DevtoolsTab | undefined>(undefined);
  readonly tabs: DevtoolsTab[] = [
    {
      key: 'info',
      name: 'Info',
      component: new ComponentPortal(DevtoolsInfoComponent)
    },
    {
      key: 'state',
      name: 'State',
      component: new ComponentPortal(DevtoolsStateComponent)
    }
  ];

  constructor(
    @Inject(DIALOG_DATA) readonly data: { tab: DevtoolsTabKey },
    private readonly _router: Router,
    private readonly _route: ActivatedRoute
  ) { }

  ngOnInit() {
    this.select(this.data.tab);
  }

  select(key: DevtoolsTabKey) {
    this.active$.next(this.tabs.find(t => t.key === key));
    this._router.navigate([], {
      relativeTo: this._route,
      queryParamsHandling: 'merge',
      queryParams: {
        ...this._route.snapshot.queryParams,
        devtools: key
      }
    });
  }
}

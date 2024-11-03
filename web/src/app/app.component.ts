import { ChangeDetectionStrategy, Component, OnDestroy, OnInit, ViewEncapsulation } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Subject, take, takeUntil } from 'rxjs';

import { DevtoolsService } from '../devtools';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss',
  host: { class: 'app' },
  changeDetection: ChangeDetectionStrategy.OnPush,
  encapsulation: ViewEncapsulation.None
})
export class AppComponent {
  private readonly _destroy$ = new Subject<void>();

  constructor(
    private readonly _router: Router,
    private readonly _route: ActivatedRoute,
    private readonly _devtools: DevtoolsService
  ) { }

  ngOnInit() {
    this._route.queryParamMap.pipe(takeUntil(this._destroy$)).subscribe(params => {
      const devtools = params.get('devtools');

      if (
        !devtools ||
        (devtools !== 'info' && devtools !== 'state')
      ) return;

      this._devtools.open(devtools)?.closed.pipe(take(1)).subscribe(() => {
        this._router.navigateByUrl(this._router.createUrlTree([], {
          queryParams: {
            ...this._route.snapshot.queryParams,
            devtools: null
          },
          queryParamsHandling: 'merge',
          preserveFragment: true
        }));
      });
    });
  }

  ngOnDestroy() {
    this._destroy$.next();
    this._destroy$.complete();
  }
}

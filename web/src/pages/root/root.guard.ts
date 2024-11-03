import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, Router } from '@angular/router';

import { Api } from '../../api';
import { State } from '../../state';

@Injectable()
export class RootGuard {
  constructor(
    private readonly _router: Router,
    private readonly _api: Api,
    private readonly _state: State
  ) { }

  async canActivate(_: ActivatedRouteSnapshot) {
    const token = localStorage.getItem('fmq--token');

    if (!this._state.topics$.get('topics')) {
      try {
        const topics = await this._api.topics.get();
        this._state.topics$.set('topics', topics);
      } catch (err) {
        return this._router.createUrlTree(['/sign-in']);
      }
    }

    return token ? true : this._router.createUrlTree(['/sign-in']);
  }

  canActivateChild(route: ActivatedRouteSnapshot) {
    return this.canActivate(route);
  }
}

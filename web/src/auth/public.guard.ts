import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, Router } from '@angular/router';

@Injectable()
export class PublicGuard {
  constructor(private readonly _router: Router) { }

  canActivate(_route: ActivatedRouteSnapshot) {
    const token = localStorage.getItem('fmq--token');
    return !token ? true : this._router.createUrlTree(['/']);
  }

  canActivateChild(route: ActivatedRouteSnapshot) {
    return this.canActivate(route);
  }
}

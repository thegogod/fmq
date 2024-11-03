import { ActivatedRouteSnapshot } from '@angular/router';

export function getRouteParam(route: ActivatedRouteSnapshot, param: string): string {
  if (route.params[param]) return route.params[param];
  if (!route.parent) throw new Error(`route param ${param} not found`);

  return getRouteParam(route.parent, param);
}

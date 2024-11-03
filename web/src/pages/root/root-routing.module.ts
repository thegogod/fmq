import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { AuthModule, PrivateGuard } from '../../auth';

import { RootComponent } from './root.component';

const routes: Routes = [
  {
    path: '',
    canActivate: [PrivateGuard],
    canActivateChild: [PrivateGuard],
    component: RootComponent,
    children: [
      {
        path: 'dashboard',
        loadChildren: () => import('./dashboard/dashboard.module').then(m => m.DashboardModule)
      },
      {
        path: '**',
        redirectTo: 'dashboard'
      }
    ]
  },
  {
    path: '**',
    redirectTo: '/'
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule, AuthModule],
  providers: []
})
export class RootRoutingModule { }

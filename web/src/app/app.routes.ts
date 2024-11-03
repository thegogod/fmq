import { Routes } from '@angular/router';

export const routes: Routes = [
  {
    path: 'sign-in',
    loadChildren: () => import('../pages/sign-in/sign-in.module').then(m => m.SignInModule)
  },
  {
    path: '',
    loadChildren: () => import('../pages/root/root.module').then(m => m.RootModule)
  },
  {
    path: '**',
    redirectTo: '/'
  }
];

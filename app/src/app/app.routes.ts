import { Routes } from '@angular/router';
import {LayoutPublic} from '../layouts/layout-public/layout-public';
import {LayoutPrivate} from '../layouts/layout-private/layout-private';

export const routes: Routes = [
  // public layout
  {
    path: '',
    component: LayoutPublic,
    children: [
      { path: '', redirectTo: '/login', pathMatch: 'full' },
      {
        path: 'login',
        loadComponent: () => import('../pages/public/login/login.page').then(c => c.LoginPage)
      },
      {
        path: 'register',
        loadComponent: () => import('../pages/public/register/register.page').then(c => c.RegisterPage)
      },
    ]
  },

  // private layout
  {
    path: 'dashboard',
    // TODO : bloqué l'accès si pas connecté
    component: LayoutPrivate,
    children: [
      {
        path: '',
        loadComponent: () => import('../pages/private/dashboard/dashboard.page').then(c => c.DashboardPage)
      },
      {
        path: 'design-system',
        loadComponent: () => import('../pages/private/design-system/design-system.page').then(c => c.DesignSystemPage)
      },
    ]
  }
];

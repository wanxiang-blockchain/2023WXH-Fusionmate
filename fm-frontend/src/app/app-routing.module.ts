import { NgModule } from '@angular/core';
import { PreloadAllModules, RouterModule, Routes } from '@angular/router';

const routes: Routes = [
  {
    path: '', loadChildren: () => import('./home/home.module').then(m => m.HomeModule),
    title: 'FusionMate'
  },
  {
    path: 'create', loadChildren: () => import('./create/create.module').then(m => m.CreateModule),
    title: 'Create - FusionMate'
  },
  {
    path: 'list', loadChildren: () => import('./list/list.module').then(m => m.ListModule),
    title: 'List - FusionMate'
  },
  {
    path: 'detail/:collectionID', loadChildren: () => import('./detail/detail.module').then(m => m.DetailModule),
    title: 'Detail - FusionMate'
  },
  {
    path: '**',
    redirectTo: ''
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes, { preloadingStrategy: PreloadAllModules })],
  exports: [RouterModule]
})
export class AppRoutingModule { }

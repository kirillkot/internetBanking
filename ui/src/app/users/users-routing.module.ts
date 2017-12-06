import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { UserFormComponent } from './user-form.component';
import { UserListComponent } from './user-list.component';


const routes: Routes = [
  { path: 'users', component: UserListComponent },
  { path: 'users/form', component: UserFormComponent },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class UsersRoutingModule { }

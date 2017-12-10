import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { UserFormComponent } from './user-form.component';
import { UserListComponent } from './user-list.component';
import { UserEditComponent } from './user-edit.component';


const routes: Routes = [
  { path: 'users', component: UserListComponent },
  { path: 'users/form', component: UserFormComponent },
  { path: 'users/edit', component: UserEditComponent },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class UsersRoutingModule { }

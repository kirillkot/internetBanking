import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { UserFormComponent } from './user-form.component';
import { UserListComponent } from './user-list.component';
import { UserLoginComponent } from './user-login.component';


const routes: Routes = [
  { path: 'users', component: UserListComponent },
  { path: 'users/form', component: UserFormComponent },
  { path: 'users/login', component: UserLoginComponent },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class UsersRoutingModule { }

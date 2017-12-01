import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { UserFormComponent } from './user-form.component';


const routes: Routes = [
    { path: 'users', component: UserFormComponent },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class UsersRoutingModule { }

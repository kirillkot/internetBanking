import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';

import { MaterialModule } from '../material/material.module';

import { UsersRoutingModule } from './users-routing.module';
import { UserService } from './user.service';
import { UserFormComponent } from './user-form.component';
import { UserListComponent } from './user-list.component';
import { UserLoginComponent } from './user-login.component';
import { AbstractModule } from '../abstract/abstract.module';

@NgModule({
  imports: [
    CommonModule,
    HttpClientModule,
    MaterialModule,
    UsersRoutingModule,
    AbstractModule
  ],
  declarations: [
    UserFormComponent,
    UserListComponent,
    UserLoginComponent
  ],
  providers: [UserService]
})
export class UsersModule { }

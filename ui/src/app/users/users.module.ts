import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';

import { MaterialModule } from '../material/material.module';

import { UsersRoutingModule } from './users-routing.module';
import { UserService } from './user.service';
import { UserFormComponent } from './user-form.component';

@NgModule({
  imports: [
    CommonModule,
    HttpClientModule,
    MaterialModule,
    UsersRoutingModule
  ],
  declarations: [UserFormComponent],
  providers: [UserService]
})
export class UsersModule { }

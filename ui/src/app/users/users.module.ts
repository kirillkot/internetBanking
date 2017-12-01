import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { UsersRoutingModule } from './users-routing.module';
import { UserService } from './user.service';
import { UserFormComponent } from './user-form.component';

@NgModule({
  imports: [
    CommonModule,
    UsersRoutingModule
  ],
  declarations: [UserFormComponent],
  providers: [UserService]
})
export class UsersModule { }

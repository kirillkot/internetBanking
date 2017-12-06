import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MaterialModule } from '../material/material.module';

import { AuthGuard } from './auth.guard';
import { LoginService } from './login.service';

import { LoginComponent } from './login.component';
import { LoginRoutingModule } from './login-routing.module';

@NgModule({
  imports: [
    CommonModule,
    MaterialModule,
    LoginRoutingModule
  ],
  providers: [LoginService, AuthGuard],
  declarations: [
    LoginComponent,
  ]
})
export class LoginModule { }

import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { BankComponent } from './bank/bank.component';
import { AdminComponent } from './admin/admin.component';
import { PageNotFoundComponent } from './not-found/not-found.component';

import { UsersModule } from './users/users.module';
import { AppRoutingModule } from './app-routing.module';


@NgModule({
  declarations: [
    AppComponent,
    BankComponent,
    PageNotFoundComponent,
    AdminComponent
  ],
  imports: [
    UsersModule,
    BrowserModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }

import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';


import { AppComponent } from './app.component';
import { AppRoutingModule } from './/app-routing.module';
import { BankComponent } from './bank/bank.component';
import { PageNotFoundComponent } from './not-found/not-found.component';
import { UsersModule } from './users/users.module';


@NgModule({
  declarations: [
    AppComponent,
    BankComponent,
    PageNotFoundComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    UsersModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }

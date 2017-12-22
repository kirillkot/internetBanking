import { NgModule } from '@angular/core';
import { HTTP_INTERCEPTORS } from '@angular/common/http';
import { BrowserModule } from '@angular/platform-browser';

import { Router } from '@angular/router';

import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { PageNotFoundComponent } from './not-found/not-found.component';

import { UsersModule } from './users/users.module';
import { MaterialModule } from './material/material.module';
import { AppRoutingModule } from './app-routing.module';
import { CardsModule } from './cards/cards.module';

import { LoginService } from './login.service';
import { AuthInterceptor } from './auth.interceptor';
import { PaymentsModule } from './payments/payments.module';
import { ConverterService } from './converter.service';


@NgModule({
  declarations: [
    AppComponent,
    PageNotFoundComponent,
    LoginComponent
  ],
  imports: [
    UsersModule,
    CardsModule,
    PaymentsModule,
    BrowserModule,
    AppRoutingModule,
    MaterialModule,
  ],
  providers: [
    {
      provide: HTTP_INTERCEPTORS,
      useClass: AuthInterceptor,
      multi: true,
    },
    LoginService,
    ConverterService
  ],
  bootstrap: [AppComponent]
})
export class AppModule {
  constructor(router: Router) {
    console.log('Routes: ', JSON.stringify(router.config, undefined, 2));
  }
}

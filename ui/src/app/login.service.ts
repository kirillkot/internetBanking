import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { Observable } from 'rxjs/Observable';
import { Subscription } from 'rxjs/Subscription';

import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { AuthCreds, AuthCredsService } from './auth-creds.service';


interface LoginResponse {
  is_admin: boolean;
}

@Injectable()
export class LoginService {
  constructor(
    private http: HttpClient,
    private service: AuthCredsService,
  ) { }

  private errorHanlder(err: any): void {
    console.log(`Login Service: login: failed: ${err}`);
  }

  login(creds: AuthCreds): Observable<void> {
    return this.http
      .get<LoginResponse>('/api/login/', {
        headers:new HttpHeaders().set('Authorization', this.service.buildAuthHeader(creds)),
      })
      .map(data => this.service.setCreds(creds, data.is_admin));
  }

  logout(): void {
    this.service.setCreds(null, false);
  }

}

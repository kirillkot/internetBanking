import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

import { Observable } from 'rxjs/Observable';
import { Subscription } from 'rxjs/Subscription';

import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';


export interface LoginCreds {
  username: string;
  password: string;
}

interface LoginResponse {
  is_admin: boolean;
}

@Injectable()
export class LoginService {
  private creds: LoginCreds = null;
  private admin: boolean;

  constructor(
    private http: HttpClient,
  ) { }

  private errorHanlder(err: any): void {
    console.log(`Login Service: login: failed: ${err}`);
  }

  private buildAuthHeader(creds: LoginCreds): string {
     return 'Basic '+btoa(creds.username + ':' + creds.password);
  }

  login(creds: LoginCreds): Observable<void> {
    return this.http
      .get<LoginResponse>('/api/login', {
        headers:new HttpHeaders().set('Authorization', this.buildAuthHeader(creds)),
      })
      .map(data => {
        this.creds = creds;
        this.admin = data.is_admin;
      });
  }

  getAuthHeader(): string {
    if (this.creds === null) {
      return '';
    }
    return this.buildAuthHeader(this.creds);
  }

  isAdmin(): boolean {
    return this.admin;
  }

  logout(): void {
    this.creds = null;
    this.admin = false;
  }

}

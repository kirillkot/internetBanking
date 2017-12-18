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
  private is_admin: boolean = false;

  constructor(
    private http: HttpClient,
  ) { }

  private errorHanlder(err: any): void {
    console.log(`Login Service: login: failed: ${err}`);
  }

  login(creds: LoginCreds): Observable<void> {
    return this.http
      .post<LoginResponse>('/api/login/', creds)
      .map(data => {
        this.is_admin = data.is_admin;
      });
  }

  isAdmin(): boolean {
    return this.is_admin;
  }

  logout(): void {
    this.is_admin = false;
    document.cookie = 'auth=; path=/;';
  }

}

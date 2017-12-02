import { Observable } from 'rxjs/Observable';

import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { HttpErrorResponse } from '@angular/common/http';


export interface UserForm {
  username: string;
  isAdmin: boolean;
  password: string;
}

export interface User {
  id: number;
  username: string;
  isAdmin: boolean;
}

export interface UserLoginForm {
  username: string;
  password: string;
}

interface UserLoginResponse {
  user: User;
  token: string;
}


@Injectable()
export class UserService {
  current: User;
  token = '';

  constructor(
    private http: HttpClient,
  ) { }

  errorHandler(err: any) {
    console.log(`User Service: error: ${err}`)
  }

  create(user: UserForm): User {
    let newuser: User;
    console.log(`Create user: init with ${user}`)
    this.http
      .post<User>('/api/users/', user)
      .subscribe(
        data => {
          console.log(`Create user: success ${data}`)
          newuser = data;
        },
        this.errorHandler,
      );
    return newuser;
  }

  getUsers(): Observable<User[]> {
    return this.http.get<User[]>('/api/users/');
  }

  delete(id: number): boolean {
    console.log(`Delete user: with ${id}`)
    let ok = false;
    this.http
      .delete(`/api/users/${id}/`)
      .subscribe(
        data => {
          console.log(`Delete user: success ${id}`)
          ok = true;
        },
        this.errorHandler,
      );
    return ok;
  }

  login(user: UserLoginForm): void {
    console.log(`Login: with ${user.username} ${user.password}`)
    this.http
      .post<UserLoginResponse>('/api/users/login/', user)
      .subscribe(
        data => {
          console.log(`Login: success ${user.username}`);
          this.current = data.user;
          this.token = data.token;
        },
        this.errorHandler,
      );
  }
}

import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { BackendService } from '../abstract/backend.service';


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
export class UserService extends BackendService<UserForm, User> {
  current: User;
  token = '';

  constructor(
    http: HttpClient,
  ) {
    super('users', http);
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

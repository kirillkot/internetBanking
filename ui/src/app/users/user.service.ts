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

@Injectable()
export class UserService extends BackendService<UserForm, User> {
  constructor(
    http: HttpClient,
  ) {
    super('users', http);
  }

}

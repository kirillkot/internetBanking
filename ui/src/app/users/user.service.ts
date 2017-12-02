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


@Injectable()
export class UserService {

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
}

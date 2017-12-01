import { Observable } from 'rxjs/Observable';

import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';


export class UserBase {
  public id: number = 0;
  public username: string = '';
  public isAdmin: boolean = false;

  constructor() {}
}

export class User extends UserBase {
  public email: string = '';
  public firstname: string = '';
  public secondname: string = '';
  public gendor: string = '';
  public birthdate: Date;

  constructor() {
    super();
  }
}

@Injectable()
export class UserService {

  constructor(
    private http: HttpClient,
  ) { }

  create(user: User): User {
    this.http
      .post<User>('/api/users/', user)
      .subscribe(
        data => { user.id = data.id },
      );
    return user;
  }

}

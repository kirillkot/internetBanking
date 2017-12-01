import { Injectable } from '@angular/core';

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

  constructor() { }

}

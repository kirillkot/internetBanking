import { Injectable } from '@angular/core';

export class UserBase {
  public id: number;
  public username: string;
  public isAdmin: bool;

  constructor() {}
}

export class User extends UserBase {
  public firstname: string;
  public secondname: string;
  public gendor: string;

  constructor(){}
}

@Injectable()
export class UserService {

  constructor() { }

}

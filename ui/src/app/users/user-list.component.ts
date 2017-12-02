import { Component, OnInit } from '@angular/core';
import { Location } from '@angular/common';

import { User } from './user.service';
import { UserService } from './user.service';


@Component({
  selector: 'app-user-list',
  templateUrl: './user-list.component.html',
  styleUrls: ['./user-list.component.css']
})
export class UserListComponent implements OnInit {
  users: User[]

  constructor(
    private location: Location,
    private service: UserService,
  ) { }

  ngOnInit() {
    this.getUsers()
  }

  getUsers(): void {
    this.users = this.service.getUsers();
    console.log(this.users)
  }

  back(): void {
    this.location.back();
  }

}

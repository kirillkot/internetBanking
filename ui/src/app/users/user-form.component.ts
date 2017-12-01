import { Component, OnInit } from '@angular/core';
import { Location } from '@angular/common';

import { User } from './user.service';
import { UserService } from './user.service';

@Component({
  selector: 'app-user-form',
  templateUrl: './user-form.component.html',
  styleUrls: ['./user-form.component.css']
})
export class UserFormComponent implements OnInit {
  user: User;

  constructor(
    private location: Location,
    private service: UserService,
  ) {
    this.user = new User();
  }

  ngOnInit() {
  }

  create(): User {
    return this.service.create(this.user);
  }

  back(): void {
    this.location.back();
  }

}

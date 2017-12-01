import { Component, OnInit } from '@angular/core';

import { User } from './user.service';

@Component({
  selector: 'app-user-form',
  templateUrl: './user-form.component.html',
  styleUrls: ['./user-form.component.css']
})
export class UserFormComponent implements OnInit {
  user: User;

  constructor() {
    this.user = new User();
  }

  ngOnInit() {
  }

}

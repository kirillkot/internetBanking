import { Component, OnInit, ChangeDetectorRef } from '@angular/core';

import { ManageListComponent } from '../abstract/manage-list.component';

import { UserForm, User } from './user.service';
import { UserService } from './user.service';


@Component({
  selector: 'app-user-list',
  templateUrl: './user-list.component.html',
  styleUrls: ['./user-list.component.css']
})
export class UserListComponent extends ManageListComponent<UserForm, User> {
  displayedColumns = ['id', 'username', 'admin', 'actions'];

  constructor(
    service: UserService,
    detector: ChangeDetectorRef,
  ) {
    super(service, detector);
  }

}

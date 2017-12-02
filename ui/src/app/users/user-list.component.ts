import { of } from 'rxjs/observable/of';
import { Location } from '@angular/common';
import { Observable } from 'rxjs/Observable';
import { DataSource } from '@angular/cdk/collections';
import { Component, OnInit } from '@angular/core';

import { MatTableDataSource } from '@angular/material';

import { User } from './user.service';
import { UserService } from './user.service';


@Component({
  selector: 'app-user-list',
  templateUrl: './user-list.component.html',
  styleUrls: ['./user-list.component.css']
})
export class UserListComponent implements OnInit {
  displayedColumns = ['id', 'username', 'admin', 'actions'];
  dataSource: UsersDataSource;

  constructor(
    private service: UserService,
  ) {
    this.dataSource = new UsersDataSource(service);
  }

  ngOnInit() { }

  delete(id: number): boolean {
    console.log(`Users List: delete ${id}`)
    return this.service.delete(id)
  }
}

export class UsersDataSource extends DataSource<any> {
  constructor(private service: UserService) {
    super();
  }
  connect(): Observable<User[]> {
    return this.service.getUsers();
  }
  disconnect() {}
}

import { Observable } from 'rxjs/Observable';
import { Component, OnInit, ChangeDetectorRef } from '@angular/core';

import { BasicDataSource } from '../abstract/data.source';

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
    private detector: ChangeDetectorRef,
  ) {
    this.dataSource = new UsersDataSource(service);
  }

  ngOnInit() { }

  refresh(): void {
    this.dataSource = new UsersDataSource(this.service);
    this.detector.detectChanges();
  }

  delete(id: number): void {
    console.log(`Users List: delete ${id}`)
    this.service.delete(id);
    this.refresh();
  }
}

export class UsersDataSource extends BasicDataSource<User> { }

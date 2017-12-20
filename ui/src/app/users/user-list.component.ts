import { Component, Inject, OnInit, ChangeDetectorRef } from '@angular/core';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material';

import { ManageListComponent } from '../abstract/manage-list.component';

import { UserForm, User } from './user.service';
import { UserService } from './user.service';


@Component({
  selector: 'app-user-list',
  templateUrl: './user-list.component.html',
  styleUrls: ['./user-list.component.css']
})
export class UserListComponent extends ManageListComponent<UserForm, User> {
  displayedColumns = ['id', 'username', 'admin', 'qr', 'actions'];

  constructor(
    service: UserService,
    detector: ChangeDetectorRef,
    public dialog: MatDialog,
  ) {
    super(service, detector);
  }

  showQR(user: User): void {
    this.dialog.open(UserQRComponent, {
      data: user,
    });
  }

}

@Component({
  selector: 'user-qr-component',
  template: `
  <h2>QR Code (for Two Factor Authority)</h2>
  <img [src]="qr()"/>
`,
})
export class UserQRComponent {
  constructor(
    public dialogRef: MatDialogRef<UserQRComponent>,
    @Inject(MAT_DIALOG_DATA) public data: User,
  ) { }

  qr(): string {
    return `data:image/png;base64,${this.data.qr}`
  }

  onNoClick(): void {
    this.dialogRef.close();
  }

}

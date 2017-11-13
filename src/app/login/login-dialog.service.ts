import { Injectable, Component, Inject } from '@angular/core';
import { MatDialog } from '@angular/material';
import { LoginComponent } from './login.component';
import * as types from '../types';

@Injectable()
export class LoginDialogService {
  constructor(public dialog: MatDialog) {}

  openDialog(isLogin: boolean, callback: (token: types.AccessToken) => void): void {
    const dialogRef = this.dialog.open(LoginComponent, {
      width: '500px',
      data: { 'isLogin': isLogin, 'callback': callback }
    });

    dialogRef.afterClosed().subscribe(result => {
    });
  }
}

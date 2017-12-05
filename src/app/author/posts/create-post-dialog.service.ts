import { Injectable, Component, Inject } from '@angular/core';
import { MatDialog } from '@angular/material';
import { CreatePostComponent } from './create-post/create-post.component';
import * as types from '../../types';

@Injectable()
export class CreatePostDialogService {

  constructor(
    public dialog: MatDialog,
  ) {}

  openDialog(user: types.User, callback: () => void): void {
    const dialogRef = this.dialog.open(CreatePostComponent, {
      width: '700px',
      data: { user: user, callback: callback }
    });

    dialogRef.afterClosed().subscribe(result => {});
  }
}

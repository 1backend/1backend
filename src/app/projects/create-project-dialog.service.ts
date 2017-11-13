import { Injectable, Component, Inject } from '@angular/core';
import { MatDialog } from '@angular/material';
import { CreateProjectDialogComponent } from './create-project-dialog/create-project-dialog.component';
import * as types from '../types';

@Injectable()
export class CreateProjectDialogService {

  animal: string;
  name: string;
  dialogRef;

  constructor(public dialog: MatDialog) {}

  openDialog(callback: (p: types.Project) => void): void {
    this.dialogRef = this.dialog.open(CreateProjectDialogComponent, {
      width: '80%',
      data: {callback: callback}
    });

    this.dialogRef.afterClosed().subscribe(result => {
    });
  }

  closeDialog() {
    this.dialogRef.close();
  }
}

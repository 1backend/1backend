import { Injectable, Component, Inject } from '@angular/core';
import { MatDialog } from '@angular/material';
import { CreateIssueComponent } from './create-issue/create-issue.component';
import * as types from '../../../types';

@Injectable()
export class CreateIssueDialogService {
  constructor(public dialog: MatDialog) {}

  openDialog(project: types.Project): void {
    const dialogRef = this.dialog.open(CreateIssueComponent, {
      width: '700px',
      data: { project: project }
    });

    dialogRef.afterClosed().subscribe(result => {});
  }
}

import { Component, OnInit, Inject, Input } from '@angular/core';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material';
import { HttpClient, HttpParams } from '@angular/common/http';
import { environment } from '../../../../../environments/environment';
import { SessionService } from '../../../../session.service';
import * as types from '../../../../types';
import { FormControl, Validators } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-create-issue',
  templateUrl: './create-issue.component.html',
  styleUrls: ['./create-issue.component.css']
})
export class CreateIssueComponent implements OnInit {
  issueTitle = '';
  commentContent = '';
  errorString;

  onNoClick(): void {
    this.dialogRef.close();
  }

  constructor(
    @Inject(MAT_DIALOG_DATA) public data: any,
    public dialogRef: MatDialogRef<CreateIssueComponent>,
    private http: HttpClient,
    private router: Router,
    private ss: SessionService
  ) {}

  ngOnInit() {}

  addIssue() {
    if (this.commentContent.length < 1 || this.issueTitle.length < 1) {
      this.errorString = 'Title and description must be filled.';
      return;
    }
    const that = this;
    this.http
      .post(environment.backendUrl + '/v1/issue', {
        issue: {
          title: this.issueTitle,
          projectId: this.data.project.Id
        },
        comment: {
          content: this.commentContent
        },
        token: this.ss.getToken()
      })
      .toPromise().then(
        d => {
          this.data.callback();
          this.dialogRef.close();
        }
      );
  }
}

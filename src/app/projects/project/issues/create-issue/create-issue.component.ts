import { Component, OnInit, Inject, Input } from '@angular/core';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material';
import { HttpClient, HttpParams } from '@angular/common/http';
import { ConstService } from '../../../../const.service';
import { SessionService } from '../../../../session.service';
import * as types from '../../../../types';
import { FormControl, Validators } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-create-issue',
  templateUrl: './create-issue.component.html',
  styleUrls: ['./create-issue.component.css'],
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
    private _const: ConstService,
    private ss: SessionService
  ) { }


  ngOnInit() {
  }

  addIssue() {
    if (this.commentContent.length < 1 || this.issueTitle.length < 1) {
      this.errorString = 'Title and description must be filled.';
      return;
    }
    const that = this;
    this.http.post(this._const.url + '/v1/issue', {
      'issue': {
        'title': this.issueTitle,
        'projectId' : this.data.project.Id
      },
      'comment': {
        'content': this.commentContent
      },
      'token': this.ss.getToken()
    }).subscribe(data => {
      location.reload();
    }, error => {
    });
  }

}

import { Component, OnInit, Input } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Router } from '@angular/router';
import * as types from '../../../../types';
import { ConstService } from '../../../../const.service';
import { SessionService } from '../../../../session.service';

@Component({
  selector: 'app-issue',
  templateUrl: './issue.component.html',
  styleUrls: ['./issue.component.css']
})
export class IssueComponent implements OnInit {
  @Input() issueId: string;
  @Input() author;
  @Input() projectName;

  issuerComment: types.Comment = {} as types.Comment;
  projectId: string;
  issue: types.Issue = {};
  commentContent = '';

  constructor(
    private http: HttpClient,
    private _const: ConstService,
    private ss: SessionService
  ) {}

  ngOnInit() {
    this.getIssue();
  }

  addComment() {
    const that = this;
    this.http
      .post(this._const.url + '/v1/comment', {
        comment: {
          content: this.commentContent,
          issueId: this.issueId
        },
        token: this.ss.getToken()
      })
      .subscribe(
        data => {
          this.getIssue();
          this.commentContent = '';
        },
        error => {}
      );
  }

  getIssue() {
    let p = new HttpParams();
    p = p.set('issueId', this.issueId);
    this.http
      .get<types.Issue>(this._const.url + '/v1/issue', { params: p })
      .subscribe(
        issue => {
          this.issue = issue;
          if (this.issue.Comments) {
            this.issue.Comments = this.issue.Comments.sort((a, b) => {
              if (a.CreatedAt === b.CreatedAt) {
                return 0;
              }
              if (a.CreatedAt < b.CreatedAt) {
                return 1;
              }
              return -1;
            });
          }
          this.issuerComment = this.issue.Comments.pop();
          this.issue.Comments = this.issue.Comments.reverse();
        },
        error => {
          console.log(error);
        }
      );
  }

  back() {}
}

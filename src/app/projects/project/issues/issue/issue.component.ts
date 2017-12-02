import { Component, OnInit, Input } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Router } from '@angular/router';
import * as types from '../../../../types';
import { environment } from '../../../../../environments/environment';
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
  currentPage = 0;
  index = 0;

  constructor(
    private http: HttpClient,
    private ss: SessionService
  ) {}

  ngOnInit() {
    this.getIssue();
  }

  addComment() {
    const that = this;
    this.http
      .post(environment.backendUrl + '/v1/comment', {
        comment: {
          content: this.commentContent,
          issueId: this.issueId
        },
        token: this.ss.getToken()
      })
      .toPromise().then(
        data => {
          this.getIssue();
          this.commentContent = '';
        }
      );
  }

  getIssue() {
    let p = new HttpParams();
    p = p.set('issueId', this.issueId);
    this.http
      .get<types.Issue>(environment.backendUrl + '/v1/issue', { params: p })
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
          for (let i = 0; i < this.issue.Comments.length; i++) {
            this.issue.Comments[i].Index = i + 1;
          }
        },
        error => {
          console.log(error);
        }
      );
  }
  pageChanged($event: any) {
    this.currentPage = $event.pageIndex;
  }

  getIndex(): number {
    return this.index++;
  }
}

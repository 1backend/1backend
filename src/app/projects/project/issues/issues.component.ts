import { Component, OnInit, Input, Inject } from '@angular/core';
import * as types from '../../../types';
import { HttpClient, HttpParams } from '@angular/common/http';
import { ConstService } from '../../../const.service';
import { ActivatedRoute, Router } from '@angular/router';
import { CreateIssueDialogService } from '../../../projects/project/issues/create-issue-dialog.service';
import { FilterPipe } from '../../../filter.pipe';

@Component({
  selector: 'app-issues',
  templateUrl: './issues.component.html',
  styleUrls: ['./issues.component.css']
})
export class IssuesComponent implements OnInit {
  @Input() project: types.Project;
  issues: types.Issue[] = [];
  search: '';
  currentPage = 0;
  loaded = false;

  constructor(
    private route: ActivatedRoute,
    private http: HttpClient,
    private _const: ConstService,
    private router: Router,
    private cris: CreateIssueDialogService
  ) {}

  create() {
    this.cris.openDialog(this.project);
  }
  pageChanged($event: any) {
    this.currentPage = $event.pageIndex;
  }

  refresh() {
    let p = new HttpParams();
    p = p.set('projectId', this.project.Id);
    this.http
      .get<types.Issue[]>(this._const.url + '/v1/issues', { params: p })
      .subscribe(
        issues => {
          if (issues) {
            this.issues = issues.sort((a, b) => {
              if (a.CreatedAt === b.CreatedAt) {
                return 0;
              }
              if (a.CreatedAt < b.CreatedAt) {
                return 1;
              }
              return -1;
            });
          }
          this.loaded = true;
        },
        error => {
          console.log(error);
        }
      );
  }

  ngOnInit() {
    this.refresh();
  }
}

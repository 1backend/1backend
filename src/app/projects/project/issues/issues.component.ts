import { Component, OnInit, Input, Inject } from '@angular/core';
import * as types from '../../../types';
import { HttpClient, HttpParams } from '@angular/common/http';
import { environment } from '../../../../environments/environment';
import { ActivatedRoute, Router } from '@angular/router';
import { CreateIssueDialogService } from '../../../projects/project/issues/create-issue-dialog.service';
import { FilterPipe } from '../../../filter.pipe';
import { LoginDialogService } from '../../../login/login-dialog.service';
import { SessionService } from '../../../session.service';

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
    private router: Router,
    private cris: CreateIssueDialogService,
    private lds: LoginDialogService,
    private ss: SessionService
  ) {}

  create() {
    if (!this.ss.getToken()) {
      this.lds.openDialog(true, () => {
        this.cris.openDialog(this.project, () => this.refresh());
      });
    } else {
      this.cris.openDialog(this.project, () => this.refresh());
    }
  }

  pageChanged($event: any) {
    this.currentPage = $event.pageIndex;
  }

  refresh() {
    let p = new HttpParams();
    p = p.set('projectId', this.project.Id);
    this.http
      .get<types.Issue[]>(environment.backendUrl + '/v1/issues', { params: p })
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

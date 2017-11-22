import { Component, OnInit, Input } from '@angular/core';
import { CreateProjectDialogService } from '../create-project-dialog.service';
import { UserService } from '../../user.service';
import { Router } from '@angular/router';
import * as types from '../../types';
import { HttpClient, HttpParams } from '@angular/common/http';
import { ConstService } from '../../const.service';
import { SessionService } from '../../session.service';
import { ActivatedRoute } from '@angular/router';
import { LoginDialogService } from '../../login/login-dialog.service';

@Component({
  selector: 'app-project-list',
  templateUrl: './project-list.component.html',
  styleUrls: ['./project-list.component.css']
})
export class ProjectListComponent implements OnInit {
  @Input() projects: types.Project[] = [];
  @Input() refresh: () => void;

  search: string;
  currentPage = 0;
  author = '';
  isProjectsPage = false;

  constructor(
    private cp: CreateProjectDialogService,
    private lds: LoginDialogService,
    private router: Router,
    public us: UserService,
    private http: HttpClient,
    private _const: ConstService,
    private ss: SessionService,
    private route: ActivatedRoute
  ) {
    this.author = this.route.snapshot.params['author'];
    this.isProjectsPage = this.router.isActive('projects', false);
  }

  ngOnInit() {}

  create() {
    const that = this;
    if (!this.ss.getToken()) {
      this.lds.openDialog(true, () => {
        this.cp.openDialog(proj => {
          that.cp.closeDialog();
          that.router.navigate(['/' + proj.Author + '/' + proj.Name]);
        });
      });
    } else {
      this.cp.openDialog(proj => {
        that.cp.closeDialog();
        that.router.navigate(['/' + proj.Author + '/' + proj.Name]);
      });
    }
  }

  pageChanged($event: any) {
    this.currentPage = $event.pageIndex;
  }
  star(p: types.Project) {
    const that = this;
    this.http.put(this._const.url + '/v1/star', {
      'projectId': p.Id,
      'token': this.ss.getToken(),
    }).subscribe(() => {
      p.Stars++;
    }, error => {
      console.log(error.error);
    });
  }
}

import { Component, OnInit, Input } from '@angular/core';
import { CreateProjectDialogService } from '../create-project-dialog.service';
import { UserService } from '../../user.service';
import { Router } from '@angular/router';
import * as types from '../../types';
import { HttpClient, HttpParams } from '@angular/common/http';
import { ConstService } from '../../const.service';
import { SessionService } from '../../session.service';

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

  constructor(
    private cp: CreateProjectDialogService,
    private router: Router,
    public us: UserService,
    private http: HttpClient,
    private _const: ConstService,
    private ss: SessionService,
  ) {
   }

  ngOnInit(
  ) {
  }

  create() {
    const that = this;
    this.cp.openDialog(proj => {
      that.cp.closeDialog();
      that.router.navigate(['/' + proj.Author + '/' + proj.Name]);
    });
  }

  pageChanged($event: any) {
    this.currentPage = $event.pageIndex;
  }
  star(p: types.Project) {
    const that = this;
    this.http.put(this._const.url + '/v1/star', {
      'projectId': p.Id,
      'token': this.ss.getToken()
    }).subscribe(star => {
    }, error => {
      console.log(error.error);
    });
  }

}

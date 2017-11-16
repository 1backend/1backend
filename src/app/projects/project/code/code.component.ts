import { Component, OnInit, Input } from '@angular/core';
import * as types from '../../../types';
import { FilterPipe } from '../../../filter.pipe';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Router } from '@angular/router';
import { SessionService } from '../../../session.service';
import { UserService } from '../../../user.service';
import { ConstService } from '../../../const.service';

@Component({
  selector: 'app-code',
  templateUrl: './code.component.html',
  styleUrls: ['./code.component.css']
})
export class CodeComponent implements OnInit {
  @Input() project: types.Project;
  @Input() refresh: () => void;

  user: types.User;
  search: string;
  endpoints: types.Endpoint[] = [];
  leftWidth = 49;
  rightWidth = 49;
  newEndpoint: types.Endpoint = {
    Method: 'GET',
    Url: '/'
  };
  status: boolean;
  readme: string;
  currentPage = 0;
  mode = '';
  readOnly = true;

  constructor(
    private http: HttpClient,
    private router: Router,
    private us: UserService,
    private ss: SessionService,
    private _const: ConstService
  ) {
    this.user = this.us.user;
  }

  ngOnInit() {
    this.getMode();
    if (this.user.Nick === this.project.Author) {
      this.readOnly = false;
    }
  }

  add() {
    this.project.Endpoints.push(this.newEndpoint);
    this.save();
  }

  save() {
    this.http
    .put(this._const.url + '/v1/project', {
      project: this.project,
      token: this.ss.getToken(),
    })
    .subscribe(
      data => {
        this.refresh();
      },
      error => {}
    );
  }

  getMode() {
    switch (this.project.Mode) {
      case 'go': {
        this.mode = 'golang';
        break;
      }
      case 'nodejs': {
        this.mode = 'javascript';
        break;
      }
      case 'typescript': {
        this. mode = 'typescript';
      }
    }
  }

  pageChanged($event: any) {
    this.currentPage = $event.pageIndex;
  }

  
  delete(e: types.Endpoint) {
    let p = new HttpParams;
    p = p.set('id', e.Id);
    this.http.delete(this._const.url + '/v1/endpoint', { params: p }).subscribe(
      data => {
        this.refresh();
      },
      error => {}
    );
  }

  setRightFullScreen() {
    this.rightWidth = 80;
    this.leftWidth = 18;
  }
  smallScreen() {
    this.rightWidth = 49;
    this.leftWidth = 49;
  }
  setLeftFullScreen() {
    this.rightWidth = 19;
    this.leftWidth = 80;
  }

  isLeftFullScreen(): boolean {
    return this.leftWidth === 80;
  }

  isRightFullScreen(): boolean {
    return this.rightWidth === 80;
  }

  reveal(e: types.Endpoint) {
    e.Selected = !e.Selected;
  }

  goSql() {
    this.router.navigate([
      '/' + this.project.Author + '/' + this.project.Name + '/' + 'sql-console'
    ]);
    location.reload(); // XD
  }
}

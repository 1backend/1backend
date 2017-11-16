import { Component, OnInit, Input } from '@angular/core';
import { Location } from '@angular/common';
import { ActivatedRoute, Router } from '@angular/router';
import { HttpClient, HttpParams } from '@angular/common/http';
import { MatProgressSpinnerModule, MatTabGroup } from '@angular/material';
import * as types from '../../types';
import { SessionService } from '../../session.service';
import { ConstService } from '../../const.service';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material';
import { UserService } from '../../user.service';

interface PingResponse {
  pong: boolean;
}

@Component({
  selector: 'app-project',
  templateUrl: './project.component.html',
  styleUrls: ['./project.component.css']
})
export class ProjectComponent implements OnInit {
  selectedIndex = 0;
  tab = '';
  author = '';
  projectName = '';
  color = 'primary';
  mode = 'indeterminate';
  value = 50;
  lastBuild: types.Build;
  leftWidth = 19;
  that = this;
  status: boolean;
  issueId = "";
  loaded = false;

  project: types.Project = {
    Endpoints: [],
    Builds: []
  };

  constructor(
    private route: ActivatedRoute,
    private http: HttpClient,
    private ss: SessionService,
    private location: Location,
    private router: Router,
    private _const: ConstService,
    public us: UserService
  ) {
    this.author = this.route.snapshot.params['author'];
    this.projectName = this.route.snapshot.params['project'];
    this.tab = this.route.snapshot.params['tab'];
    this.issueId = this.route.snapshot.params['issueId'];
    this.makeRefresh()();
    if (this.tab === 'sql-console') {
      this.selectedIndex = 2;
    }
    if (this.tab === 'builds') {
      this.selectedIndex = 1;
    }
    if (this.tab === 'issues' || this.issueId) {
      this.selectedIndex = 3;
    }
  }

  makeRefresh(): () => void {
    const that = this;
    return () => {
      let p = new HttpParams();
      console.log(that.leftWidth);
      p = p.set('author', that.author);
      p = p.set('project', that.projectName);
      p = p.set('token', that.ss.getToken());
      that.http
        .get<types.Project>(this._const.url + '/v1/project', {
          params: p
        })
        .subscribe(
          proj => {
            if (proj.Builds) {
              proj.Builds = proj.Builds.sort((a, b) => {
                if (a.CreatedAt === b.CreatedAt) {
                  return 0;
                }
                if (a.CreatedAt < b.CreatedAt) {
                  return 1;
                }
                return -1;
              });
            }
            if (proj.Builds) {
              that.lastBuild = proj.Builds[0];
            }
            if (proj.Endpoints) {
              proj.Endpoints = proj.Endpoints.sort((a, b) => {
                if (a.CreatedAt === b.CreatedAt) {
                  return 0;
                }
                if (a.CreatedAt < b.CreatedAt) {
                  return 1;
                }
                return -1;
              });
            }
            if (!proj.ReadMe) {
              proj.ReadMe = proj.Name + '\n===\nThis project doesn\'t have a readme yet.';
            }
            that.project = proj;
            that.loaded = true;
          },
          error => {
            console.log(error);
          }
        );
    };
  }

  ngOnInit() {
    this.getStatus();
  }

  delete() {
    this.http.delete(this._const.url + '/v1/project', {}).subscribe(
      data => {
        this.makeRefresh()();
      },
      error => {
        alert(JSON.stringify(error));
      }
    );
  }

  selectedIndexChange(tabGroup: MatTabGroup) {
    const pid = tabGroup._tabs.find((e, i, a) => i === tabGroup.selectedIndex)
      .content.viewContainerRef.element.nativeElement.dataset.pid;
    if (pid !== 'code') {
      this.location.go('/' + this.project.Author + '/' + this.project.Name + '/' + pid);
    } else {
      this.location.go('/' + this.project.Author + '/' + this.project.Name);
    }
  }

  getStatus() {
    this.http
      .get<PingResponse>(
        this._const.url + '/app/' +
          this.author +
          '/' +
          this.projectName +
          '/ping'
      )
      .subscribe(data => {
        this.status = data.pong;
      });
  }
}

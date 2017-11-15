import { Component, OnInit } from '@angular/core';
import * as types from '../types';
import { FilterPipe } from '../filter.pipe';
import { HttpClient, HttpParams } from '@angular/common/http';
import { SessionService } from '../session.service';
import { ConstService } from '../const.service';

@Component({
  selector: 'app-projects',
  templateUrl: './projects.component.html',
  styleUrls: ['./projects.component.css']
})
export class ProjectsComponent implements OnInit {
  currentPage = 0;
  search = '';
  projects: types.Project[] = [];
  lastUpdated: types.Project;

  constructor(
    private http: HttpClient,
    private ss: SessionService,
    private _const: ConstService
  ) {
    this.refresh();
  }

  refresh() {
    const that = this;
    let p = new HttpParams();
    p = p.set('token', this.ss.getToken());
    this.http
      .get<types.Project[]>(this._const.url + '/v1/projects', { params: p })
      .subscribe(
        projs => {
          this.projects = projs.sort((a, b) => {
            for (let p of projs) {
              if (a.UpdatedAt === b.UpdatedAt) {
                return 0;
              }
            } if (a.UpdatedAt < b.UpdatedAt) {
              return 1;
            }
            return -1;
          });
          that.lastUpdated = projs[0];
        },
        error => {
          console.log(error);
        }
      );
  }
  pageChanged($event: any) {
    this.currentPage = $event.pageIndex;
  }

  ngOnInit() {}

}

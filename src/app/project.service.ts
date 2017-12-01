import { Injectable } from '@angular/core';
import * as types from './types';
import { HttpClient, HttpParams } from '@angular/common/http';
import { SessionService } from './session.service';
import { ConstService } from './const.service';
import { Observable } from 'rxjs/Observable';

interface PingResponse {
  pong: boolean;
}

@Injectable()
export class ProjectService {
  constructor(
    private http: HttpClient,
    private _const: ConstService,
    private ss: SessionService
  ) {}

  // returns last updated projects first
  listByNick(nick: string): Promise<types.Project[]> {
    let p = new HttpParams();
    p = p.set('nick', nick);
    p = p.set('token', this.ss.getToken());
    return new Promise<types.Project[]>((resolve, reject) => {
      this.http
        .get<types.Project[]>(this._const.url + '/v1/projects', {
          params: p
        })
        .toPromise()
        .then(projs => {
          projs = projs.sort((a, b) => {
            if (a.UpdatedAt === b.UpdatedAt) {
              return 0;
            }
            if (a.UpdatedAt < b.UpdatedAt) {
              return 1;
            }
            return -1;
          });
          resolve(projs);
        });
    });
  }

  getByAuthorAndProjectName(
    author: string,
    projectName: string
  ): Promise<types.Project> {
    return new Promise<types.Project>((resolve, reject) => {
      let p = new HttpParams();
      p = p.set('author', author);
      p = p.set('project', projectName);
      p = p.set('token', this.ss.getToken());
      this.http
        .get<types.Project>(this._const.url + '/v1/project', {
          params: p
        })
        .toPromise()
        .then(proj => {
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
          resolve(proj);
        });
    });
  }

  getStatus(author: string, projectName: string): Promise<PingResponse> {
    return new Promise<PingResponse>((resolve, reject) => {
      this.http
        .get<PingResponse>(
          this._const.url + '/app/' + author + '/' + projectName + '/ping'
        )
        .toPromise();
    });
  }

  star(projectId: string): Promise<void> {
    return new Promise<void>((resolve, reject) => {
      this.http
        .put(this._const.url + '/v1/star', {
          projectId: projectId,
          token: this.ss.getToken()
        })
        .toPromise();
    });
  }

  unstar(projectId: string): Promise<void> {
    return new Promise<void>((resolve, reject) => {
      let p = new HttpParams();
      p = p.set('projectId', projectId);
      p = p.set('token', this.ss.getToken());
      this.http
        .delete(this._const.url + '/v1/star', {
          params: p
        })
        .toPromise();
    });
  }

  list(): Promise<types.Project[]> {
    let p = new HttpParams();
    p = p.set('token', this.ss.getToken());
    return new Promise<types.Project[]>((resolve, reject) => {
      this.http
        .get<types.Project[]>(this._const.url + '/v1/projects', {
          params: p
        })
        .toPromise()
        .then(projs => {
          projs = projs.sort((a, b) => {
            if (a.UpdatedAt === b.UpdatedAt) {
              return 0;
            }
            if (a.UpdatedAt < b.UpdatedAt) {
              return 1;
            }
            return -1;
          });
          resolve(projs);
        });
    });
  }

  save(project: types.Project): Promise<void> {
    return new Promise<void>((resolve, reject) => {
      this.http
        .put(this._const.url + '/v1/project', {
          project: project,
          token: this.ss.getToken()
        })
        .toPromise();
    });
  }
}

import { Injectable } from '@angular/core';
import * as types from './types';
import { HttpClient, HttpParams } from '@angular/common/http';
import { SessionService } from './session.service';
import { environment } from '../environments/environment';
import { Observable } from 'rxjs/Observable';

interface PingResponse {
  pong: boolean;
}

@Injectable()
export class ProjectService {
  constructor(
    private http: HttpClient,
    private ss: SessionService
  ) {}

  // returns last updated projects first
  listByNick(nick: string): Promise<types.Project[]> {
    let p = new HttpParams();
    p = p.set('nick', nick);
    p = p.set('token', this.ss.getToken());
    return this.http
      .get<types.Project[]>(environment.backendUrl + '/v1/projects', {
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
        return projs;
      });
  }

  getByAuthorAndProjectName(
    author: string,
    projectName: string
  ): Promise<types.Project> {
    let p = new HttpParams();
    p = p.set('author', author);
    p = p.set('project', projectName);
    p = p.set('token', this.ss.getToken());
    return this.http
      .get<types.Project>(environment.backendUrl + '/v1/project', {
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
        return proj;
      });
  }

  getStatus(author: string, projectName: string): Promise<PingResponse> {
    return this.http
      .get<PingResponse>(
        environment.backendUrl + '/app/' + author + '/' + projectName + '/ping'
      )
      .toPromise();
  }

  star(projectId: string): Promise<void> {
    return this.http
      .put<void>(environment.backendUrl + '/v1/star', {
        projectId: projectId,
        token: this.ss.getToken()
      })
      .toPromise();
  }

  unstar(projectId: string): Promise<void> {
    let p = new HttpParams();
    p = p.set('projectId', projectId);
    p = p.set('token', this.ss.getToken());
    return this.http
      .delete<void>(environment.backendUrl + '/v1/star', {
        params: p
      })
      .toPromise();
  }

  list(): Promise<types.Project[]> {
    let p = new HttpParams();
    p = p.set('token', this.ss.getToken());
    return this.http
      .get<types.Project[]>(environment.backendUrl + '/v1/projects', {
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
        return projs;
      });
  }

  save(project: types.Project): Promise<void> {
    return this.http
      .put<void>(environment.backendUrl + '/v1/project', {
        project: project,
        token: this.ss.getToken()
      })
      .toPromise();
  }
}

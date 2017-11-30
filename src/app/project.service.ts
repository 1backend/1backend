import { Injectable } from '@angular/core';
import * as types from './types';
import { HttpClient, HttpParams } from '@angular/common/http';
import { SessionService } from './session.service';
import { ConstService } from './const.service';
import { Observable } from 'rxjs/Observable';

@Injectable()
export class ProjectService {
  constructor(
    private http: HttpClient,
    private _const: ConstService,
    private ss: SessionService
  ) {
  }

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
      .subscribe(
        projs => {
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
        },
        error => {
          reject(error);
        }
      );
    });
  }
}

import { Injectable } from '@angular/core';
import * as types from './types';
import { HttpClient, HttpParams } from '@angular/common/http';
import { SessionService } from './session.service';
import { ConstService } from './const.service';

@Injectable()

export class UserService {
  user: types.User = {} as types.User;

  constructor(
    private http: HttpClient,
    private _const: ConstService,
    private sess: SessionService,
  ) {
    this.get().then(user => {
      for (const k of Object.keys(user)) {
        this.user[k] = user[k];
      }
    });
  }
  // gets current user
  get(): Promise<types.User> {
    return new Promise<types.User>((resolve, reject) => {
      if (!this.sess.getToken() || this.sess.getToken().length === 0) {
        setTimeout(function () {
          resolve({} as types.User);
        }, 1); // hack to get around navbar router.isActive not being ready immediately
        return;
      }
      this.http.get<types.User>(this._const.url + '/v1/user?token=' + this.sess.getToken()).subscribe(data => {
        for (const k of Object.keys(data)) {
          this.user[k] = data[k];
        }
        resolve(data);
      }, error => {
        reject(error);
      });
    });
  }

  save(u: types.User): Promise<void> {
    return new Promise<void>((resolve, reject) => {
      this.http.post(this._const.url + '/v1/user', {
        'password': u.Password,
        'user': {
          'Name': u.Name,
          'Id': u.Id
        },
        'token': this.sess.getToken()
      }).subscribe(data => {
        resolve();
      }, error => {
        reject(error);
      });
    });
  }
}

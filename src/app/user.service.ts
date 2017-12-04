import { Injectable } from '@angular/core';
import * as types from './types';
import { HttpClient, HttpParams } from '@angular/common/http';
import { SessionService } from './session.service';
import { Observable } from 'rxjs/Observable';
import { environment } from '../environments/environment';

interface LoginResponse {
  token: types.AccessToken;
}

interface RegisterResponse {
  token: types.AccessToken;
}

@Injectable()
export class UserService {
  user: types.User = {} as types.User;

  constructor(private http: HttpClient, private sess: SessionService) {
    this.get().then(user => {
      for (const k of Object.keys(user)) {
        this.user[k] = user[k];
      }
    });
  }

  loggedIn(): boolean {
    return this.sess.getToken().length > 0;
  }

  logout() {
    this.sess.setToken('');
  }

  // gets current user
  get(): Promise<types.User> {
    return new Promise<types.User>((resolve, reject) => {
      if (!this.sess.getToken() || this.sess.getToken().length === 0) {
        setTimeout(function() {
          resolve({} as types.User);
        }, 1); // hack to get around navbar router.isActive not being ready immediately
        return;
      }
      this.http
        .get<types.User>(
          environment.backendUrl + '/v1/user?token=' + this.sess.getToken()
        )
        .toPromise()
        .then(user => {
          for (const k of Object.keys(user)) {
            this.user[k] = user[k];
          }
          resolve(user);
        });
    });
  }

  getByNick(nick: string): Promise<types.User> {
    let p = new HttpParams();
    p = p.set('token', this.sess.getToken());
    p = p.set('nick', nick);
    return this.http
      .get<types.User>(environment.backendUrl + '/v1/user', { params: p })
      .toPromise();
  }

  save(u: types.User): Promise<void> {
    return this.http
      .post<void>(environment.backendUrl + '/v1/user', {
        password: u.Password,
        user: {
          Name: u.Name,
          Id: u.Id
        },
        token: this.sess.getToken()
      })
      .toPromise();
  }

  // saves the object in the user service
  saveSelf(): Promise<void> {
    return this.http
      .put<void>(environment.backendUrl + '/v1/user', {
        user: {
          avatarLink: this.user.AvatarLink,
          name: this.user.Name,
          email: this.user.Email,
          nick: this.user.Nick
        },
        token: this.sess.getToken()
      })
      .toPromise();
  }

  login(email: string, password: string): Promise<LoginResponse> {
    return this.http
      .post<LoginResponse>(environment.backendUrl + '/v1/login', {
        email: email,
        password: password
      })
      .toPromise()
      .then(loginResp => {
        this.sess.setToken(loginResp.token.Token);
        return loginResp;
      });
  }

  register(
    userName: string,
    email: string,
    password: string
  ): Promise<RegisterResponse> {
    return this.http
      .post<RegisterResponse>(environment.backendUrl + '/v1/register', {
        nick: userName,
        email: email,
        password: password
      })
      .toPromise()
      .then(registerResp => {
        this.sess.setToken(registerResp.token.Token);
        return registerResp;
      });
  }
}

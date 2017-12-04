import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import * as types from './types';
import { SessionService } from './session.service';
import { environment } from '../environments/environment';

@Injectable()
export class CommentService {
  constructor(private ss: SessionService, private http: HttpClient) {}

  update(comment: types.Comment): Promise<void> {
    return this.http
      .put<void>(environment.backendUrl + '/v1/comment', {
        comment: comment,
        token: this.ss.getToken()
      })
      .toPromise();
  }
}

import { Injectable } from '@angular/core';
import * as types from './types';
import { environment } from '../environments/environment';
import { HttpClient, HttpParams } from '@angular/common/http';
import { SessionService } from './session.service';
import { Http } from '@angular/http/src/http';

@Injectable()
export class PostService {
  constructor(private http: HttpClient, private ss: SessionService) {}

  list(author: string): Promise<types.Post[]> {
    let p = new HttpParams();
    p = p.set('nick', author);
    p = p.set('token', this.ss.getToken());
    return this.http
      .get<types.Post[]>(environment.backendUrl + '/v1/posts', { params: p })
      .toPromise()
      .then(posts => {
        posts = posts.sort((a, b) => {
          if (a.UpdatedAt === b.UpdatedAt) {
            return 0;
          }
          if (a.UpdatedAt < b.UpdatedAt) {
            return 1;
          }
          return -1;
        });
        return posts;
      });
  }

  get(postId: string): Promise<types.Post> {
    let p = new HttpParams();
    p = p.set('postId', postId);
    return this.http
      .get<types.Post>(environment.backendUrl + '/v1/post', { params: p })
      .toPromise();
  }

  edit(post: types.Post): Promise<types.Post> {
    return this.http
      .put(environment.backendUrl + '/v1/post', {
        post: {
          content: post.Content,
          title: post.Title,
          user: post.User
        }
      })
      .toPromise();
  }

  create(post: types.Post): Promise<void> {
    return this.http
      .post<void>(environment.backendUrl + '/v1/post', {
        post: post,
        token: this.ss.getToken()
      })
      .toPromise();
  }
}

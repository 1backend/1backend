import { Component, OnInit, ViewEncapsulation, Input } from '@angular/core';
import * as types from '../../types';
import { UserService } from '../../user.service';
import { Router, ActivatedRoute } from '@angular/router';
import { HttpClient, HttpParams } from '@angular/common/http';
import { ConstService } from '../../const.service';
import { error } from 'selenium-webdriver';

@Component({
  selector: 'app-posts',
  templateUrl: './posts.component.html',
  styleUrls: ['./posts.component.css'],
  encapsulation: ViewEncapsulation.None
})
export class PostsComponent implements OnInit {
  @Input () author: string;
  posts: types.Posts[] = [];
  postId: '';

  constructor(
    public us: UserService,
    private router: Router,
    private http: HttpClient,
    private route: ActivatedRoute,
    private _const: ConstService
  ) {
   }

   getPosts() {
    let p = new HttpParams();
    p = p.set('nick', this.author);
    this.http.get<types.Posts[]>(this._const.url + '/v1/posts', { params: p })
    .subscribe(
      posts => {
        this.posts = posts;
      },
      error => {
        console.log('error');
      });
  }

  ngOnInit() {
    this.getPosts();
  }

}

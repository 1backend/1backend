import { Component, OnInit, ViewEncapsulation, Input } from '@angular/core';
import * as types from '../../types';
import { UserService } from '../../user.service';
import { Router, ActivatedRoute } from '@angular/router';

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
    private route: ActivatedRoute,
  ) {
   }

   getPosts() {
    //
  }

  ngOnInit() {
    this.getPosts();
  }

}

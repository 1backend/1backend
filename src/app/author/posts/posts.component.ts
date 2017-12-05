import { Component, OnInit, ViewEncapsulation, Input } from '@angular/core';
import * as types from '../../types';
import { UserService } from '../../user.service';
import { Router, ActivatedRoute } from '@angular/router';
import { PostService } from '../../post.service';
import { CreatePostDialogService } from './create-post-dialog.service';

@Component({
  selector: 'app-posts',
  templateUrl: './posts.component.html',
  styleUrls: ['./posts.component.css'],
  encapsulation: ViewEncapsulation.None
})
export class PostsComponent implements OnInit {
  @Input() author: string;
  posts: types.Post[] = [];
  search = '';


  constructor(
    public us: UserService,
    private router: Router,
    private route: ActivatedRoute,
    private ps: PostService,
    private cpds: CreatePostDialogService
  ) {}

  getPosts() {
    this.ps
      .list(this.author)
      .then(posts => {
        this.posts = posts;
      })
      .catch(err => {
        err = console.log('error');
      });
  }

  create() {
    this.cpds.openDialog(this.us.user, () => this.getPosts());
  }

  ngOnInit() {
    this.getPosts();
  }
}

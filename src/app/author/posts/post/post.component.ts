import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { PostService } from '../../../post.service';
import * as types from '../../../types';
import { UserService } from '../../../user.service';
@Component({
  selector: 'app-post',
  templateUrl: './post.component.html',
  styleUrls: ['./post.component.css'],
  encapsulation: ViewEncapsulation.None
})
export class PostComponent implements OnInit {
  postId: string;
  post: types.Post = {};

  constructor(
    private route: ActivatedRoute,
    private ps: PostService,
    public us: UserService
  ) {}

  getPost() {
    this.ps
      .get(this.postId)
      .then(post => {
        this.post = post;
      })
      .catch(err => {
        err = console.log('error');
      });
  }

  editPost() {
    this.post.Editing = !this.post.Editing;
  }

  updatePost() {
    this.ps
      .edit(this.post)
      .then(post => {
        this.post = post;
        this.getPost();
      })
      .catch(err => (err = console.log('error')));
  }

  ngOnInit() {
    this.postId = this.route.snapshot.params['postId'];
    this.getPost();
  }
}

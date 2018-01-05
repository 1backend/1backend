import {
  Component,
  OnInit,
  ViewEncapsulation,
  Input,
  Inject
} from '@angular/core';
import { PostService } from '../../../post.service';
import { UserService } from '../../../user.service';
import * as types from '../../../types';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material';

@Component({
  selector: 'app-create-post',
  templateUrl: './create-post.component.html',
  styleUrls: ['./create-post.component.css'],
  encapsulation: ViewEncapsulation.None
})
export class CreatePostComponent implements OnInit {
  @Input() author: string;
  post: types.Post = {};

  constructor(
    private ps: PostService,
    public us: UserService,
    @Inject(MAT_DIALOG_DATA) public data: any,
    public dialogRef: MatDialogRef<CreatePostComponent>
  ) {}

  createPost() {
    this.ps
      .create(this.post)
      .then(d => {
        this.data.callback();
        this.dialogRef.close();
      })
      .catch(err => console.log('error'));
  }

  ngOnInit() {}
}

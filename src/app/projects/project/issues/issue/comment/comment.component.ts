import {
  Component,
  OnInit,
  ViewEncapsulation,
  Output,
  EventEmitter,
  Input
} from '@angular/core';
import * as types from '../../../../../types';
import { environment } from '../../../../../../environments/environment';
import { UserService } from '../../../../../user.service';
import { CommentService } from '../../../../../comment.service';

@Component({
  selector: 'app-comment',
  templateUrl: './comment.component.html',
  styleUrls: ['./comment.component.css'],
  encapsulation: ViewEncapsulation.None
})
export class CommentComponent implements OnInit {
  @Output() onCommentUpdate = new EventEmitter<void>();
  @Input() comment: types.Comment;

  constructor(private us: UserService, private cs: CommentService) {}

  ngOnInit() {}

  updateComment(comment: types.Comment) {
    this.cs.update(comment).then(() => {
      this.onCommentUpdate.emit();
    });
  }

  editComment(comment: types.Comment) {
    comment.Editing = true;
  }
}

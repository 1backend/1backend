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
  editedCommentContent = '';

  constructor(private us: UserService, private cs: CommentService) {}

  ngOnInit() {
    this.editedCommentContent = this.comment.Content;
  }

  updateComment() {
    this.comment.Content = this.editedCommentContent;
    this.cs.update(this.comment).then(() => {
      this.onCommentUpdate.emit();
    });
  }

  editComment() {
    this.comment.Editing = !this.comment.Editing;
  }
}

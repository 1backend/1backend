import {
  Component,
  OnInit,
  EventEmitter,
  Input,
  Output,
  ViewEncapsulation
} from '@angular/core';
import { UserService } from '../../user.service';
import { NotificationsService } from 'angular2-notifications';
import * as types from '../../types';

@Component({
  selector: 'app-profile-edit',
  templateUrl: './profile-edit.component.html',
  styleUrls: ['./profile-edit.component.css'],
  encapsulation: ViewEncapsulation.None
})
export class ProfileEditComponent implements OnInit {
  @Input() user: types.User;
  @Output() onSelfSave = new EventEmitter<void>();
  saved = true;
  oldPw = '';
  newPw = '';
  newPwCf = '';

  constructor(public us: UserService, private notif: NotificationsService) {}

  ngOnInit() {}

  save() {
    if (!this.valid()) {
      return;
    }
    this.us.saveSelf().then(data => {
      this.saved = true;
      this.onSelfSave.emit();
    });
  }

  edit() {
    this.saved = false;
  }

  changePassword() {
    if (!this.pwChangeValid()) {
      return;
    }
    this.us.changePassword(this.oldPw, this.newPw);
  }

  valid(): boolean {
    if (!this.user.Nick) {
      this.notif.error('Nickname is empty');
      return false;
    }
    if (!this.user.Email) {
      this.notif.error('Email is empty');
      return false;
    }
    return true;
  }

  pwChangeValid(): boolean {
    if (this.newPw.length < 1) {
      this.notif.error('Password is empty');
      return false;
    }
    if (this.newPw !== this.newPwCf) {
      this.notif.error('New password and confirm do not match');
      return false;
    }
    return true;
  }
}

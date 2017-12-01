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

@Component({
  selector: 'app-profile-edit',
  templateUrl: './profile-edit.component.html',
  styleUrls: ['./profile-edit.component.css'],
  encapsulation: ViewEncapsulation.None
})
export class ProfileEditComponent implements OnInit {
  @Input() author = '';
  @Output() onSelfSave = new EventEmitter<void>();
  saved = true;

  constructor(public us: UserService, private notif: NotificationsService) {}

  ngOnInit() {}

  save() {
    if (!this.valid()) {
      return;
    }
    this.us.saveSelf().toPromise().then(
      data => {
        this.saved = true;
        this.onSelfSave.emit();
      }
    );
  }

  edit() {
    this.saved = false;
  }

  valid(): boolean {
    if (!this.us.user.Nick) {
      this.notif.error('Nickname is empty');
      return false;
    }
    if (!this.us.user.Email) {
      this.notif.error('Email is empty');
      return false;
    }
    return true;
  }
}

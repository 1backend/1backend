import {
  Component,
  OnInit,
  EventEmitter,
  Output,
  ViewEncapsulation
} from '@angular/core';
import { UserService } from '../../user.service';
import { EventService } from '../../event.service';
import { NotificationsService } from 'angular2-notifications';

@Component({
  selector: 'app-profile-edit',
  templateUrl: './profile-edit.component.html',
  styleUrls: ['./profile-edit.component.css'],
  encapsulation: ViewEncapsulation.None
})
export class ProfileEditComponent implements OnInit {
  @Output() onSelfSave = new EventEmitter<void>();
  saved = true;

  constructor(
    private us: UserService,
    private ev: EventService,
    private notif: NotificationsService
  ) {}

  ngOnInit() {}

  save() {
    if (!this.valid()) {
      return;
    }
    this.us.saveSelf().subscribe(
      data => {
        this.saved = true;
        this.onSelfSave.emit();
      },
      error => {}
    );
  }

  edit() {
    this.saved = false;
  }

  valid(): boolean {
    if (!this.us.user.Nick) {
      this.notif.error('Nickname is empty');
      return true;
    }
    if (!this.us.user.Email) {
      this.notif.error('Email is empty');
      return true;
    }
    return false;
  }
}

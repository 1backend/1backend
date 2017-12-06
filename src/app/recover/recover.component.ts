import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { NotificationsService } from 'angular2-notifications';
import { Title } from '@angular/platform-browser';

@Component({
  selector: 'app-recover',
  templateUrl: './recover.component.html',
  styleUrls: ['../login/login.component.css']
})
export class RecoverComponent implements OnInit {
  email: string;

  constructor(
    private http: HttpClient,
    private title: Title,
    private notif: NotificationsService
  ) {}

  ngOnInit() {
    this.title.setTitle('Recover');
  }

  sendReset() {
    this.http
      .post(environment.backendUrl + '/v1/send-reset', {
        'email': this.email
      })
      .subscribe(
        rsp => {
          this.notif.success('Email has been sent');
        },
        error => {
          this.notif.error('Something went wrong');
        }
      );
  }
}

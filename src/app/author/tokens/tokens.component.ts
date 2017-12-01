import { Component, OnInit, ViewEncapsulation } from '@angular/core';
import { TokenService } from '../../token.service';
import { UserService } from '../../user.service';
import { NotificationsService } from 'angular2-notifications';

@Component({
  selector: 'app-tokens',
  templateUrl: './tokens.component.html',
  styleUrls: ['./tokens.component.css'],
  encapsulation: ViewEncapsulation.None
})
export class TokensComponent implements OnInit {
  serviceTokenName = '';
  serviceTokenDescription = '';
  from = '';
  to = '';
  transferAmount = 0;

  constructor(
    private ts: TokenService,
    private us: UserService,
    private notif: NotificationsService
  ) {}

  ngOnInit() {}

  transfer() {
    this.ts.transfer(this.from, this.to, this.transferAmount).then(() => {
      this.us.get();
      this.notif.success('Successfully transferred');
    });
  }

  createToken() {
    this.ts
      .createToken(this.serviceTokenName, this.serviceTokenDescription)
      .then(() => {
        this.us.get();
        this.notif.success('Token successfully created');
      });
  }
}

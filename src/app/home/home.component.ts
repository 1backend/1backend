import { Component, OnInit } from '@angular/core';
import { LoginDialogService } from '../login/login-dialog.service';
import { SessionService } from '../session.service';
import { Title } from '@angular/platform-browser';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  constructor(
    private loginDialog: LoginDialogService,
    private ss: SessionService,
    private title: Title
  ) {}

  ngOnInit() {
    this.title.setTitle('1Backend');
  }

  scroll(el) {
    el.scrollIntoView();
  }
}

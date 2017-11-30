import { Component, OnInit } from '@angular/core';
import { LoginDialogService } from '../login/login-dialog.service';
import { SessionService } from '../session.service';
import { Title } from '@angular/platform-browser';
import { Router } from '@angular/router';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  pricingId = document.getElementById('pricing');
  getStartedId = document.getElementById('getStarted');

  constructor(
    private loginDialog: LoginDialogService,
    private ss: SessionService,
    private title: Title,
    private router: Router
  ) {}

  ngOnInit() {
    this.title.setTitle('1Backend');
  }

  scroll(el) {
    el.scrollIntoView();
  }
}

import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'app';

  constructor(public router: Router) {}

  public options = {
    timeOut: 3000,
    showProgressbar: true,
    pauseOnHover: false,
    clickToClose: false
  };

  pathOnly(): string {
    return this.router.url.split('#')[0].split('?')[0];
  }
}

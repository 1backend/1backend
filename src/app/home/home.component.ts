import { Component, OnInit } from "@angular/core";
import { LoginDialogService } from "../login/login-dialog.service";
import { SessionService } from "../session.service";

@Component({
  selector: "app-home",
  templateUrl: "./home.component.html",
  styleUrls: ["./home.component.css"]
})
export class HomeComponent implements OnInit {
  constructor(
    private loginDialog: LoginDialogService,
    private ss: SessionService
  ) {}

  ngOnInit() {}

  scroll(el) {
    el.scrollIntoView();
  }
}

/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Component } from "@angular/core";
import { IonRouterOutlet } from "@ionic/angular/standalone";



@Component({
  selector: "app-root",
  templateUrl: "./app.component.html",
  styleUrls: ["./app.component.scss"],
  imports: [IonRouterOutlet],
})
export class AppComponent {
  title = "singulatron-angular-app";

  constructor() {}

  ngOnInit(): void {}
}

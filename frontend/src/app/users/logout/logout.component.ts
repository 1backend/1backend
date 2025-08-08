/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Component } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import {
	IonList,
	IonItem,
	IonLabel,
	IonButton,
} from '@ionic/angular/standalone';
import { TranslatePipe } from '../../translate.pipe';
import { TranslateModule } from '@ngx-translate/core';
import { CenteredComponent } from '../../components/centered/centered.component';
import { ChangeDetectionStrategy } from '@angular/core';
import { PageComponent } from '../../components/page/page.component';
import { IconMenuComponent } from '../../components/icon-menu/icon-menu.component';
import { UserService } from '../../services/user.service';

@Component({
	selector: 'app-logout',
	standalone: true,
	imports: [
		IonList,
		IonItem,
		IonLabel,
		IonButton,
		PageComponent,
		IconMenuComponent,
		CenteredComponent,
		FormsModule,
		ReactiveFormsModule,
		TranslateModule,
		TranslatePipe,
	],
	templateUrl: './logout.component.html',
	styleUrl: './logout.component.scss',
	changeDetection: ChangeDetectionStrategy.OnPush,
})
export class LogoutComponent {
	constructor(private userService: UserService) {}

	logout() {
		this.userService.removeToken();
		window.location.pathname = '/';
	}
}

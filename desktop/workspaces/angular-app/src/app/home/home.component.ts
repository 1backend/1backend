/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { ChangeDetectionStrategy, Component } from '@angular/core';
import { TranslatePipe } from '../translate.pipe';
import { TranslateModule } from '@ngx-translate/core';
import { RouterLink } from '@angular/router';
import { NgFor, NgIf } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { IonInput, IonButton } from '@ionic/angular/standalone';
import { CenteredComponent } from '../components/centered/centered.component';
import { PageComponent } from '../components/page/page.component';
import { IconMenuComponent } from '../components/icon-menu/icon-menu.component';
import { UserService } from '../services/user.service';

@Component({
	selector: 'app-home',
	templateUrl: './home.component.html',
	styleUrl: './home.component.css',
	imports: [
		IconMenuComponent,
		PageComponent,
		CenteredComponent,
		IonInput,
		IonButton,
		NgFor,
		NgIf,
		FormsModule,
		RouterLink,
		TranslateModule,
		TranslatePipe,
	],
	changeDetection: ChangeDetectionStrategy.OnPush,
})
export class HomeComponent {
	searchTerm = '';

	buttons = [
		{
			title: 'AI',
			link: '/startup',
		},
		{
			title: 'Users',
			link: '/users',
		},
		{
			title: 'Cluster',
			link: '/nodes',
		},
	];

	constructor(private userService: UserService) {}

	async ngOnInit() {
		this.userService.noop();
	}
}

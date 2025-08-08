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
import { IonInput } from '@ionic/angular/standalone';
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
			description:
				'Use the ChatGPT-inspired local AI platform. No third parties have access to your chats—100% private.',
		},
		{
			title: 'Chat',
			link: '/chat',
			description: 'Chat with any downloaded AI models.',
		},
		{
			title: 'Models',
			link: '/model-explorer',
			description:
				'Download models, set the default model, and pause or resume downloads. View downloaded AI models.',
		},
		{
			title: 'Prompts',
			link: '/prompts',
			description:
				'View and search prompts. Filter by user, model, and more. Cancel prompts or allow them to be retried with exponential backoff.',
		},
		{
			title: 'Users',
			link: '/users',
			description: 'Manage users across all namespaces.',
		},
		//{
		//	title: 'Roles',
		//	link: '/roles',
		//	description: 'View and edit roles and their corresponding permissions.',
		//},
		{
			title: 'Cluster',
			link: '/nodes',
			description: 'Manage your servers.',
		},
		{
			title: 'GPUs',
			link: '/nodes',
			description:
				'Monitor GPU temperatures, utilization, driver versions, and more.',
		},
	];

	constructor(private userService: UserService) {}

	async ngOnInit() {
		this.userService.noop();
	}
}

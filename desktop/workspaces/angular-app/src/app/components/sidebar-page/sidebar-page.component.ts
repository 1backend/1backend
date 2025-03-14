/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import {
	Input,
	Component,
	ViewChild,
	ChangeDetectionStrategy,
	ChangeDetectorRef,
} from '@angular/core';
import { NavController, IonContent } from '@ionic/angular/standalone';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { map } from 'rxjs/operators';
import { NgStyle, NgIf } from '@angular/common';
import { addIcons } from 'ionicons';
import {
	chatboxEllipsesOutline,
	arrowUpCircleOutline,
	createOutline,
	playCircleOutline,
	trashOutline,
	listOutline,
	peopleOutline,
	personAddOutline,
	shirtOutline,
	homeOutline,
} from 'ionicons/icons';

type appGroup = 'ai-group' | 'users-group' | '';

@Component({
	selector: 'b-sidebar-page',
	templateUrl: './sidebar-page.component.html',
	styleUrls: ['./sidebar-page.component.css'],
	imports: [RouterLink, IonContent, NgStyle, NgIf],
	changeDetection: ChangeDetectionStrategy.OnPush,
})
export class SidebarPageComponent {
	id = Math.random().toString(36).slice(7);
	currentPath = '';

	@ViewChild('template', { static: true }) template: any;

	@Input() title: string = '';
	@Input() icon: string = '';
	@Input() noModal: boolean = false;
	@Input() appsModal: boolean = false;
	@Input() noBackButton: boolean = false;
	@Input() themeKey: string = '';

	constructor(
		public navCtrl: NavController,
		private activatedRoute: ActivatedRoute,
		private cd: ChangeDetectorRef
	) {
		addIcons({
			'create-outline': createOutline,
			'trash-outline': trashOutline,
			'play-circle-outline': playCircleOutline,
			'shirt-outline': shirtOutline,
			'person-add-outline': personAddOutline,
			'people-outline': peopleOutline,
			'list-outline': listOutline,
			'chatbox-ellipses-outline': chatboxEllipsesOutline,
			'arrow-up-circle-outline': arrowUpCircleOutline,
			'home-outline': homeOutline,
		});
	}

	ngOnInit() {
		this.activatedRoute.url
			.pipe(map((segments) => segments.join('/')))
			.subscribe((path) => {
				this.currentPath = path;
				this.cd.markForCheck();
			});
	}

	group(): appGroup {
		if (
			this.currentPath === 'startup' ||
			this.currentPath === 'chat' ||
			this.currentPath === 'model-explorer'
		) {
			return 'ai-group';
		}

		if (
			this.currentPath === 'users' ||
			this.currentPath === 'add-user' ||
			this.currentPath === 'roles'
		) {
			return 'users-group';
		}

		return '';
	}

	rout = {
		activeEntry: '',
	};
}

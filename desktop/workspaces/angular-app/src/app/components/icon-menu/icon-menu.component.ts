import {
	Component,
	ChangeDetectionStrategy,
	ChangeDetectorRef,
} from '@angular/core';
import { NavController } from '@ionic/angular/standalone';
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
	codeWorkingOutline
} from 'ionicons/icons';

type appGroup = 'ai-group' | 'users-group' | 'cluster' | '';

@Component({
	selector: 'app-icon-menu',
	imports: [NgStyle, RouterLink, NgIf],
	templateUrl: './icon-menu.component.html',
	styleUrl: './icon-menu.component.scss',
	changeDetection: ChangeDetectionStrategy.OnPush,
})
export class IconMenuComponent {
	currentPath = '';

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
			'code-working-outline': codeWorkingOutline
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
			this.currentPath === 'model-explorer' ||
			this.currentPath === 'prompts'
		) {
			return 'ai-group';
		}

		if (
			this.currentPath === 'users' ||
			this.currentPath === 'add-user' ||
			this.currentPath === 'roles' ||
			this.currentPath === 'logout'
		) {
			return 'users-group';
		}

		if (this.currentPath === 'nodes') {
			return 'cluster';
		}

		return '';
	}

	rout = {
		activeEntry: '',
	};
}

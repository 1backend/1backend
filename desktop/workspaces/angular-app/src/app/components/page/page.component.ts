/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import {
	ContentChildren,
	QueryList,
	Component,
	ChangeDetectionStrategy,
	ChangeDetectorRef,
	Input,
	HostListener,
	AfterContentInit,
	ViewChild,
	ViewContainerRef,
} from '@angular/core';
import { Subscription } from 'rxjs';
import { NgFor } from '@angular/common';
import {
	IonMenu,
	IonApp,
	IonContent,
	IonList,
	IonItem,
	IonIcon,
	IonHeader,
	IonToolbar,
	IonButtons,
	IonMenuButton,
	IonTitle,
} from '@ionic/angular/standalone';
import {
	AsyncPipe,
	NgTemplateOutlet,
	NgStyle,
	NgIf,
	NgClass,
} from '@angular/common';
import { MobileService } from '../../services/mobile.service';
import { FooterService } from '../../services/footer.service';
import { ServerService } from '../../services/server.service';
import { Router, NavigationStart } from '@angular/router';

@Component({
	selector: 'app-page',
	templateUrl: './page.component.html',
	styleUrl: './page.component.scss',
	imports: [
		AsyncPipe,
		IonApp,
		IonContent,
		IonList,
		IonItem,
		IonIcon,
		IonMenu,
		IonHeader,
		IonButtons,
		IonToolbar,
		IonMenuButton,
		IonTitle,
		NgTemplateOutlet,
		NgStyle,
		NgIf,
		NgFor,
		NgClass,
	],
	changeDetection: ChangeDetectionStrategy.OnPush,
})
export class PageComponent implements AfterContentInit {
	/** We give unique IDs to the ion-menu and ion-content elements
	 * so multiple of them can coexists, since the page component is not a singleton one.
	 */
	id: string = '';

	@Input() menuWidth = '90%';
	@Input() menuEnabled = true;
	@Input() columnWidths: string[] = [];
	@Input() mobileColumnWidths: string[] = [];
	@Input() title: string = '';
	@Input() breakpoint: number = 768; // Default breakpoint is 768px

	@ContentChildren('columnContent') columnsContent!: QueryList<any>;
	@ContentChildren('mainContent') mainContent!: QueryList<any>;
	@ViewChild('footerContainer', { read: ViewContainerRef, static: false })
	footerContainer!: ViewContainerRef;

	@ViewChild(IonMenu, { static: true }) menu!: IonMenu;

	columns: any[] = [];
	main: any;

	private subscriptions: Subscription[] = [];

	constructor(
		public mobile: MobileService,
		public footer: FooterService,
		private cd: ChangeDetectorRef,
		private server: ServerService,
		private router: Router
	) {
		this.id = this.server.id('page');
		this.cd.markForCheck();
	}

	@HostListener('window:resize', ['$event'])
	onResize() {
		this.mobile.setMobileStatus(window.innerWidth < this.breakpoint);
	}

	ngOnInit() {
		this.subscriptions.push(
			this.router.events.subscribe((event) => {
				if (event instanceof NavigationStart) {
					// Footers get removed here, but they are inserted by
					// components because they might want to do plumbing
					// like setting up @Output() subscriptions.
					this.footer.removeFooterComponent();
					this.cd.markForCheck();
				}
			}),
			this.mobile.isMobile$.subscribe((isMobile) => {
				if (isMobile) {
					return;
				}

				this.footer.removeFooterComponent();
			})
		);

		this.mobile.setMobileStatus(window.innerWidth < this.breakpoint);

		this.footer.footerComponent$.subscribe((componentReference) => {
			// @todo incredibly nasty hack
			// #footerContainer in the html template is not available at this
			// point because the *ngIf="footerComponent" is just being evaluated
			setTimeout(() => {
				if (componentReference && this.footerContainer) {
					this.footerContainer.insert(componentReference.hostView);
				}
			}, 1);
		});
	}

	ionViewDidLeave() {
		for (const s of this.subscriptions) {
			s.unsubscribe();
		}
	}

	ngAfterContentInit(): void {
		this.columns = this.columnsContent.toArray();
		this.main = this.mainContent.first;
	}

	getColumnWidth(index: number): string {
		if (this.mobile.getMobileStatus() && this.mobileColumnWidths[index]) {
			return this.mobileColumnWidths[index];
		}
		return this.columnWidths[index] || 'auto';
	}

	getBreakpointQuery(): string {
		return `(min-width: ${this.breakpoint}px)`;
	}

	toggleMenu() {
		this.menu.toggle();
	}
}

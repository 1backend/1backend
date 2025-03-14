import { Injectable,  ComponentRef } from '@angular/core';
import { ReplaySubject } from 'rxjs';
import { Router, ActivatedRoute, NavigationStart } from '@angular/router';
import { map, filter } from 'rxjs/operators';

@Injectable({
	providedIn: 'root',
})
export class FooterService {
	private hasFooter = false;

	// eslint-disable-next-line
	private footerComponentRef: ComponentRef<any> | null = null;
	footerComponentSubject = new ReplaySubject<ComponentRef<any> | null>(1);
	footerComponent$ = this.footerComponentSubject.asObservable();

	constructor(
		private activatedRoute: ActivatedRoute,
		private router: Router
	) {
		// Since ionic lifecycle hooks dont seem to
		// be triggering properly - nor ngOnDestroy -
		// we need to do this hack here.
		// We always remove the footer at NavigationStart and then
		// components can set their own footers when they detect NavigationEnd
		// on their path. A huge drawback is that components must know
		// their path.
		this.router.events
			.pipe(
				filter((event) => event instanceof NavigationStart),
				map(() => this.activatedRoute),
				map((route) => {
					while (route.firstChild) route = route.firstChild;
					return route;
				}),
				filter((route) => route.outlet === 'primary'),
				map((route) => route.snapshot.url.join('/'))
			)
			.subscribe(() => {
				this.removeFooterComponent();
			});
	}

	hasFooterComponent(): boolean {
		return this.hasFooter;
	}

	updateFooterComponent(componentInstance: ComponentRef<any> | null) {
		this.hasFooter = true;
		this.footerComponentSubject.next(componentInstance);
	}

	removeFooterComponent() {
		this.hasFooter = false;
		// eslint-disable-next-line
		this.footerComponentSubject.next(null);

		if (this.footerComponentRef) {
			this.footerComponentRef.destroy();
			// eslint-disable-next-line
			this.footerComponentRef = null;
		}
	}
}

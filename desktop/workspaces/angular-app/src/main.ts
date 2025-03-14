// import { provideExperimentalZonelessChangeDetection } from '@angular/core';
import { enableProdMode } from '@angular/core';

import { environment } from './environments/environment';
import { AppComponent } from './app/app.component';
import { provideAnimationsAsync } from '@angular/platform-browser/animations/async';
import { bootstrapApplication } from '@angular/platform-browser';
import {
	provideHttpClient,
	withInterceptorsFromDi,
	withFetch,
} from '@angular/common/http';
import { OPENORCH_SERVICE_CONFIG } from './app/services/server.service';
import { API_SERVICE_CONFIG } from './app/api.service';
import { provideRouter, RouteReuseStrategy } from '@angular/router';
import { routes } from './app/app-routing.module';
import {
	IonicRouteStrategy,
	provideIonicAngular,
} from '@ionic/angular/standalone';

if (environment.production) {
	enableProdMode();
}

bootstrapApplication(AppComponent, {
	providers: [
		{ provide: RouteReuseStrategy, useClass: IonicRouteStrategy },
		provideIonicAngular(),
		provideHttpClient(withFetch(), withInterceptorsFromDi()),
		provideRouter(routes),
		// provideExperimentalZonelessChangeDetection(),
		{
			provide: OPENORCH_SERVICE_CONFIG,
			useValue: { env: environment },
		},
		{
			provide: API_SERVICE_CONFIG,
			useValue: { env: environment },
		},
		provideAnimationsAsync(),
	],
	// eslint-disable-next-line
}).catch((error) => console.error(error));

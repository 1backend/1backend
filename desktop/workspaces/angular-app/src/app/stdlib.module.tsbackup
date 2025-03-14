/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { ModuleWithProviders, NgModule } from '@angular/core';

import { CommonModule } from '@angular/common';
import { IonicModule } from '@ionic/angular';
import {
	provideHttpClient,
	withInterceptorsFromDi,
} from '@angular/common/http';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import {
	ApiService,
	API_SERVICE_CONFIG,
	ApiServiceConfig,
} from './api.service';

import { CookieService } from 'ngx-cookie-service';
import { TranslatePipe, TRANSLATE_OBJECT } from './translate.pipe';
import translations from './i18n.json';
import { RouterModule } from '@angular/router';
import { CenteredComponent } from './components/centered/centered.component';
import { SidebarPageComponent } from './components/sidebar-page/sidebar-page.component';

export interface StdlibModuleConfig {
	apiServiceConfig: ApiServiceConfig;
}

@NgModule({
    exports: [
        IonicModule,
        FormsModule,
        ReactiveFormsModule,
        TranslatePipe,
        CenteredComponent,
        CommonModule,
        SidebarPageComponent,
    ],
    imports: [
        CommonModule,
        FormsModule,
        ReactiveFormsModule,
        IonicModule,
        RouterModule,
        TranslatePipe, CenteredComponent, SidebarPageComponent,
    ],
    providers: [
        {
            provide: TRANSLATE_OBJECT,
            useValue: translations,
        },
        TranslatePipe,
        provideHttpClient(withInterceptorsFromDi()),
    ],
})
export class StdlibModule {
	static forRoot(
		config: StdlibModuleConfig
	): ModuleWithProviders<StdlibModule> {
		return {
			ngModule: StdlibModule,
			providers: [
				{
					provide: API_SERVICE_CONFIG,
					useValue: config.apiServiceConfig,
				},
				ApiService,
				CookieService,
			],
		};
	}
}

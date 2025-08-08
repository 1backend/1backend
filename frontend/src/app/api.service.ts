/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Injectable, InjectionToken } from '@angular/core';

import { BehaviorSubject } from 'rxjs';

export interface Event {
	type: string;
}


export interface Environment {
	production: boolean;
	brandName: string;
	shortBrandName: string;
}

export interface ApiServiceConfig {
	env: Environment;
}
export const API_SERVICE_CONFIG = new InjectionToken<ApiServiceConfig>(
	'ApiServiceConfig'
);

export interface Thread {
	id: string;
	name: string;
	messages?: Array<Message>;
}

export const defaultThreadName = 'New chat';

export interface Message {
	id: string;
	content?: string;
	userId?: string;
}

@Injectable({
	providedIn: 'root',
})
export class ApiService {
	private locale = 'en';

	public firehose: BehaviorSubject<Event> = new BehaviorSubject<Event>({
		type: 'noop',
	});

	constructor() {}

	public setLocale(s: string) {
		this.locale = s;
	}

	public getLocale(): string {
		return this.locale;
	}
}

export interface ReadByWebsitesRequest {
	host: string;
}

export interface VersionResponse {
	windows?: Version;
	linux?: Version;
	mac?: Version;
}

export interface Version {
	version?: string;
	downloadPageURL?: string;
	downloadURL?: string;
	releaseDate?: Date | string;
	changeLog?: string;
}

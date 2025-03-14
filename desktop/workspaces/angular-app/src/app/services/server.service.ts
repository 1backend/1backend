/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Injectable, Inject, InjectionToken } from '@angular/core';
import { CookieService } from 'ngx-cookie-service';
import Sonyflake from 'sonyflake';

const base62 = '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz';
const sonyflake = new Sonyflake();

export interface Environment {
	production: boolean;
	brandName: string;
	shortBrandName: string;
	backendAddress: string;
	localPromptAddress: string;
	serverAddress: string;
}

export interface ServerServiceConfig {
	env: Environment;
}
export const OPENORCH_SERVICE_CONFIG =
	new InjectionToken<ServerServiceConfig>('ServerServiceConfig');

@Injectable({
	providedIn: 'root',
})
export class ServerService {
	public config: ServerServiceConfig;

	constructor(
		private cs: CookieService,
		@Inject(OPENORCH_SERVICE_CONFIG) config: ServerServiceConfig
	) {
		this.config = config;
	}

	token(): string {
		return this.cs.get('the_token');
	}

	addr(): string {
		return this.config.env.serverAddress;
	}

	id(prefix: string): string {
		const numberString = sonyflake.nextId();
		const number = BigInt(numberString);

		if (number === BigInt(0)) {
			return `${prefix}_0`;
		}

		let result = '';
		let number_ = number;

		while (number_ > 0) {
			const remainder = Number(number_ % BigInt(62));
			number_ = number_ / BigInt(62);
			result = base62[remainder] + result;
		}

		return `${prefix}_${result}`;
	}
}

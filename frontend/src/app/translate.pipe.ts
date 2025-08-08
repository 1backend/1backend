/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Pipe, PipeTransform } from '@angular/core';
import { ApiService } from './api.service';

// key to language to translation
export type TranslateObject = { [key: string]: { [key: string]: string } };

@Pipe({
	name: 'translate',
	pure: true,
	standalone: true,
})
export class TranslatePipe implements PipeTransform {
	constructor(
		private apiService: ApiService,
	) {}

	transform(key: string, 	translations?: TranslateObject): string {
		if (!translations) {
			return key
		}
		if (!translations[key]) {
			return key;
		}
		const lang = this.apiService.getLocale();
		if (lang) {
			return translations[key][lang];
		}
		if (translations[key]['en']) {
			return translations[key]['en'];
		}
		return key;
	}
}

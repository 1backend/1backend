/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import {
	Component,
	ChangeDetectionStrategy,
	ChangeDetectorRef,
} from '@angular/core';
import { NgFor, NgIf } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { IonButton } from '@ionic/angular/standalone';
import { PromptService } from '../services/prompt.service';
import { UserService } from '../services/user.service';
import { first } from 'rxjs';
import { PageComponent } from '../components/page/page.component';
import { IconMenuComponent } from '../components/icon-menu/icon-menu.component';
import { CenteredComponent } from '../components/centered/centered.component';
import { PromptComponent } from './prompt/prompt.component';
import { QueryParser } from '../services/query.service';
import {
	DatastoreFilter,
	PromptSvcPrompt as Prompt,
	PromptSvcListPromptsRequest,
} from '@openorch/client';

import { ActivatedRoute, Router } from '@angular/router';

@Component({
	selector: 'app-prompts',
	imports: [
		IonButton,
		CenteredComponent,
		PageComponent,
		IconMenuComponent,
		NgFor,
		NgIf,
		FormsModule,
		PromptComponent,
	],
	changeDetection: ChangeDetectionStrategy.OnPush,
	templateUrl: './prompts.component.html',
	styleUrls: ['./prompts.component.scss'],
})
export class PromptsComponent {
	prompts: Prompt[] = [];
	after: any;
	request = {
		statuses: [
			'scheduled',
			'running',
			'completed',
			'errored',
			'abandoned',
			'canceled',
		],
		desc: true,
	};
	count = 0;
	searchTerm = '';
	queryParser: QueryParser;

	constructor(
		private cd: ChangeDetectorRef,
		private userService: UserService,
		private promptService: PromptService,
		private activatedRoute: ActivatedRoute,
		private router: Router
	) {
		this.queryParser = new QueryParser();
		this.queryParser.defaultConditionFunc = (value: any): DatastoreFilter => {
			return {
				op: 'containsSubstring',
				fields: ['modelId'],
				jsonValues: JSON.stringify([value]),
			};
		};

		this.userService.user$.pipe(first()).subscribe(() => {
			this.initializeOnLogin();
		});
	}

	private initializeOnLogin() {
		this.activatedRoute.queryParams.subscribe((parameters) => {
			this.searchTerm =
				this.queryParser.convertQueryParamsToSearchTerm(parameters);

			this.fetchPrompts();
			this.cd.markForCheck();
		});
	}

	public redirect() {
		const query = this.queryParser.parse(this.searchTerm);
		const kv = filtersToKeyValue(
			query.filters
				? query.filters.filter((v) => {
						return v.fields?.includes('modelId');
					})
				: []
		);

		if (Object.keys(kv)?.length) {
			this.router.navigate([], {
				queryParams: kv,
			});
			return;
		}

		if (this.searchTerm) {
			this.router.navigate([], {
				queryParams: { search: this.searchTerm },
			});
			return;
		}

		this.router.navigate([], {
			queryParams: {},
		});
	}

	public async fetchPrompts() {
		const query = this.queryParser.parse(this.searchTerm);
		query.count = true;
		query.filters = query.filters || [];

		if (!query.filters.some((f) => f.fields?.includes('status'))) {
			query.filters.push({
				fields: ['status'],
				jsonValues: JSON.stringify(this.request.statuses),
				op: 'isInList',
			});
		}
		query.orderBys = [{ field: 'createdAt', desc: true }];

		const request: PromptSvcListPromptsRequest = {
			query: query,
		};

		if (this.after) {
			request.query!.jsonAfter = JSON.stringify([this.after]);
		}

		const response = await this.promptService.promptList(request);

		// eslint-disable-next-line
		if (response.prompts && this.after) {
			this.prompts = [...this.prompts, ...response.prompts];
		} else if (response.prompts) {
			this.prompts = response.prompts;
		} else {
			this.prompts = [];
		}

		this.count = response.count || 0;

		// eslint-disable-next-line
		if (response.after && response.after != null) {
			this.after = response.after;
		} else {
			this.after = undefined;
		}

		this.cd.markForCheck();
	}

	async loadMoreData() {
		if (!this.after) {
			console.log('No more prompts to load');
			return;
		}
		await this.fetchPrompts();
	}
}

function filtersToKeyValue(filters: DatastoreFilter[]): {
	[key: string]: any;
} {
	if (!filters) {
		return {};
	}
	const object: { [key: string]: any } = {};

	for (const filter of filters) {
		object[filter.fields![0]] = JSON.parse(filter.jsonValues!)[0];
	}

	return object;
}

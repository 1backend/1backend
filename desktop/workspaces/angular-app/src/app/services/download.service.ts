/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Injectable } from '@angular/core';
import { ServerService } from './server.service';
import { FirehoseService } from './firehose.service';
import {
	ReplaySubject,
	first,
	filter,
	throttleTime,
	exhaustMap,
	from,
	catchError,
	of,
} from 'rxjs';
import { UserService } from './user.service';
import {
	FileSvcApi,
	Configuration,
	FileSvcDownload as Download,
	FileSvcDownloadsResponse,
} from '@openorch/client';

export interface FileSvcConfig {
	downloadFolder: string;
}

export interface DownloadStatusChangeEvent {
	allDownloads: Download[];
}

@Injectable({
	providedIn: 'root',
})
export class DownloadService {
	downloadService!: FileSvcApi;

	onFileDownloadStatusSubject = new ReplaySubject<DownloadStatusChangeEvent>(1);
	onFileDownloadStatus$ = this.onFileDownloadStatusSubject.asObservable();

	constructor(
		private server: ServerService,
		private firehoseService: FirehoseService,
		private userService: UserService
	) {
		this.init();
		this.userService.user$.pipe(first()).subscribe(() => {
			this.downloadService = new FileSvcApi(
				new Configuration({
					basePath: this.server.addr(),
					apiKey: this.server.token(),
				})
			);
			this.loggedInInit();
		});
	}

	async init() {
		this.firehoseService.firehoseEvent$
			.pipe(
				filter((event) => event.name === 'downloadStatusChange'),
				throttleTime(500),
				exhaustMap(() =>
					from(this.downloadList()).pipe(
						catchError((error) => {
							console.error('Error fetching download list:', error);
							return of({ downloads: [] });
						})
					)
				)
			)
			.subscribe((rsp) => {
				this.onFileDownloadStatusSubject.next({
					allDownloads: rsp?.downloads || [],
				});
			});
	}

	async loggedInInit() {
		try {
			const rsp = await this.downloadList();
			this.onFileDownloadStatusSubject.next({
				allDownloads: rsp?.downloads as Download[],
			});
		} catch (error) {
			console.error('Error in pollFileDownloadStatus', {
				error: JSON.stringify(error),
			});
		}
	}

	async downloadDo(url: string) {
		this.downloadService.downloadFile({
			body: {
				url: url,
			},
		});
	}

	async downloadPause(url: string) {
		this.downloadService.pauseDownload({
			url: url,
		});
	}

	async downloadList(): Promise<FileSvcDownloadsResponse> {
		return this.downloadService.listDownloads();
	}
}

/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Component, Input, Output, EventEmitter } from '@angular/core';
import { Subscription, throttleTime } from 'rxjs';
import { DownloadService } from '../services/download.service';
import { FileSvcDownload as Download } from '@openorch/client';
import { TranslatePipe } from '../translate.pipe';
import { TranslateModule } from '@ngx-translate/core';
import { IonSpinner, IonIcon } from '@ionic/angular/standalone';
import { NgIf, DecimalPipe } from '@angular/common';
import { addIcons } from 'ionicons';
import { pauseOutline, playOutline } from 'ionicons/icons';

@Component({
	selector: 'app-downloading',
	templateUrl: './downloading.component.html',
	styleUrl: './downloading.component.scss',
	imports: [
		IonSpinner,
		IonIcon,
		NgIf,
		DecimalPipe,
		TranslateModule,
		TranslatePipe,
	],
})
export class DownloadingComponent {
	@Input() url!: string;
	@Output() downloadStatusChange = new EventEmitter<Download>();

	details!: Download;

	constructor(private downloadService: DownloadService) {
		addIcons({
			play: playOutline,
			pause: pauseOutline,
		});
	}

	subscriptions: Subscription[] = [];

	ngOnInit() {
		this.subscriptions.push(
			this.downloadService.onFileDownloadStatus$.subscribe((data) => {
				const d = data.allDownloads?.find((d) => {
					return d.url == this.url;
				});

				if (!d) {
					return;
				}

				this.details = d;
				this.downloadStatusChange.emit(this.details);
			}),
			this.downloadService.onFileDownloadStatus$
				.pipe(throttleTime(10 * 1000))
				.subscribe((data) => {
					const d = data.allDownloads?.find((d) => {
						return d.url == this.url;
					});
					if (!d) {
						return;
					}
					if (d.status == 'inProgress') {
						console.debug('Download is in progress', {
							url: d.url,
							status: d.status,
							progress: d.progress,
							error: d.error,
						});
					}
				})
		);
	}

	ionViewWillLeave() {
		for (const v of this.subscriptions) {
			v.unsubscribe();
		}
	}

	pauseDownload() {
		this.downloadService.downloadPause(this.url);
	}

	resumeDownload() {
		this.downloadService.downloadDo(this.url);
	}

	isValid(): boolean {
		try {
			new URL(this.url);
			return true;
		} catch {}
		return false;
	}
}

/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Component, OnInit, ViewChild, ElementRef } from '@angular/core';
import { ElectronIpcService } from '../services/electron-ipc.service';
import { WindowApiConst } from 'shared-lib';
import { ElectronAppService } from '../services/electron-app.service';
import { combineLatest, Subscription } from 'rxjs';
import { DownloadService, FileSvcConfig } from '../services/download.service';
import { ModelService, ModelSvcConfig } from '../services/model.service';
import { ModelSvcModel } from '@openorch/client';
import { ContainerService } from '../services/container.service';
import { ConfigService } from '../services/config.service';
import {
	FileSvcDownload as Download,
	ConfigSvcConfig as Config,
} from '@openorch/client';
import { TranslatePipe } from '../translate.pipe';
import { TranslateModule } from '@ngx-translate/core';
import { DownloadingComponent } from '../downloading/downloading.component';
import { RouterLink } from '@angular/router';
import { NgIf, NgStyle, AsyncPipe } from '@angular/common';
import { CenteredComponent } from '../components/centered/centered.component';
import { PageComponent } from '../components/page/page.component';
import { IconMenuComponent } from '../components/icon-menu/icon-menu.component';
import {
	IonList,
	IonItem,
	IonIcon,
	IonLabel,
	IonButton,
	IonSpinner,
} from '@ionic/angular/standalone';
import { addIcons } from 'ionicons';
import {
	laptopOutline,
	chevronDown,
	chevronForward,
	cubeOutline,
	hardwareChipOutline,
} from 'ionicons/icons';

@Component({
	selector: 'app-startup',
	templateUrl: './startup.component.html',
	styleUrl: './startup.component.scss',
	imports: [
		IonList,
		IonItem,
		IonIcon,
		IonLabel,
		IonButton,
		IonSpinner,
		IconMenuComponent,
		PageComponent,
		CenteredComponent,
		NgIf,
		NgStyle,
		RouterLink,
		DownloadingComponent,
		AsyncPipe,
		TranslateModule,
		TranslatePipe,
	],
})
export class StartupComponent implements OnInit {
	@ViewChild('logContainer') private logContainer!: ElementRef;
	log = 'Installation logs will be streamed here. Please wait...\n';
	scrollToBottom(): void {
		try {
			this.logContainer.nativeElement.scrollTop =
				this.logContainer.nativeElement.scrollHeight;
		} catch {}
	}

	models: ModelSvcModel[] = [];
	allIsWell = false;
	isDownloading = false;
	downloaded = false;
	restartIsRequired = false;

	downloadFolder: string = '';

	showSections = {
		model: false,
		dependencies: false,
		starting: false,
	};

	toggleSection(section: string) {
		type ShowSectionsKeys = keyof typeof this.showSections;
		this.showSections[section as ShowSectionsKeys] =
			!this.showSections[section as ShowSectionsKeys];
	}

	constructor(
		private ipcService: ElectronIpcService,
		public lapi: ElectronAppService,
		public configService: ConfigService,
		public downloadService: DownloadService,
		public containerService: ContainerService,
		public modelService: ModelService
	) {
		addIcons({
			'laptop-outline': laptopOutline,
			'chevron-down': chevronDown,
			'chevron-forward': chevronForward,
			'cube-outline': cubeOutline,
			'hardware-chip-outline': hardwareChipOutline,
		});
	}

	handleDownloadStatus(data: Download) {
		this.isDownloading = data.status == 'inProgress' || data.status == 'paused';
		this.downloaded = data.status == 'completed';
	}

	selectedModelName(cu: Config | null): string {
		if (!cu) {
			return '';
		}

		const model = this.models?.find(
			(v) => v.id == (cu?.data['model-svc'] as ModelSvcConfig)?.currentModelId
		);
		const displayName = [model?.name, model?.flavour, model?.version].join(' ');
		return displayName;
	}

	selectedModel(cu: Config | null): ModelSvcModel | undefined {
		if (!cu) {
			return;
		}
		return this.models?.find(
			(v) => v.id == (cu?.data['model-svc'] as ModelSvcConfig)?.currentModelId
		);
	}

	ionViewWillLeave() {
		for (const v of this.subscriptions) {
			v.unsubscribe();
		}
	}

	private subscriptions: Subscription[] = [];

	async ngOnInit(): Promise<void> {
		this.subscriptions.push(
			this.lapi.onRuntimeInstallLog$.subscribe((data) => {
				if (data == this.log) {
					return;
				}

				for (let l of data.replace(this.log, '').trim().split('\n')) {
					l = l?.trim();
					if (l) {
						console.log('Install log: ' + l);
					}
				}

				this.log = data;
				if (
					this.log?.includes('RESTART REQUIRED') ||
					this.log?.includes('restart is required')
				) {
					this.restartIsRequired = true;
				}

				this.scrollToBottom();
			})
		);
		this.models = await this.modelService.getModels();

		this.subscriptions.push(
			this.lapi.onFolderSelect$.subscribe((data) => {
				this.downloadFolder = data.location;
			})
		);

		let assetsReady = false;
		this.subscriptions.push(
			this.modelService.onModelCheck$.subscribe((data) => {
				if (data.assetsReady === undefined) {
					return;
				}
				if (data.assetsReady !== assetsReady) {
					assetsReady = data.assetsReady;
				}
			})
		);

		combineLatest([
			this.containerService.onContainerInfo$,
			this.modelService.onModelCheck$,
		]).subscribe(([dockerInfo, modelCheck]) => {
			if (this.allIsWell) {
				return;
			}
			if (!dockerInfo.hasDocker) {
				this.showSections.dependencies = true;
			} else if (modelCheck.assetsReady == false) {
				this.showSections.model = true;
			} else {
				this.showSections.starting = true;
			}
		});

		this.subscriptions.push(
			this.modelService.onModelLaunch$.subscribe(async () => {
				if (this.allIsWell) {
					return;
				}
				this.showSections.starting = false;
				console.log('Received model launch event');
				this.allIsWell = true;
			})
		);
	}

	openFolderSelect() {
		this.ipcService.send(WindowApiConst.SELECT_FOLDER_REQUEST, {});
	}

	async download() {
		const config = this.configService.lastConfig;
		const modelId = (config as unknown as ConfigData).data['model-svc']
			.currentModelId;
		if (!modelId) {
			throw 'Model id is empty';
		}
		const model = this.models?.find((v) => v.id == modelId);
		if (!model) {
			throw `Cannot find model with id ${modelId}`;
		}

		for (const asset of model.assets!) {
			this.downloadService.downloadDo(asset.url);
		}
	}

	isRuntimeInstalling = false;
	async installRuntime() {
		this.ipcService.send(WindowApiConst.INSTALL_RUNTIME_REQUEST, {});
		this.isRuntimeInstalling = true;
	}

	hasModelAsset(model?: ModelSvcModel): boolean {
		return model?.assets?.find((a) => a.envVarKey == 'MODEL') != undefined;
	}

	modelAssetUrl(model?: ModelSvcModel): string {
		return model?.assets?.find((a) => a.envVarKey == 'MODEL')?.url || '';
	}
}

interface ConfigData {
	data: {
		'model-svc': ModelSvcConfig;
		'file-svc': FileSvcConfig;
	};
}

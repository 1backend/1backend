/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Injectable } from '@angular/core';
import { ServerService } from './server.service';
import { ContainerService } from './container.service';
import { ReplaySubject, combineLatest } from 'rxjs';
import {
	ModelSvcApi,
	ModelSvcModel as Model,
	ModelSvcListModelsResponse as GetModelsResponse,
	Configuration,
	ModelSvcStatusResponse,
	ModelSvcPlatform as Platform
} from '@openorch/client';
import {
	OnModelLaunch,
	OnModelCheck,
} from 'shared-lib/models/event-request-response';

export interface ModelSvcConfig {
	currentModelId: string;
}

@Injectable({
	providedIn: 'root',
})
export class ModelService {
	private modelService!: ModelSvcApi;

	private initInProgress: boolean = false;

	private onModelCheckSubject = new ReplaySubject<OnModelCheck>(1);
	/** Emitted any time when the currently selected model is checked */
	public onModelCheck$ = this.onModelCheckSubject.asObservable();

	private onModelLaunchSubject = new ReplaySubject<OnModelLaunch>(1);
	/** Emitted when the model is launched and available shortly */
	public onModelLaunch$ = this.onModelLaunchSubject.asObservable();

	private onModelReadySubject = new ReplaySubject<OnModelReady>(1);
	public onModelReady$ = this.onModelReadySubject.asObservable();

	private platformsSubject = new ReplaySubject<Platform[]>(1);
	public platforms$ = this.platformsSubject.asObservable();


	constructor(
		private server: ServerService,
		private dockerService: ContainerService
	) {
		this.modelService = new ModelSvcApi(
			new Configuration({
				basePath: this.server.addr(),
				apiKey: this.server.token(),
			})
		);

		this.initPlatforms()

		// @todo nothing to trigger model start so we resolve to polling
		setInterval(() => {
			this.init();
		}, 2000);

		this.listenToModelReady();
	}

	models: Model[] = [];

	async getModels(): Promise<Model[]> {
		if (this.models?.length) {
			return this.models;
		}

		const rsp: GetModelsResponse = await this.modelService.listModels();
		return rsp.models!;
	}

	async initPlatforms(){
		const platformsRsp = await this.modelService.listPlatforms()
		this.platformsSubject.next(platformsRsp.platforms)
	}

	private listenToModelReady(): void {
		combineLatest([
			this.dockerService.onContainerInfo$,
			this.onModelCheck$,
		]).subscribe(([dockerInfo, modelCheck]) => {
			if (dockerInfo.hasDocker && modelCheck.assetsReady) {
				this.onModelReadySubject.next({ modelReady: true });
			}
		});
	}

	async init() {
		try {
			if (this.initInProgress) {
				return;
			}
			this.initInProgress = true;

			this.models = await this.getModels();
			const rsp = await this.modelStatus();

			this.onModelCheckSubject.next({
				assetsReady: rsp?.status?.assetsReady,
			});

			if (rsp?.status?.running) {
				this.onModelLaunchSubject.next({});
			}

			if (rsp?.status?.assetsReady) {
				await this.modelStart();
			}
		} catch (error) {
			console.log(error);
			console.error('Error in model.service init', {
				error: JSON.stringify(error),
			});
		} finally {
			this.initInProgress = false;
		}
	}

	async modelStatus(modelId?: string): Promise<ModelSvcStatusResponse> {
		if (modelId) {
			return await this.modelService.getModelStatus({
				modelId: modelId,
			});
		}

		return this.modelService.getDefaultModelStatus();
	}

	async modelStart(modelId?: string): Promise<ModelStartResponse> {
		if (modelId) {
			return this.modelService.startModel({
				modelId: modelId,
			});
		}

		return this.modelService.startDefaultModel();
	}

	async makeDefault(id: string) {
		this.modelService.makeDefault({
			modelId: id,
		});
	}
}

export interface OnModelReady {
	modelReady: boolean;
}

// eslint-disable-next-line
interface ModelStartResponse {}

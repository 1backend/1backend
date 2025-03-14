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
import { first } from 'rxjs';
import { UserService } from './user.service';
import {
	Configuration,
	DataSvcApi,
	DatastoreFilter,
	DataSvcCreateObjectRequest as CreateObjectRequest,
	DataSvcQueryRequest as QueryRequest,
	DataSvcQueryResponse as QueryResponse,
	DataSvcObject as DataObject,
	DataSvcUpdateObjectsRequest as UpdateObjectsRequest,
	// DataSvcUpdateObjectsResponse as UpdateObjectsResponse,
	DataSvcUpsertObjectRequest as UpsertObjectRequest,
	DataSvcCreateObjectResponse,
	DataSvcUpsertObjectResponse,
	// DataSvcDeleteObjectRequest as DeleteObjectRequest,
	// DataSvcDeleteObjectResponse as DeleteObjectResponse,
} from '@openorch/client';

@Injectable({
	providedIn: 'root',
})
export class DataService {
	dataService!: DataSvcApi;

	constructor(
		private server: ServerService,
		private userService: UserService,
		private firehoseService: FirehoseService
	) {
		this.userService.user$.pipe(first()).subscribe(() => {
			this.init();
			this.dataService = new DataSvcApi(
				new Configuration({
					apiKey: this.server.token(),
					basePath: this.server.addr(),
				})
			);
		});
	}

	async init() {
		this.firehoseService.firehoseEvent$.subscribe(async (event) => {
			switch (event.name) {
			}
			return;
		});
	}

	async create(
		table: string,
		object: DataObject
	): Promise<DataSvcCreateObjectResponse> {
		object.table = table;
		const request: CreateObjectRequest = {
			object: object,
		};

		return this.dataService.createObject({ body: request });
	}

	async find(
		table: string,
		filters: DatastoreFilter[]
	): Promise<QueryResponse> {
		const request: QueryRequest = {
			table: table,
			query: {
				filters: filters,
			},
		};

		return this.dataService.queryObjects({ body: request });
	}

	async upsert(
		table: string,
		object: DataObject
	): Promise<DataSvcUpsertObjectResponse> {
		object.table = table;
		const request: UpsertObjectRequest = {
			object: object,
		};

		return this.dataService.upsertObject({
			objectId: object.id!,
			body: request,
		});
	}

	async update(
		table: string,
		filters: DatastoreFilter[],
		object: DataObject
	): Promise<object> {
		const request: UpdateObjectsRequest = {
			table: table,
			filters: filters,
			object: object,
		};

		return this.dataService.updateObjects({
			body: request,
		});
	}

	async delete(table: string, filters: DatastoreFilter[]): Promise<any> {
		console.log(table, filters);

		// const request = {
		// 	table: table,
		// 	filters: filters,
		// };
		//eturn this.dataService.deleteObjects({
		//	body: request,
		//);
	}
}

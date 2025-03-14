/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Injectable } from '@angular/core';
import { ServerService } from './server.service';
import { Observable, Subject, first } from 'rxjs';
import { UserService } from './user.service';
import { catchError } from 'rxjs/operators';

@Injectable({
	providedIn: 'root',
})
export class FirehoseService {
	private firehoseEventSubject = new Subject<FirehoseEvent>();
	public firehoseEvent$ = this.firehoseEventSubject.asObservable();

	constructor(
		private server: ServerService,
		private userService: UserService
	) {
		this.userService.user$.pipe(first()).subscribe(() => {
			this.init();
		});
	}

	async init() {
		this.firehoseSubscribe().subscribe((event) => {
			this.firehoseEventSubject.next(event);
		});
	}

	private resubCount = 0;
	private firehoseSubscribe(): Observable<FirehoseEvent> {
		return new Observable<FirehoseEvent>((observer) => {
			const controller = new AbortController();
			const { signal } = controller;

			const subscribe = () => {
				console.info('Subscribing to the firehose');

				const uri =
					this.server.config.env.serverAddress +
					'/firehose-svc/events/subscribe';

				const headers = {
					Authorization: 'Bearer ' + this.userService.getToken(),
					'Content-Type': 'application/json',
				};

				fetch(uri, {
					method: 'GET',
					headers: headers,
					signal: signal,
				})
					.then((response) => {
						if (!response || !response.body) {
							observer.error(`Response is empty`);
							return;
						}
						if (!response.ok) {
							observer.error(`HTTP error! status: ${response.status}`);
							return;
						}
						const reader = response.body.getReader();
						return new ReadableStream({
							start(controller) {
								function push() {
									reader
										.read()
										.then(({ done, value }) => {
											if (done) {
												controller.close();
												observer.complete();
												return;
											}
											// Convert the Uint8Array to string
											const text = new TextDecoder().decode(value);
											const lines = text.split('\n');
											for (const line of lines) {
												const trimmedLine = line.trim();

												if (
													trimmedLine === '' ||
													trimmedLine === 'data: ' ||
													trimmedLine === 'data: [DONE]'
												) {
													// Skip empty lines, lines containing only 'data: ', or "[DONE]" markers
													continue;
												}

												const cleanedText = trimmedLine
													.replaceAll(/^data: /gm, '')
													.trim();

												try {
													const json = JSON.parse(cleanedText);
													observer.next(json);
												} catch (error) {
													console.error(
														'Error parsing firehose response chunk JSON',
														{
															error: error,
															promptResponseChunk: cleanedText,
														}
													);
													// Decide how we want to handle parsing errors.
													// For continuous streaming, we might not want to call observer.error() here
													// unless it's a critical error that requires stopping the stream.
												}
											}

											// Call push again outside the loop to continue reading
											push();
										})
										.catch((error) => {
											if (
												error instanceof Error &&
												error.message.includes('BodyStreamBuffer was aborted')
											) {
												// we ignore this because this is normal
											} else {
												console.error('Error reading from stream', {
													error: JSON.stringify(error),
												});

												observer.error(error);
												controller.error(error);
											}
											observer.error(error);
											controller.error(error);
										});
								}
								push();
							},
						});
					})
					.catch((error) => {
						observer.error(error);
					});
			};

			sleep(this.resubCount * 20).then(() => {
				subscribe();
			});

			return () => {
				controller.abort(); // This ensures fetch is aborted when unsubscribing
			};
		}).pipe(
			catchError((error) => {
				console.error('Firehose subscription error', {
					error: JSON.stringify(error),
				});

				this.resubCount++;
				return this.firehoseSubscribe();
			})
		);
	}
}

export interface FirehoseEvent {
	name: string;
	data: any;
}

function sleep(ms: number) {
	return new Promise((resolve) => setTimeout(resolve, ms));
}

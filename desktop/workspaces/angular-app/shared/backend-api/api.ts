/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import * as zlib from 'zlib';
import brotliCompress from 'compress-brotli';

// BROTLI_ZLIB doesnt seem to react to the quality param
// tried 1-11

// Benchmark:
// GZIP best compression
// Cumulative average saving: 82.02%
// Cumulative average compression time: 10.362933684931509 ms (73 samples)
//
// BROTLI_COMPRESS quality: 9. Higher quality values take 10x longer to compress (???)
// Cumulative average saving: 86.39%
// Cumulative average compression time: 23.408338123287667 ms (73 samples)

const options: zlib.ZlibOptions = {
	// turned off as debugging shows an almost worse
	// cumulative average saving with this enabled
	level: zlib.constants.Z_BEST_COMPRESSION,
};

const { compress } = brotliCompress({
	compressOptions: {
		chunkSize: 64 * 1024,
		// these get passed down when compressing in place
		// check below
		// params: {
		// 	[zlib.constants.BROTLI_PARAM_QUALITY]: 8,
		// 	[zlib.constants.BROTLI_PARAM_MODE]: zlib.constants.BROTLI_MODE_TEXT,
		// },
	} as any,
	decompressOptions: {
		chunkSize: 1024,
		params: {
			[zlib.constants.BROTLI_PARAM_MODE]: zlib.constants.BROTLI_MODE_TEXT,
		},
	},
});

const debug = false;
let totalOriginalSize = 0;
let totalCompressedSize = 0;
let totalTimeTaken = 0;
let totalRequests = 0;
let samples = 0;
type modeType = 'GZIP' | 'BROTLI_ZLIB' | 'BROTLI_COMPRESS';
const mode = 'BROTLI_COMPRESS' as modeType;

export async function post<T>(
	url: string,
	requestBody: T,
	additionalHeaders: Record<string, string> = {}
): Promise<any> {
	try {
		// Convert the request body to a JSON string
		const jsonData: string = JSON.stringify(requestBody);
		const originalSize = jsonData.length;
		if (debug) {
			console.log(`Size before compression: ${originalSize} bytes`);
			totalOriginalSize += originalSize;
		}

		// Compress the data using gzip
		const startTime = process.hrtime();
		samples++;
		const compressedData = await new Promise<Buffer>(
			async (resolve, reject) => {
				let compressFunc = zlib.gzip;
				switch (mode) {
					case 'GZIP':
					case 'BROTLI_ZLIB':
						if (mode == 'BROTLI_ZLIB') {
							compressFunc = zlib.brotliCompress;
						}
						compressFunc(jsonData, options, (err, buffer) => {
							if (err) {
								reject(err);
							} else {
								resolve(buffer);
							}
							if (debug) {
								const endTime = process.hrtime(startTime);
								const elapsedTimeInMilliseconds =
									endTime[0] * 1000 + endTime[1] / 1000000;
								totalTimeTaken += elapsedTimeInMilliseconds;
								console.log(
									`Compression time: ${elapsedTimeInMilliseconds.toFixed(2)} ms`
								);
							}
						});
						break;
					case 'BROTLI_COMPRESS':
						try {
							let res = await compress(requestBody as any, {
								params: {
									[zlib.constants.BROTLI_PARAM_MODE]:
										zlib.constants.BROTLI_MODE_TEXT,
									[zlib.constants.BROTLI_PARAM_QUALITY]: 9,
								},
							});
							resolve(Buffer.from(res));
						} catch (e) {
							reject(e);
						}
						if (debug) {
							const endTime = process.hrtime(startTime);
							const elapsedTimeInMilliseconds =
								endTime[0] * 1000 + endTime[1] / 1000000;
							totalTimeTaken += elapsedTimeInMilliseconds;
							console.log(
								`Compression time: ${elapsedTimeInMilliseconds.toFixed(2)} ms`
							);
						}
						break;
				}
			}
		);

		const compressedSize = compressedData.length;
		if (debug) {
			totalCompressedSize += compressedSize;
			const currentSaving =
				((originalSize - compressedSize) / originalSize) * 100;
			const cumulativeAverageSaving =
				((totalOriginalSize - totalCompressedSize) / totalOriginalSize) * 100;

			console.log(
				`Size after compression: ${compressedSize} bytes, current saving: ${currentSaving.toFixed(
					2
				)}%`
			);
			console.log(
				`Cumulative average saving: ${cumulativeAverageSaving.toFixed(2)}%`
			);
			console.log(
				`Cumulative average compression time: ${
					totalTimeTaken / ++totalRequests
				} ms (${samples} samples)`
			);
		}

		let compressHeader = 'gzip';
		switch (mode) {
			case 'BROTLI_COMPRESS':
			case 'BROTLI_ZLIB':
				compressHeader = 'br';
				break;
		}
		// Combine the additional headers with the default headers
		const headers = {
			'Content-Encoding': compressHeader,
			'Content-Type': 'application/json',
			...additionalHeaders,
		};

		// Perform the POST request
		const response = await fetch(url, {
			method: 'POST',
			headers: headers,
			body: compressedData,
		});

		// Attempt to parse the response as JSON
		let clonedResponse: Response | undefined;
		try {
			// Clone the response before parsing it
			clonedResponse = response.clone();
			const jsonData = await response.json(); // Attempt to parse the original response
			return jsonData; // Return the parsed JSON data if successful
		} catch (e) {
			if (clonedResponse == undefined) {
				throw new Error(`Response was not cloned successfully`);
			}
			// If parsing fails, use the cloned response to read the raw text
			const rawResponse = await clonedResponse.text();
			throw new Error(
				`Invalid JSON: ${rawResponse}, error: ${(e as any)?.message}`
			);
		}
	} catch (error) {
		throw new Error(
			typeof error === 'string' ? error : (error as any)?.message
		);
	}
}

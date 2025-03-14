import { tryClick, testRun, tryType } from './utils';
import { Page } from '@playwright/test';

testRun(
	'Wait for new thread to update and stop changing',
	async (page: Page): Promise<void> => {
		tryClick(page, '.sng-test-ai-button', 5000);
		tryClick(page, '.sng-test-download-button', 300000);
		tryClick(page, '.sng-test-runtime-button', 300000);
		await tryClick(page, '.sng-test-go-button', 300000);

		await tryClick(page, '.sng-test-open-new-thread', 10000);

		await tryType(
			page,
			'.sng-test-chat-textarea textarea',
			'Tell me three examples of fruits. Names are enough. Keep it short.',
			300000
		);

		await tryClick(page, '.sng-test-chat-send-button', 300000);
		await page.waitForSelector('.sng-test-streaming-is-in-progress', {
			state: 'visible',
			timeout: 30000,
		});

		console.log('Streaming now.');

		await page.waitForSelector('.sng-test-streaming-is-in-progress', {
			state: 'detached',
			timeout: 30000,
		});
		console.log('Streaming finished');

		await page.waitForFunction(
			() => {
				const firstThreadItem = document.querySelector(
					'.sng-test-thread-item-index-0'
				);
				return (
					firstThreadItem &&
					(firstThreadItem?.textContent?.includes('fruit') ||
						firstThreadItem?.textContent?.includes('Fruit'))
				);
			},
			{ timeout: 300000 }
		);
	},
	{ playWrightKey: 'wait-for-thread-update' }
);

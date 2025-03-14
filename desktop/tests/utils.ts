import { _electron as electron } from 'playwright';
import { Page } from '@playwright/test';
import { test, expect } from '@playwright/test';
import { findLatestBuild, parseElectronApp } from 'electron-playwright-helpers';
import { join } from 'path';

// this seems to get rid of an ECONNNERR when running tests
const delayTime = 500;

export async function tryInteract(
	page: Page,
	selector: string,
	action,
	timeout = 30000
) {
	const startTime = Date.now();
	let isActionSuccessful = false;

	async function checkAndPerformAction() {
		try {
			const element = await page.locator(selector);
			const isVisible = await element.isVisible();
			if (isVisible) {
				console.log('Performing action on', selector);
				await action(element); // Perform the passed action
				return true;
			}
		} catch (error) {
			// Interpret specific errors as the application exiting
			if (
				error.message.includes('Application exited') ||
				error.message.includes('page is closed') ||
				error.message.includes('has been closed')
			) {
				console.log('Application has exited or page is closed.');
				return true; // Exit the polling loop
			}
			console.error(
				`An error occurred while looking for "${selector}":`,
				error.message
			);
		}
		return false;
	}

	// Polling mechanism
	while (Date.now() - startTime < timeout) {
		if (await checkAndPerformAction()) {
			return; // Action was performed successfully or application exited, exit function
		}
		// Short delay for polling
		await new Promise((resolve) => setTimeout(resolve, 100)); // 100ms polling interval
	}

	if (!isActionSuccessful) {
		console.log(
			selector,
			'not found or not interactable within timeout period'
		);
	}
}

// Convenience function for clicking
export async function tryClick(page: Page, selector: string, timeout = 30000) {
	await tryInteract(
		page,
		selector,
		async (element) => {
			console.log('Clicking', selector);
			await element.click();
		},
		timeout
	);
}

// Convenience function for typing
export async function tryType(
	page: Page,
	selector: string,
	text,
	timeout = 30000
) {
	await tryInteract(
		page,
		selector,
		async (element) => {
			console.log('Typing into', selector);
			await element.fill(text);
		},
		timeout
	);
}

export interface TestOptions {
	playWrightKey?: string;
	timeout?: number;
}

export async function testRun(
	name: string,
	cb: (page: Page) => Promise<void>,
	options?: TestOptions
) {
	test(name, async ({}, testInfo) => {
		await new Promise((resolve) => setTimeout(resolve, delayTime));

		testInfo.setTimeout(options?.timeout || 300000);

		const latestBuild = findLatestBuild();
		const appInfo = parseElectronApp(latestBuild);

		let app = await electron.launch({
			args: [appInfo.main], // main file from package.json
			executablePath: appInfo.executable, // path to the Electron executable
			timeout: 10000,
		});

		// Capture backend (main process) logs
		app.on('console', (message) => {
			console.log(
				`${new Date().toLocaleString()} [Main Process] ${message.text()}`
			);
		});

		const window = await app.firstWindow();
		if (!window) {
			throw 'Window is undefined';
		}

		// Capture frontend (renderer process) logs
		window.on('console', (message) => {
			console.log(
				`${new Date().toLocaleString()} [Renderer Process] ${message.text()}`
			);
		});

		const context = window.context();
		await context.tracing.start({
			screenshots: true,
			snapshots: true,
			sources: true,
		});

		try {
			await cb(window);
		} finally {
			// This ensures the trace is saved even if the test fails.
			await context.tracing.stop({
				path: join(
					'playwright-traces',
					`${options?.playWrightKey || name}-trace.zip`
				),
			});
			await app.close();
		}
	});
}

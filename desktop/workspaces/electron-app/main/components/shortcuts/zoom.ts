/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { BrowserWindow, GlobalShortcut } from 'electron';

export function registerZoomShortcuts(
	window: BrowserWindow,
	globalShortcut: GlobalShortcut
) {
	if (!window) {
		console.error('Cannot register zoom shortcuts: window is null');
		return;
	}

	const { webContents } = window;

	function zoomIn() {
		const currentZoom = webContents.getZoomFactor();
		webContents.setZoomFactor(currentZoom + 0.1);
	}

	function zoomOut() {
		const currentZoom = webContents.getZoomFactor();
		webContents.setZoomFactor(currentZoom - 0.1);
	}

	function resetZoom() {
		webContents.setZoomFactor(1.0);
	}

	// @todo is there a way to generalize this?
	globalShortcut.register('CommandOrControl+Plus', zoomIn); // Most standard layouts
	globalShortcut.register('CommandOrControl+Shift+=', zoomIn); // Alternative: + on some keyboards
	globalShortcut.register('Shift+CommandOrControl+=', zoomIn); // Some English layouts
	globalShortcut.register('CommandOrControl+Shift+3', zoomIn); // Layouts where + is Shift+3
	globalShortcut.register('CommandOrControl+Shift+1', zoomIn); // German and Swiss layouts
	// globalShortcut.register('CommandOrControl+Shift+´', zoomIn); // Swedish, Finnish, Norwegian layouts
	// globalShortcut.register('CommandOrControl+Shift+`', zoomIn); // Alternative for some layouts

	globalShortcut.register('CommandOrControl+-', zoomOut);

	globalShortcut.register('CommandOrControl+0', resetZoom);
}

import { app, BrowserWindow } from 'electron';
import * as path from 'path';

// IS THIS FILE EVEN USED AT ALL?

function createWindow() {
	const mainWindow = new BrowserWindow({
		height: 600,
		webPreferences: {
			nodeIntegration: true,
			preload: path.join(__dirname, 'preload.js'),
		},
		width: 800,
		autoHideMenuBar: true,
	});

	mainWindow.loadFile(path.join(__dirname, '../index.html'));

	mainWindow.webContents.openDevTools();
}

app.whenReady().then(() => {
	createWindow();

	app.on('activate', function () {
		if (BrowserWindow.getAllWindows().length === 0) createWindow();
	});
});

app.on('window-all-closed', () => {
	if (process.platform !== 'darwin') {
		app.quit();
	}
});

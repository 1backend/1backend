import { contextBridge, ipcRenderer, IpcRendererEvent } from 'electron';
import { WindowApiConst } from 'shared-lib';

contextBridge.exposeInMainWorld('api', {
	node: () => process.versions.node,
	chrome: () => process.versions.chrome,
	electron: () => process.versions.electron,
	send: <In>(channel: string, input: In) => {
		if (WindowApiConst.SENDING_SAFE_CHANNELS.includes(channel)) {
			ipcRenderer.send(channel, input);
		} else {
			console.warn(`Channel [${channel}] is not in the safe channels list`);
		}
	},
	receive: <Out>(channel: string, callback: (output: Out) => void) => {
		ipcRenderer.on(channel, (_event: IpcRendererEvent, ...parameters: any[]) =>
			callback(parameters[0])
		);
	},
});

console.log('The preload script has been injected successfully.');

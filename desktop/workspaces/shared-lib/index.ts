import { WindowApi } from './apis/window-api';
export * from './apis/window-api';
export * from './apis/window-api-consts';
export * from './models/config/app-config';

declare global {
	interface Window {
		api: WindowApi;
	}
}

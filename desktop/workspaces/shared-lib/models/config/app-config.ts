export interface AppConfig {
	configId: string;
	mainLogFile: string;
	mainLogLevel: string;
	isIconAvailable: boolean;
	isNodeIntegration: boolean;
	isContextIsolation: boolean;
	isSandbox: boolean;
	isEnableRemoteModule: boolean;
	isOpenDevTools: boolean;
}

import * as fs from 'fs-extra';
import _ from 'lodash';
import * as path from 'node:path';
import { AppConfig } from 'shared-lib';
import { App } from './components/app';

declare const global: Global;

declare global {
	interface Global {
		appConfig: AppConfig;
	}
}

const currentEnvironment = process.env.X_NODE_ENV || process.env.NODE_ENV;
const appConfigs = fs.readJsonSync(path.join(__dirname, 'config.json'));
const defaultConfig = appConfigs.development;
const currentConfig = appConfigs[currentEnvironment || 'development'];
global.appConfig =
	currentEnvironment === 'development'
		? defaultConfig
		: _.merge(defaultConfig, currentConfig);

App.launch();

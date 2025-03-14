module.exports = function (config) {
	config.set({
		basePath: '',
		frameworks: [
			'jasmine',
			'@angular-devkit/build-angular',
		],
		plugins: [
			require('karma-jasmine'),
			require('karma-chrome-launcher'),
			require('karma-jasmine-html-reporter'),
			require('karma-coverage'),
			require('@angular-devkit/build-angular/plugins/karma'),
		],
		browsers: ['ChromeHeadlessNoSandbox'],
		customLaunchers: {
			ChromeHeadlessNoSandbox: {
				base: 'ChromeHeadless',
				flags: ['--no-sandbox'],
			},
		},
		client: {
			jasmine: {},
			clearContext: false,
		},
		jasmineHtmlReporter: {
			suppressAll: true,
		},
		coverageReporter: {
			dir: require('path').join(__dirname, './coverage/angular-app'),
			subdir: '.',
			reporters: [{ type: 'html' }, { type: 'text-summary' }],
		},
		reporters: ['progress', 'kjhtml'],
		port: 9876,
		colors: true,
		logLevel: config.LOG_INFO,
		autoWatch: true,
		browsers: ['Chrome'],
		singleRun: false,
		restartOnFileChange: true,
	});
};

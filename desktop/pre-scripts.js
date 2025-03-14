const https = require('https');
const fs = require('fs');
const path = require('path');
const os = require('os');

if (os.platform() === 'win32') {
	// this used to be a place where dind was copied into the build.
	// not anymore due to slow build time
}

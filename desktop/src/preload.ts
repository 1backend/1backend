// Node.js APIs are available in the preload process
window.addEventListener('DOMContentLoaded', () => {
	const replaceText = (selector: string, text: string) => {
		const element = document.getElementById(selector);
		if (element) {
			element.innerText = text;
		}
	};

	for (const type of ['chrome', 'node', 'electron']) {
		replaceText(
			`${type}-version`,
			process.versions[type as keyof NodeJS.ProcessVersions]
		);
	}
});

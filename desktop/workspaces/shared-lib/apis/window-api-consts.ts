export class WindowApiConst {
	public static readonly FRONTEND_READY_REQUEST = 'frontendReadyRequest';

	public static readonly SELECT_FOLDER_REQUEST = 'selectFolderRequest';
	public static readonly ON_FOLDER_SELECT = 'onFolderSelect';

	static readonly COPY_TO_CLIPBOARD_REQUEST = 'copyToClipboardRequest';

	static readonly ON_SYSTEM_LANGUAGE = 'onSystemLanguage';

	static readonly INSTALL_RUNTIME_REQUEST = 'installRuntimeRequest';

	static readonly ON_RUNTIME_INSTALL_LOG = 'onRuntimeInstallLog';

	static readonly LOG_REQUEST = 'logRequest';

	static readonly DISABLE_LOGGING_REQUEST = 'disableLoggingRequest';
	static readonly ENABLE_LOGGING_REQUEST = 'enableLoggingRequest';

	// angular -> electron
	public static readonly SENDING_SAFE_CHANNELS = [
		WindowApiConst.FRONTEND_READY_REQUEST,
		WindowApiConst.SELECT_FOLDER_REQUEST,
		WindowApiConst.COPY_TO_CLIPBOARD_REQUEST,
		WindowApiConst.INSTALL_RUNTIME_REQUEST,
		WindowApiConst.LOG_REQUEST,
		WindowApiConst.DISABLE_LOGGING_REQUEST,
		WindowApiConst.ENABLE_LOGGING_REQUEST,
	];

	// electron -> angular
	public static readonly RECEIVING_SAFE_CHANNELS = [
		WindowApiConst.ON_FOLDER_SELECT,
		WindowApiConst.ON_SYSTEM_LANGUAGE,
		WindowApiConst.ON_RUNTIME_INSTALL_LOG,
	];
}

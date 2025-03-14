export interface WindowApi {
	receive<Out>(channel: string, callback: (output: Out) => void): void;

	send<In>(channel: string, input: In): void;
}

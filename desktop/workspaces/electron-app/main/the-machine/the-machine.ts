import { EventService } from '../services/event-service';
import { OSManager } from './os_manager';

export class Machine {
	public eventService: EventService;

	constructor(private assetFolder: string) {
		this.eventService = new EventService();

		new OSManager(this.assetFolder, this.eventService);
	}
}

/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Injectable } from '@angular/core';
import { DataService } from './generic.service';
import { ServerService } from './server.service';
import { UserService } from './user.service';
import { first } from 'rxjs';
import { DataSvcObject } from '@openorch/client';

const CHARACTERS_TABLE_NAME = 'characters';
const SELECTED_CHARACTERS_TABLE_NAME = 'selected-characters';

@Injectable({
	providedIn: 'root',
})
export class CharacterService {
	public selectedCharacter!: Character;

	constructor(
		private server: ServerService,
		private dataService: DataService,
		private userService: UserService
	) {
		this.userService.user$.pipe(first()).subscribe(() => {
			this.init();
		});
	}

	async init() {
		const selectedCharacter = await this.getSelectedCharacter();
		if (selectedCharacter) {
			this.selectedCharacter = selectedCharacter;
		}
	}

	async loadCharacters(): Promise<Character[]> {
		const response = await this.dataService.find(CHARACTERS_TABLE_NAME, [
			// all(),
		]);
		return response?.objects as Character[];
	}

	async upsertCharacter(character: Character) {
		const exists = await this.getCharacter(character.id!);
		return exists
			? await this.updateCharacter(character)
			: await this.createNewCharacter(character);
	}

	async createNewCharacter(character: Character) {
		const id = this.server.id('char');
		const now = new Date().toISOString();
		character.id = id;
		await this.dataService.create(CHARACTERS_TABLE_NAME, {
			...character,
			id,
			createdAt: now,
			updatedAt: now,
		});
	}

	async deleteCharacter(character: Character) {
		if (character.id === this.selectedCharacter?.id) {
			this.selectedCharacter = {} as any;
			await this.deleteCharacterSelection(character.id!);
		}
		await this.dataService.delete(CHARACTERS_TABLE_NAME, [
			// idCondition(character.id!),
		]);
	}

	async updateCharacter(character: Character) {
		const now = new Date().toISOString();
		character.updatedAt = now;
		this.dataService.update(
			CHARACTERS_TABLE_NAME,
			[
				//idCondition(character.id!)
			],
			{
				...character,
			}
		);
	}

	async getCharacter(characterId: string): Promise<Character | undefined> {
		console.log(characterId);
		try {
			const response = await this.dataService.find(CHARACTERS_TABLE_NAME, [
				// idCondition(characterId),
			]);
			return response?.objects?.[0] as any as Character;
		} catch {
			return;
		}
	}

	async getSelectedCharacter(): Promise<Character | undefined> {
		const characterSelection = await this.getCharacterSelection();
		const selectedCharacterId = characterSelection?.data
			?.selectedCharacterId as any as string;
		if (!selectedCharacterId) {
			return;
		}
		const character = await this.getCharacter(selectedCharacterId);
		if (character) {
			this.selectedCharacter = character;
		}
		return character;
	}

	async selectCharacter(character: Character): Promise<void> {
		this.selectedCharacter = character;
		this.upsertCharacterSelection(character.id!);
	}

	async upsertCharacterSelection(selectedCharacterId: string): Promise<void> {
		const now = new Date().toISOString();
		let characterSelection = await this.getCharacterSelection();
		if (!characterSelection) {
			characterSelection = initCharacterSelection();
			characterSelection.id = this.server.id('char');
			characterSelection.createdAt = now;
		}
		characterSelection.updatedAt = now;
		characterSelection.data.selectedCharacterId = selectedCharacterId;
		this.dataService.upsert(SELECTED_CHARACTERS_TABLE_NAME, {
			...characterSelection,
		});
	}

	async getCharacterSelection(): Promise<SelectedCharacter | null> {
		const userId = await this.userService.getUserId();
		console.log(userId);
		const response = await this.dataService.find(
			SELECTED_CHARACTERS_TABLE_NAME,
			[
				// userIdCondition(userId)
			]
		);
		return response?.objects?.[0] as SelectedCharacter;
	}

	/**
	 * This will remove only the selection from SELECTED_CHARACTERS_TABLE_NAME
	 */
	async deleteCharacterSelection(characterId: string) {
		console.log(characterId);
		this.dataService.delete(SELECTED_CHARACTERS_TABLE_NAME, [
			//idCondition(characterId),
		]);
	}
}

export interface Character extends DataSvcObject {
	data: {
		name: string;
		behaviour: string;
		private: boolean;
	};
}

export interface SelectedCharacter extends DataSvcObject {
	data: {
		selectedCharacterId: string;
	};
}

export function initCharacter(): Character {
	return {
		data: {
			name: '',
			behaviour: '',
			private: false,
		},
		id: '',
		table: '',
		createdAt: '',
		updatedAt: '',
	};
}

export function initCharacterSelection(): SelectedCharacter {
	return {
		data: {
			selectedCharacterId: '',
		},
		id: '',
		table: '',
		createdAt: '',
		updatedAt: '',
	};
}

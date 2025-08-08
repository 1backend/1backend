/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import {
	Component,
	ViewEncapsulation,
	ViewChild,
	ChangeDetectionStrategy,
	ChangeDetectorRef,
} from '@angular/core';
import { NgFor, NgIf } from '@angular/common';
import {
	IonModal,
	IonHeader,
	IonToolbar,
	IonTitle,
	IonButtons,
	IonButton,
	IonIcon,
	IonContent,
	IonSegment,
	IonLabel,
	IonSegmentButton,
	IonItem,
	IonTextarea,
} from '@ionic/angular/standalone';
import { FormsModule } from '@angular/forms';
import {
	CharacterService,
	Character,
	initCharacter,
} from '../../services/character.service';
import { addIcons } from 'ionicons';
import { createOutline, playCircleOutline, closeOutline } from 'ionicons/icons';

@Component({
	selector: 'app-ai-character',
	templateUrl: './character.component.html',
	styleUrl: './character.component.scss',
	imports: [
		IonModal,
		IonHeader,
		IonToolbar,
		IonTitle,
		IonButtons,
		IonButton,
		IonIcon,
		IonContent,
		IonSegment,
		IonLabel,
		NgFor,
		NgIf,
		FormsModule,
		IonSegmentButton,
		IonItem,
		IonTextarea,
	],
	encapsulation: ViewEncapsulation.None,
	changeDetection: ChangeDetectionStrategy.OnPush,
})
export class CharacterComponent {
	@ViewChild(IonModal) modal!: IonModal;
	public isOpen: boolean = false;
	public segment = 'select';
	public editingCharacter: Character = initCharacter();
	public characters: Character[] = [];

	constructor(
		private characterService: CharacterService,
		private cd: ChangeDetectorRef
	) {
		addIcons({
			'create-outline': createOutline,
			'close-outline': closeOutline,
			'play-circle-outline': playCircleOutline,
		});
	}

	async ngOnInit() {
		await this.loadCharacters();
	}

	async loadCharacters() {
		this.characters = await this.characterService.loadCharacters();
	}

	public clearEditingCharacter(): void {
		this.editingCharacter = initCharacter();
	}

	selectCharacter(character: Character) {
		this.characterService.selectCharacter(character);
	}

	async upsertCharacter(character: Character) {
		await this.characterService.upsertCharacter(character);
		this.clearEditingCharacter();
		await this.loadCharacters();
	}

	async deleteCharacter(character: Character) {
		await this.characterService.deleteCharacter(character);
		await this.loadCharacters();
	}

	async selectCharacterForEdit(character: Character) {
		this.editingCharacter = character;
		this.segment = 'create';
	}

	getModeText(): string {
		return this.editingCharacter?.id ? 'Edit' : 'Create';
	}

	show(): void {
		this.isOpen = true;
		this.cd.markForCheck();
	}

	close(): void {
		this.isOpen = false;
		this.cd.markForCheck();
	}
}

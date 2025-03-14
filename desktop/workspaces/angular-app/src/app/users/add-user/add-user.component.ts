/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Component, OnInit } from '@angular/core';
import {
	FormBuilder,
	FormGroup,
	Validators,
	FormsModule,
	ReactiveFormsModule,
} from '@angular/forms';
import { UserService } from '../../services/user.service';
import { UserSvcUser as User, UserSvcRole as Role } from '@openorch/client';
import { first } from 'rxjs';
import {
	ToastController,
	IonCardContent,
	IonItem,
	IonSelectOption,
	IonSelect,
	IonRow,
	IonCol,
	IonButton,
} from '@ionic/angular/standalone';
import { TranslatePipe } from '../../translate.pipe';
import { TranslateModule } from '@ngx-translate/core';
import { NgFor } from '@angular/common';
import { CenteredComponent } from '../../components/centered/centered.component';
import { ChangeDetectorRef, ChangeDetectionStrategy } from '@angular/core';
import { PageComponent } from '../../components/page/page.component';
import { IconMenuComponent } from '../../components/icon-menu/icon-menu.component';

@Component({
	selector: 'app-add-user',
	templateUrl: './add-user.component.html',
	styleUrls: ['./add-user.component.scss'],
	imports: [
		IonRow,
		IonCol,
		IonButton,
		IonCardContent,
		IonItem,
		IonSelectOption,
		IonSelect,
		PageComponent,
		IconMenuComponent,
		CenteredComponent,
		FormsModule,
		ReactiveFormsModule,
		NgFor,
		TranslateModule,
		TranslatePipe,
	],
	changeDetection: ChangeDetectionStrategy.OnPush,
})
export class AddUserComponent implements OnInit {
	addUserForm!: FormGroup;
	roles: Role[] = [];

	constructor(
		private fb: FormBuilder,
		private userService: UserService,
		private toast: ToastController,
		private cd: ChangeDetectorRef
	) {
		this.userService.user$.pipe(first()).subscribe(() => {
			this.loggedInInit();
		});
	}

	ngOnInit() {
		this.addUserForm = this.fb.group({
			email: ['', [Validators.required, Validators.email]],
			name: ['', Validators.required],
			password: ['', Validators.required],
			passwordConfirmation: ['', Validators.required],
			roles: [[], Validators.required],
		});
	}

	async loggedInInit() {
		const rsp = await this.userService.getRoles();
		this.roles = rsp.roles!;
		this.cd.markForCheck();
	}

	async createUser() {
		if (this.addUserForm.invalid) {
			return;
		}

		const { slug, name, password, passwordConfirmation, roles } =
			this.addUserForm.value;

		if (password !== passwordConfirmation) {
			const toast = await this.toast.create({
				message: 'Passwords do not match',
				duration: 5000,
				color: 'danger',
				cssClass: 'white-text',
				position: 'middle',
			});
			toast.present();
			return;
		}

		const user: User = { slug, name };

		try {
			await this.userService.createUser(user, password, roles);
			const toast = await this.toast.create({
				message: `User ${slug} saved`,
				duration: 5000,
				color: 'secondary',
				position: 'middle',
			});
			toast.present();

			this.addUserForm.reset();
		} catch (error) {
			let errorMessage = 'An unexpected error occurred';
			try {
				errorMessage = (JSON.parse(error as string) as { error: string })
					?.error;
			} catch {}

			const toast = await this.toast.create({
				color: 'danger',
				cssClass: 'white-text',
				message: errorMessage,
				duration: 5000,
				position: 'middle',
			});
			toast.present();
		}
	}
}

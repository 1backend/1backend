import { Component } from '@angular/core';
import { UserService } from '../../services/user.service';
import { first } from 'rxjs';
import { NgFor, NgIf } from '@angular/common';
import { FormsModule } from '@angular/forms';
import {
	IonList,
	IonListHeader,
	IonLabel,
	IonItem,
	IonButton,
	IonCheckbox,
} from '@ionic/angular/standalone';
import { ChangeDetectorRef, ChangeDetectionStrategy } from '@angular/core';
import { PageComponent } from '../../components/page/page.component';
import { IconMenuComponent } from '../../components/icon-menu/icon-menu.component';

@Component({
	selector: 'app-roles',
	templateUrl: './roles.component.html',
	styleUrls: ['./roles.component.css'],
	imports: [
		IonList,
		IonListHeader,
		IonLabel,
		IonItem,
		IonButton,
		IonCheckbox,
		PageComponent,
		IconMenuComponent,
		FormsModule,
		NgFor,
		NgIf,
	],
	changeDetection: ChangeDetectionStrategy.OnPush,
})
export class RolesComponent {
	roles: string[] = [];
	permissions: string[] = [];
	selectedRole: string | undefined;
	selectedRolePermissions: Set<string> = new Set<string>();

	roleSearchQuery: string = '';
	permissionSearchQuery: string = '';

	constructor(
		private userService: UserService,
		private cd: ChangeDetectorRef
	) {
		this.userService.user$.pipe(first()).subscribe(() => {
			this.loggedInInit();
		});
	}

	async loggedInInit() {
		const rsp = await this.userService.getRoles();
		this.roles = await rsp.roles!;

		this.cd.markForCheck();
	}

	selectRole(role: string) {
		this.selectedRole = role;
		this.loadRolePermissions(role);
	}

	async loadRolePermissions(role: string) {
		this.selectedRolePermissions.clear();
		const rsp = await this.userService.getPermissions(role!);
		this.permissions = rsp.permissions!;
		// @todo fix this
		// if (role.permissionIds) {
		// 	for (const id of role.permissionIds) {
		// 		this.selectedRolePermissions.add(id);
		// 	}
		// }
		this.cd.markForCheck();
	}

	togglePermission(permissionId: string) {
		if (this.selectedRolePermissions.has(permissionId)) {
			this.selectedRolePermissions.delete(permissionId);
		} else {
			this.selectedRolePermissions.add(permissionId);
		}
	}

	async savePermissions() {
		if (this.selectedRole) {
			// const permissionIds = [...this.selectedRolePermissions];
			// @todo
		}
	}

	filteredRoles() {
		return this.roles.filter((role) =>
			role?.toLowerCase().includes(this.roleSearchQuery.toLowerCase())
		);
	}

	filteredPermissions() {
		return this.permissions.filter((permission) =>
			permission
				.toLowerCase()
				.includes(this.permissionSearchQuery.toLowerCase())
		);
	}
}

/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Routes } from '@angular/router';
import { StartupComponent } from './startup/startup.component';
import { ChatComponent } from './chat/chat.component';
import { ModelExplorerComponent } from './model-explorer/model-explorer.component';
import { HomeComponent } from './home/home.component';
import { LoginComponent } from './login/login.component';
import { UsersComponent } from './users/users.component';
import { AddUserComponent } from './users/add-user/add-user.component';
import { RolesComponent } from './users/roles/roles.component';
import { LogoutComponent } from './users/logout/logout.component';
import { PromptsComponent } from './prompts/prompts.component';
import { NodesComponent } from './nodes/nodes.component';

export const routes: Routes = [
	{
		path: '',
		component: HomeComponent,
	},
	{
		path: 'startup',
		component: StartupComponent,
	},
	{
		path: 'chat',
		component: ChatComponent,
	},
	{
		path: 'model-explorer',
		component: ModelExplorerComponent,
	},
	{
		path: 'users',
		component: UsersComponent,
	},
	{
		path: 'add-user',
		component: AddUserComponent,
	},
	{
		path: 'roles',
		component: RolesComponent,
	},
	{
		path: 'login',
		component: LoginComponent,
	},
	{
		path: 'logout',
		component: LogoutComponent,
	},
	{
		path: 'prompts',
		component: PromptsComponent,
	},
	{
		path: 'nodes',
		component: NodesComponent,
	},
];

import {
  Component,
  OnInit,
  ViewEncapsulation,
  Input,
  Output,
  EventEmitter
} from '@angular/core';
import * as types from '../../../types';
import { ProjectService } from '../../../project.service';
import { FormsModule } from '@angular/forms';
import { UserService } from '../../../user.service';
import { Router } from '@angular/router';
import { NotificationsService } from 'angular2-notifications';

@Component({
  selector: 'app-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.css'],
  encapsulation: ViewEncapsulation.None
})
export class SettingsComponent implements OnInit {
  @Input() project: types.Project;
  @Output() onProjectSaved = new EventEmitter<void>();
  show = false;
  callerId = '';
  color = 'accent';
  privateChecked = false;
  privateDisabled = false;
  isPremiumMember = false;
  newProjectName;
  nameChanged = false;
  deleteProjectName: string;
  newVersion;
  newNamespace;
  newDescription;

  constructor(
    private ps: ProjectService,
    private us: UserService,
    private router: Router,
    private notif: NotificationsService
  ) {}

  ngOnInit() {
    this.isPremiumMember = this.us.user.Premium;
    this.privateChecked = !this.project.OpenSource;
    this.newProjectName = this.project.Name;
    this.newVersion = this.project.Version;
    this.newNamespace = this.project.Namespace;
    this.newDescription = this.project.Description;
  }

  showCallerId() {
    if (this.callerId !== '') {
      this.show = true;
      return;
    }
    this.ps.getCallerId(this.project).then(callerIdResponse => {
      this.callerId = callerIdResponse.CallerId;
      this.show = true;
    });
  }

  hideCallerId() {
    this.show = false;
  }

  saveSettings() {
    this.project.Version = this.newVersion;
    this.project.Description = this.newDescription;
    this.project.Namespace = this.newNamespace;
    this.project.OpenSource = !this.privateChecked;
    this.ps.update(this.project).then(() => {
      this.onProjectSaved.emit();
    });
  }

  deleteProject() {
    if (this.project.Name !== this.deleteProjectName) {
      this.notif.error('Incorrect project name');
      return;
    }
    this.ps
      .delete(this.project)
      .then(() => {
        this.router.navigate(['/' + this.us.user.Nick]);
      })
      .catch(err => {
        this.notif.error(err);
      });
  }

  changeName() {
    if (this.project.Name !== this.newProjectName) {
      this.project.Name = this.newProjectName;
      this.nameChanged = true;
    }
    this.ps.update(this.project).then(() => {
      if (this.nameChanged) {
        this.router.navigate([
          '/' + this.project.Author + '/' + this.project.Name
        ]);
        this.onProjectSaved.emit();
      } else {
        this.onProjectSaved.emit();
      }
    });
  }
}

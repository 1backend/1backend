import { Component, OnInit, ViewEncapsulation, Input, Output, EventEmitter } from '@angular/core';
import * as types from '../../../types';
import { ProjectService } from '../../../project.service';
import { FormsModule } from '@angular/forms';
import { UserService } from '../../../user.service';
import { Router } from '@angular/router';

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

  constructor(private ps: ProjectService, private us: UserService, private router: Router) { }

  ngOnInit() {
    this.isPremiumMember = this.us.user.Premium;
    this.privateChecked = !this.project.Public;
    this.newProjectName = this.project.Name;
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
    if(this.project.Name != this.newProjectName) {
      this.project.Name = this.newProjectName;
      this.nameChanged = true;
    }
    this.project.Public = !this.privateChecked;
    this.ps.update(this.project).then(() => {
      if ( this.nameChanged ) {
        this.router.navigate(['/' + this.project.Author + '/' + this.project.Name]);
        this.onProjectSaved.emit();
      } else {
        this.onProjectSaved.emit();
      }
    })
  }
}

import { Component, OnInit, ViewEncapsulation, Input, Output, EventEmitter } from '@angular/core';
import * as types from '../../../types';
import { ProjectService } from '../../../project.service';
import { FormsModule } from '@angular/forms';
import { UserService } from '../../../user.service';

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

  constructor(private ps: ProjectService, private us: UserService) { }

  ngOnInit() {
    this.isPremiumMember = this.us.user.Premium;
    this.privateChecked != this.project.OpenSource;
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
    this.project.OpenSource = !this.privateChecked;
    this.ps.update(this.project).then(() => {
      this.onProjectSaved.emit();
    })
  }

}

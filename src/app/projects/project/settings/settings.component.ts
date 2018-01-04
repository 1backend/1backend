import { Component, OnInit, ViewEncapsulation, Input } from '@angular/core';
import * as types from '../../../types';
import { ProjectService } from '../../../project.service';

@Component({
  selector: 'app-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.css'],
  encapsulation: ViewEncapsulation.None
})
export class SettingsComponent implements OnInit {
  @Input() project: types.Project;
  show = false;
  callerId = '';

  constructor(private ps: ProjectService) {}

  ngOnInit() {}

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
}

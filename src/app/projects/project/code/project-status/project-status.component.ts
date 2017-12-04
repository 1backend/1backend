import { Component, OnInit, Input, ViewEncapsulation } from '@angular/core';
import { ProjectService } from '../../../../project.service';

@Component({
  selector: 'app-project-status',
  templateUrl: './project-status.component.html',
  styleUrls: ['./project-status.component.css'],
  encapsulation: ViewEncapsulation.None
})
export class ProjectStatusComponent implements OnInit {
  @Input() author: string;
  @Input() projectName: string;
  status = false;
  loaded = false;

  constructor(private ps: ProjectService) {}

  getStatus() {
    this.ps
      .getStatus(this.author, this.projectName)
      .then(pingResponse => {
        this.loaded = true;
        this.status = pingResponse.pong;
      })
      .catch(() => {
        this.loaded = true;
      });
  }

  ngOnInit() {
    this.getStatus();
  }
}

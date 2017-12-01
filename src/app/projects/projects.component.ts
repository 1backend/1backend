import { Component, OnInit } from '@angular/core';
import * as types from '../types';
import { FilterPipe } from '../filter.pipe';
import { ProjectService } from '../project.service';

@Component({
  selector: 'app-projects',
  templateUrl: './projects.component.html',
  styleUrls: ['./projects.component.css']
})
export class ProjectsComponent implements OnInit {
  currentPage = 0;
  search = '';
  projects: types.Project[] = [];

  constructor(private ps: ProjectService) {
    this.refresh();
  }

  refresh() {
    this.ps
      .list()
      .then(projects => (this.projects = projects))
      .catch(err => console.log(err));
  }

  pageChanged($event: any) {
    this.currentPage = $event.pageIndex;
  }

  ngOnInit() {}
}

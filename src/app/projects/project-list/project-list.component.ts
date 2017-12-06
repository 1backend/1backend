import { Component, OnInit, Input } from '@angular/core';
import { CreateProjectDialogService } from '../create-project-dialog.service';
import { UserService } from '../../user.service';
import { Router } from '@angular/router';
import * as types from '../../types';
import { ActivatedRoute } from '@angular/router';
import { LoginDialogService } from '../../login/login-dialog.service';
import { ProjectService } from '../../project.service';
import { Title } from '@angular/platform-browser';

@Component({
  selector: 'app-project-list',
  templateUrl: './project-list.component.html',
  styleUrls: ['./project-list.component.css']
})
export class ProjectListComponent implements OnInit {
  @Input() projects: types.Project[] = [];

  search: string;
  currentPage = 0;
  author = '';
  isProjectsPage = false;

  constructor(
    private cp: CreateProjectDialogService,
    private lds: LoginDialogService,
    private router: Router,
    public us: UserService,
    private ps: ProjectService,
    private route: ActivatedRoute,
    private title: Title
  ) {
    this.author = this.route.snapshot.params['author'];
    this.isProjectsPage = this.router.isActive('projects', false);
  }

  ngOnInit() {
    if (this.router.url === '/projects') {
      this.title.setTitle('Projects');
    } else if (this.router.url === '/' + this.author) {
      this.title.setTitle(this.author);
    }
  }

  create() {
    const that = this;
    if (!this.us.loggedIn()) {
      this.lds.openDialog(true, () => {
        this.cp.openDialog(proj => {
          that.cp.closeDialog();
          that.router.navigate(['/' + proj.Author + '/' + proj.Name]);
        });
      });
    } else {
      this.cp.openDialog(proj => {
        that.cp.closeDialog();
        that.router.navigate(['/' + proj.Author + '/' + proj.Name]);
      });
    }
  }

  pageChanged($event: any) {
    this.currentPage = $event.pageIndex;
  }

  star(p: types.Project) {
    this.ps
      .star(p.Id)
      .then(() => {
        p.Stars++;
        p.Starrers = [];
      })
      .catch(error => {
        console.log(error.error);
      });
  }

  unStar(proj: types.Project) {
    this.ps
      .unstar(proj.Id)
      .then(() => {
        proj.Stars--;
        proj.Starrers = null;
      })
      .catch(error => {});
  }
}

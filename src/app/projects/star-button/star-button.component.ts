import {
  Component,
  OnInit,
  Input,
  Output,
  EventEmitter,
  ViewEncapsulation
} from '@angular/core';
import { ProjectService } from '../../project.service';
import { UserService } from '../../user.service';
import * as types from '../../types';

@Component({
  selector: 'app-star-button',
  templateUrl: './star-button.component.html',
  styleUrls: ['./star-button.component.css'],
  encapsulation: ViewEncapsulation.None
})
export class StarButtonComponent implements OnInit {
  @Input() showCount = false;
  @Input() project: types.Project;
  @Output() starEvent = new EventEmitter<boolean>();
  constructor(private ps: ProjectService, private us: UserService) {}

  ngOnInit() {}

  starredBySelf(): boolean {
    return (
      this.project.Starrers &&
      this.project.Starrers.filter(x => x.Id === this.us.user.Id).length > 0
    );
  }

  star() {
    this.ps.star(this.project.Id).then(() => {
      this.project.Stars++;
      this.starEvent.emit(true);
    });
  }

  unStar() {
    this.ps.unstar(this.project.Id).then(() => {
      this.project.Stars--;
      this.project.Starrers = this.project.Starrers.filter(
        x => x.Id !== this.us.user.Id
      );
      this.starEvent.emit(false);
    });
  }
}

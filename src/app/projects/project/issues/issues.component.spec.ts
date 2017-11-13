import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ListIssuesComponent } from './list-issues.component';

describe('ListIssuesComponent', () => {
  let component: ListIssuesComponent;
  let fixture: ComponentFixture<ListIssuesComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ListIssuesComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ListIssuesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});

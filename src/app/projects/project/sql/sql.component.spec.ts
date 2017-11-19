import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { SqlComponent } from './sql.component';

describe('SqlConsoleComponent', () => {
  let component: SqlComponent;
  let fixture: ComponentFixture<SqlComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ SqlComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(SqlComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});

import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { SqlConsoleComponent } from './sql-console.component';

describe('SqlConsoleComponent', () => {
  let component: SqlConsoleComponent;
  let fixture: ComponentFixture<SqlConsoleComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ SqlConsoleComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(SqlConsoleComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});

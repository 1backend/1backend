import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { StarButtonComponent } from './star-button.component';

describe('StarButtonComponent', () => {
  let component: StarButtonComponent;
  let fixture: ComponentFixture<StarButtonComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ StarButtonComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(StarButtonComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

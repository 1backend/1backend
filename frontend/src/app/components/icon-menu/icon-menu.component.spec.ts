import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IconMenuComponent } from './icon-menu.component';

describe('IconMenuComponent', () => {
  let component: IconMenuComponent;
  let fixture: ComponentFixture<IconMenuComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [IconMenuComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(IconMenuComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

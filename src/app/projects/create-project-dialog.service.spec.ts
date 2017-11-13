import { TestBed, inject } from '@angular/core/testing';

import { CreateProjectDialogService } from './create-project-dialog.service';

describe('CreateProjectDialogService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [CreateProjectDialogService]
    });
  });

  it('should be created', inject([CreateProjectDialogService], (service: CreateProjectDialogService) => {
    expect(service).toBeTruthy();
  }));
});

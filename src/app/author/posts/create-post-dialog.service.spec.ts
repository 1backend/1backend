import { TestBed, inject } from '@angular/core/testing';

import { CreatePostDialogService } from './create-post-dialog.service';

describe('CreatePostDialogService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [CreatePostDialogService]
    });
  });

  it('should be created', inject([CreatePostDialogService], (service: CreatePostDialogService) => {
    expect(service).toBeTruthy();
  }));
});

import { TestBed, inject } from '@angular/core/testing';

import { CreateIssueDialogService } from './create-issue-dialog.service';

describe('CreateIssueDialogService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [CreateIssueDialogService]
    });
  });

  it('should be created', inject([CreateIssueDialogService], (service: CreateIssueDialogService) => {
    expect(service).toBeTruthy();
  }));
});

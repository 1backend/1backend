import { LengthPipe } from './length.pipe';

describe('LengthPipe', () => {
  it('create an instance', () => {
    const pipe = new LengthPipe();
    expect(pipe).toBeTruthy();
  });
});

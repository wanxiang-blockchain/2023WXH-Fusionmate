import { TestBed } from '@angular/core/testing';

import { ChainInfoService } from './chain-info.service';

describe('ChainInfoService', () => {
  let service: ChainInfoService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ChainInfoService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});

import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SystematicCheckupComponent } from './systematic-checkup.component';

describe('SystematicCheckupComponent', () => {
  let component: SystematicCheckupComponent;
  let fixture: ComponentFixture<SystematicCheckupComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [SystematicCheckupComponent]
    });
    fixture = TestBed.createComponent(SystematicCheckupComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

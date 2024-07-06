import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AppointmentManagementComponent } from './appointment-management.component';

describe('AppointmentManagementComponent', () => {
  let component: AppointmentManagementComponent;
  let fixture: ComponentFixture<AppointmentManagementComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [AppointmentManagementComponent]
    });
    fixture = TestBed.createComponent(AppointmentManagementComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

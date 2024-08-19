import { ComponentFixture, TestBed } from '@angular/core/testing';

import { StudentAppointmentManagementComponent } from './student-appointment-management.component';

describe('StudentAppointmentManagementComponent', () => {
  let component: StudentAppointmentManagementComponent;
  let fixture: ComponentFixture<StudentAppointmentManagementComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [StudentAppointmentManagementComponent]
    });
    fixture = TestBed.createComponent(StudentAppointmentManagementComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

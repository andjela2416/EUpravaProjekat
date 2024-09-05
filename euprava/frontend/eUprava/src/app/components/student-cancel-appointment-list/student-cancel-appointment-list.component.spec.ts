import { ComponentFixture, TestBed } from '@angular/core/testing';

import { StudentCancelAppointmentListComponent } from './student-cancel-appointment-list.component';

describe('StudentCancelAppointmentListComponent', () => {
  let component: StudentCancelAppointmentListComponent;
  let fixture: ComponentFixture<StudentCancelAppointmentListComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [StudentCancelAppointmentListComponent]
    });
    fixture = TestBed.createComponent(StudentCancelAppointmentListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

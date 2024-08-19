import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AppointmentListUpdateComponent } from './appointment-list-update.component';

describe('AppointmentListUpdateComponent', () => {
  let component: AppointmentListUpdateComponent;
  let fixture: ComponentFixture<AppointmentListUpdateComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [AppointmentListUpdateComponent]
    });
    fixture = TestBed.createComponent(AppointmentListUpdateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

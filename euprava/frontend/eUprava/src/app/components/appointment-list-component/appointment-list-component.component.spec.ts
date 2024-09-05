import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AppointmentListComponentComponent } from './appointment-list-component.component';

describe('AppointmentListComponentComponent', () => {
  let component: AppointmentListComponentComponent;
  let fixture: ComponentFixture<AppointmentListComponentComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [AppointmentListComponentComponent]
    });
    fixture = TestBed.createComponent(AppointmentListComponentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

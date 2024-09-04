import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-student-appointment-management',
  templateUrl: './student-appointment-management.component.html',
  styleUrls: ['./student-appointment-management.component.css']
})
export class StudentAppointmentManagementComponent {

constructor(private router: Router) {}

  scheduleAppointment() {
    this.router.navigate(['/student-appointment-list']);
  }

  cancelAppointment() {
    this.router.navigate(['/cancel-appointment']);
  }

}

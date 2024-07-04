import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-appointment-management',
  templateUrl: './appointment-management.component.html',
  styleUrls: ['./appointment-management.component.css']
})
export class AppointmentManagementComponent {

  constructor(private router: Router) {}

  createAppointment() {
    this.router.navigate(['/create-appointment']);
  }

  createSystematicCheck() {
    this.router.navigate(['/create-systematicCheck']);
  }
}

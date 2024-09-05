import { Component, OnInit } from '@angular/core';
import { AppointmentService } from 'src/app/services/appointment.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-student-appointment-list',
  templateUrl: './student-appointment-list.component.html',
  styleUrls: ['./student-appointment-list.component.css']
})
export class StudentAppointmentListComponent implements OnInit {
  appointments: any[] = [];

  constructor(private appointmentService: AppointmentService,private router: Router) { }

  ngOnInit(): void {
    this.loadAppointments();
  }

  loadAppointments(): void {
    this.appointmentService.getAppointments().subscribe(
      (data) => {
        this.appointments = data;
      },
      (error) => {
        console.error('Error loading appointments:', error);
        // Handle error as needed
      }
    );
  }

  scheduleAppointment(id: string): void {
    this.appointmentService.scheduleAppointment(id).subscribe(
      () => {
        console.log('Appointment scheduled successfully');
        alert("Appointment scheduled successfully")
        this.router.navigate(['/student-appointment-management']);
      },
      (error) => {
        console.error('Error scheduling appointment:', error);
        // Handle error as needed
      }
    );
  }
}

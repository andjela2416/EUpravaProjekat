import { Component, OnInit } from '@angular/core';
import { AppointmentService } from 'src/app/services/appointment.service';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-student-appointment-list',
  templateUrl: './student-appointment-list.component.html',
  styleUrls: ['./student-appointment-list.component.css']
})
export class StudentAppointmentListComponent implements OnInit {
  appointments: any[] = [];
  userId: string | null = '';

  constructor(private appointmentService: AppointmentService,private router: Router, private authService: AuthService) { }

  ngOnInit(): void {
    this.loadAppointments();
    this.userId = localStorage.getItem('user_id');
  }

  loadAppointments(): void {
    this.appointmentService.getFreeAppointments().subscribe(
      (data) => {
        this.appointments = data;
        this.loadDoctorsInfo();
      },
      (error) => {
        console.error('Error loading appointments:', error);
      }
    );
  }

  loadDoctorsInfo(): void {
    this.appointments.forEach((appointment) => {
      this.authService.getUser(appointment.doctor_id).subscribe(
        (doctor) => {
          appointment.doctor_id = doctor;
        },
        (error) => {
          console.error('Error loading doctor info:', error);
          // Handle error as needed
        }
      );
    });
  }

  scheduleAppointment(id: string): void {
    this.appointmentService.scheduleAppointment(id, this.userId ).subscribe(
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

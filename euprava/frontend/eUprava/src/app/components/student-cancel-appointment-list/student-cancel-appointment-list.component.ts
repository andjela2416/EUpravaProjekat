import { Component, OnInit } from '@angular/core';
import { AppointmentService } from 'src/app/services/appointment.service';
import { AuthService } from 'src/app/services/auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-student-cancel-appointment-list',
  templateUrl: './student-cancel-appointment-list.component.html',
  styleUrls: ['./student-cancel-appointment-list.component.css']
})
export class StudentCancelAppointmentListComponent implements OnInit {
  appointments: any[] = [];

  constructor(private appointmentService: AppointmentService, private router: Router, private authService: AuthService) { }

  ngOnInit(): void {
    this.loadAppointments()
  }

  loadAppointments(): void {
    this.appointmentService.getReservedAppointmentsByStudent(localStorage.getItem('user_id')).subscribe(
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
        }
      );
    });
  }

  cancelAppointment(id: string): void {
    this.appointmentService.cancelAppointment(id).subscribe(
      () => {
        console.log('Appointment canceled successfully');
        alert("Appointment canceled successfully");
        this.loadAppointments(); // Reload the appointments list
      },
      (error) => {
        console.error('Error canceling appointment:', error);
        // Handle error as needed
      }
    );
  }

  isCancelable(appointmentDate: string): boolean {
    const now = new Date();
    const appointment = new Date(appointmentDate);
    const hoursDifference = (appointment.getTime() - now.getTime()) / (1000 * 60 * 60);
    return hoursDifference >= 24;
  }
}


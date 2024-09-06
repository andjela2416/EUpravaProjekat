import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { OnInit } from '@angular/core';
import { AppointmentService } from 'src/app/services/appointment.service';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-appointment-list-update',
  templateUrl: './appointment-list-update.component.html',
  styleUrls: ['./appointment-list-update.component.css']
})
export class AppointmentListUpdateComponent implements OnInit {
  appointments: any[] = [];
  systematicAppointments: any[] = [];

  constructor(private appointmentService: AppointmentService,private router: Router,private authService: AuthService) { }

  ngOnInit(): void {
    this.loadAppointments();
  }

  loadAppointments(): void {
    this.appointmentService.getAppointmentsByDoctor().subscribe(data => {
      this.appointments = data.filter(appointment => !appointment.systematic);
      this.systematicAppointments = data.filter(appointment => appointment.systematic);
      this.loadStudentsInfo();
    });
  }

  loadStudentsInfo(): void {
    this.appointments.forEach((appointment) => {
      this.authService.getUser(appointment.student_id).subscribe(
        (student) => {
          appointment.student_id = student;
        },
        (error) => {
          console.error('Error loading doctor info:', error);
        }
      );
    });
    this.systematicAppointments.forEach((appointment) => {
      this.authService.getUser(appointment.student_id).subscribe(
        (student) => {
          appointment.student_id = student;
        },
        (error) => {
          console.error('Error loading doctor info:', error);
        }
      );
    });
  }

  updateAppointment(appointmentId: string) {
    this.router.navigate(['/update-appointment', appointmentId]);
  }

  deleteAppointment(id: string, appointmentDate: string): void {
    if (!this.rule24h(appointmentDate)) {
      alert('Ne možete obrisati termin zakazan za manje od 24 sata.');
      return;
    }

    if (confirm('Da li ste sigurni da želite da obrišete ovaj termin?')) {
      this.appointmentService.deleteAppointment(id).subscribe(() => {

        this.appointments = this.appointments.filter(app => app.id !== id);
        this.systematicAppointments = this.systematicAppointments.filter(app => app.id !== id);
      });
    }
  }

  rule24h(appointmentDate: string): boolean {
    const now = new Date();
    const appointment = new Date(appointmentDate);
    const hoursDifference = (appointment.getTime() - now.getTime()) / (1000 * 60 * 60);
    return hoursDifference >= 24;
  }

  isFutureDate(appointmentDate: string): boolean {
    const now = new Date();
    const appointment = new Date(appointmentDate);
    return appointment > now;
  }

}

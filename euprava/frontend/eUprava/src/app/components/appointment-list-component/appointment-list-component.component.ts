// appointment-list.component.ts
import { Component, OnInit } from '@angular/core';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AppointmentService } from 'src/app/services/appointment.service';

@Component({
  selector: 'app-appointment-list',
  templateUrl: './appointment-list-component.component.html',
  styleUrls: ['./appointment-list-component.component.css']
})
export class AppointmentListComponent implements OnInit {
  appointments: any[] = [];

  constructor(private appointmentService: AppointmentService) { }

  ngOnInit(): void {
    this.loadAppointments();
  }

  loadAppointments(): void {
    this.appointmentService.getAppointmentsByDoctor().subscribe(data => {
      this.appointments = data;
    });
  }

  deleteAppointment(id: string): void {
    if (confirm('Are you sure you want to delete this appointment?')) {
      this.appointmentService.deleteAppointment(id).subscribe(() => {
        this.appointments = this.appointments.filter(app => app.id !== id);
      });
    }
  }
}

import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { OnInit } from '@angular/core';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AppointmentService } from 'src/app/services/appointment.service';

@Component({
  selector: 'app-appointment-list-update',
  templateUrl: './appointment-list-update.component.html',
  styleUrls: ['./appointment-list-update.component.css']
})
export class AppointmentListUpdateComponent implements OnInit {
  appointments: any[] = [];

  constructor(private appointmentService: AppointmentService,private router: Router) { }

  ngOnInit(): void {
    this.loadAppointments();
  }

  loadAppointments(): void {
    this.appointmentService.getAppointments().subscribe(data => {
      this.appointments = data;
    });
  }

  updateAppointment(appointmentId: string) {
    this.router.navigate(['/update-appointment', appointmentId]);
  }
}

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Appointment } from '../models/appointment.model';
import { TherapyData } from '../models/appointment.model';
import { Observable } from 'rxjs';
import { environment } from 'src/app/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class AppointmentService {

  private url = "healthcare";
  constructor(private http: HttpClient) { }

  createAppointment(appointment: Appointment): Observable<any> {
    return this.http.post<any>(`${environment.baseApiUrl}/${this.url}/appointments`, appointment);
  }

  getAppointments(): Observable<any[]> {
    return this.http.get<any[]>(`${environment.baseApiUrl}/${this.url}/appointments`);
  }

  getAppointment(id: string): Observable<any> {
    return this.http.get<any>(`${environment.baseApiUrl}/${this.url}/appointmentById?id=${id}`);
  }

  updateAppointment(id: string, appointmentData: any): Observable<any> {
    return this.http.patch<any>(`${environment.baseApiUrl}/${this.url}/appointment/update/${id}`, appointmentData);
  }

  deleteAppointment(id: string): Observable<void> {
    return this.http.delete<void>(`${environment.baseApiUrl}/${this.url}/appointment/delete?id=${id}`);
  }

  createTherapy(therapy: TherapyData): Observable<any> {
    return this.http.post<any>(`${environment.baseApiUrl}/${this.url}/therapy`, therapy);
  }

  getFreeAppointments(): Observable<any[]> {
    return this.http.get<any[]>(`${environment.baseApiUrl}/${this.url}/appointments/not_reserved`);
  }

  getReservedAppointments(): Observable<any[]> {
    return this.http.get<any[]>(`${environment.baseApiUrl}/${this.url}/appointments/reserved`);
  }

  getReservedAppointmentsByStudent(): Observable<any[]> {
    return this.http.get<any[]>(`${environment.baseApiUrl}/${this.url}/appointments/reservedByStudent`);
  }

  getAppointmentsByDoctor(): Observable<any[]> {
    return this.http.get<any[]>(`${environment.baseApiUrl}/${this.url}/appointments/byUser`);
  }

  scheduleAppointment(appointmentId: string): Observable<any> {
    const requestBody = { appointment_id: appointmentId };
    return this.http.post<any>(`${environment.baseApiUrl}/${this.url}/appointments/schedule`, requestBody);
  }

  cancelAppointment(appointmentId: string): Observable<any> {
    const requestBody = { appointment_id: appointmentId };
    return this.http.post<any>(`${environment.baseApiUrl}/${this.url}/appointments/cancel`, requestBody);
  }

}

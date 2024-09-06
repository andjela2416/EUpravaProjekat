import {HttpClient, HttpParams} from '@angular/common/http';
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

  createAppointment(appointment: Appointment, userId: string): Observable<any> {
    return this.http.post<any>(`${environment.baseApiUrl}/${this.url}/appointments?doctorId=${userId}`, appointment);
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

  getTherapies(): Observable<any[]> {
    return this.http.get<any[]>(`${environment.baseApiUrl}/${this.url}/therapies`);
  }

  getHealthRecords(): Observable<any[]> {
    return this.http.get<any[]>(`${environment.baseApiUrl}/${this.url}/healthrecords`);
  }

  getFreeAppointments(): Observable<any[]> {
    return this.http.get<any[]>(`${environment.baseApiUrl}/${this.url}/appointments/not_reserved`);
  }

  getReservedAppointments(): Observable<any[]> {
    return this.http.get<any[]>(`${environment.baseApiUrl}/${this.url}/appointments/reserved`);
  }

  getReservedAppointmentsByStudent(userId: string | null ): Observable<any[]> {
    let params = new HttpParams();

    if (userId) {
      params = params.set('user_id', userId);
    }
    return this.http.get<any[]>(`${environment.baseApiUrl}/${this.url}/appointments/reservedByStudent`, { params });
  }

  getAppointmentsByDoctor(): Observable<any[]> {
    return this.http.get<any[]>(`${environment.baseApiUrl}/${this.url}/appointments/byUser`);
  }

  scheduleAppointment(appointmentId: string, userId: string | null): Observable<any> {
    const requestBody = { appointment_id: appointmentId, user_id: userId};
    return this.http.post<any>(`${environment.baseApiUrl}/${this.url}/appointments/schedule`, requestBody);
  }

  cancelAppointment(appointmentId: string): Observable<any> {
    const requestBody = { appointment_id: appointmentId };
    return this.http.post<any>(`${environment.baseApiUrl}/${this.url}/appointments/cancel`, requestBody);
  }

}

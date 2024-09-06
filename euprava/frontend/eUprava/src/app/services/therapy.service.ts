import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { TherapyData } from '../models/appointment.model'; // Assuming you've moved TherapyData to a separate model file
import { Observable } from 'rxjs';
import { environment } from 'src/app/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class TherapyService {

  private url = "food";
  constructor(private http: HttpClient) { }

  // Fetch all therapies
  getAllTherapies(): Observable<TherapyData[]> {
    return this.http.get<TherapyData[]>(`${environment.baseApiUrl}/${this.url}/therapies`);
  }

  // Optionally other CRUD methods can be added here
}

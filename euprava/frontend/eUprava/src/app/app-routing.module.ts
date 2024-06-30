import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CreateAppointmentComponent } from './components/create-appointment/create-appointment.component';
import { AppointmentManagementComponent } from './components/appointment-management/appointment-management.component';

const routes: Routes = [
  { path: 'create-appointment', component: CreateAppointmentComponent },
  { path: '', component: AppointmentManagementComponent }
  // Dodajte ostale rute ovde ako postoje
];


@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

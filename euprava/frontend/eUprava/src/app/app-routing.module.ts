import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CreateAppointmentComponent } from './components/create-appointment/create-appointment.component';
import { AppointmentManagementComponent } from './components/appointment-management/appointment-management.component';
import { SystematicCheckupComponent } from './components/systematic-checkup/systematic-checkup.component';


const routes: Routes = [
  { path: 'create-appointment', component: CreateAppointmentComponent },
  { path: '', component: AppointmentManagementComponent },
  { path: 'create-systematicCheck', component: SystematicCheckupComponent }
  // Dodajte ostale rute ovde ako postoje
];


@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

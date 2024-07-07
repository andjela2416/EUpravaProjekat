import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CreateAppointmentComponent } from './components/create-appointment/create-appointment.component';
import { AppointmentManagementComponent } from './components/appointment-management/appointment-management.component';
import { SystematicCheckupComponent } from './components/systematic-checkup/systematic-checkup.component';
import { AppointmentListComponent } from './components/appointment-list-component/appointment-list-component.component';
import { AppointmentUpdateComponent } from './components/appointment-update/appointment-update.component';
import { AppointmentListUpdateComponent } from './components/appointment-list-update/appointment-list-update.component';
import { TherapyCreateComponent } from './components/create-therapy/create-therapy.component';
import { StudentAppointmentManagementComponent } from './components/student-appointment-management/student-appointment-management.component';
import { StudentAppointmentListComponent } from './components/student-appointment-list/student-appointment-list.component';
import { StudentCancelAppointmentListComponent } from './components/student-cancel-appointment-list/student-cancel-appointment-list.component';

const routes: Routes = [
  { path: 'create-appointment', component: CreateAppointmentComponent },
  { path: '', component: AppointmentManagementComponent },
  { path: 'student-appointment-management', component: StudentAppointmentManagementComponent },
  { path: 'student-appointment-list', component: StudentAppointmentListComponent },
  { path: 'create-systematicCheck', component: SystematicCheckupComponent },
  { path: 'appointment-list', component: AppointmentListComponent },
  { path: 'update-appointment/:id', component: AppointmentUpdateComponent },
  { path: 'update-appointment-list', component: AppointmentListUpdateComponent },
  { path: 'create-therapy', component: TherapyCreateComponent },
  { path: 'cancel-appointment', component: StudentCancelAppointmentListComponent },
];


@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

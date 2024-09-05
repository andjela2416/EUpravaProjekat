import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HttpClientModule} from '@angular/common/http';
import { CreateAppointmentComponent } from './components/create-appointment/create-appointment.component';
import { AppointmentManagementComponent } from './components/appointment-management/appointment-management.component';
import { SystematicCheckupComponent } from './components/systematic-checkup/systematic-checkup.component';
import { AppointmentListComponent } from './components/appointment-list-component/appointment-list-component.component';
import { CommonModule } from '@angular/common';
import { AppointmentUpdateComponent } from './components/appointment-update/appointment-update.component';
import { AppointmentListUpdateComponent } from './components/appointment-list-update/appointment-list-update.component';
import { CreateTherapyComponent } from './components/create-therapy/create-therapy.component';
import { StudentAppointmentManagementComponent } from './components/student-appointment-management/student-appointment-management.component';
import { StudentAppointmentListComponent } from './components/student-appointment-list/student-appointment-list.component';
import { StudentCancelAppointmentListComponent } from './components/student-cancel-appointment-list/student-cancel-appointment-list.component';
import { NavbarComponent } from './components/navbar/navbar.component';
import {LoginComponent} from "./components/login/login.component";
import {RegisterComponent} from "./components/register/register.component";

@NgModule({
  declarations: [
    AppComponent,
    CreateAppointmentComponent,
    AppointmentManagementComponent,
    SystematicCheckupComponent,
    AppointmentListComponent,
    AppointmentUpdateComponent,
    AppointmentListUpdateComponent,
    CreateTherapyComponent,
    StudentAppointmentManagementComponent,
    StudentAppointmentListComponent,
    StudentCancelAppointmentListComponent,
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    AppRoutingModule,
    FormsModule,
    ReactiveFormsModule,
    CommonModule,
    NavbarComponent,
    LoginComponent,
    RegisterComponent
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }

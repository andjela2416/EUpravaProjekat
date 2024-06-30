import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { CreateAppointmentComponent } from './components/create-appointment/create-appointment.component';
import { AppointmentManagementComponent } from './components/appointment-management/appointment-management.component';
import { ToastModule } from 'ngx-toast';

@NgModule({
  declarations: [
    AppComponent,
    CreateAppointmentComponent,
    AppointmentManagementComponent
  ],
  imports: [
    BrowserModule,
    ToastModule.forRoot(),
    HttpClientModule,
    AppRoutingModule,
    FormsModule,
    ReactiveFormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }

import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AppointmentService } from 'src/app/services/appointment.service';
import { Appointment } from 'src/app/models/appointment.model';

@Component({
  selector: 'app-create-appointment',
  templateUrl: './create-appointment.component.html',
  styleUrls: ['./create-appointment.component.css']
})
export class CreateAppointmentComponent implements OnInit {
  appointmentForm: FormGroup;

  constructor(private fb: FormBuilder,private appointmentService: AppointmentService) {
    this.appointmentForm = this.fb.group({
      studentID: [''],
      date: ['', Validators.required],
      doorNumber: ['', Validators.required],
      description: [''],
      systematic: [false],
    });
  }

  ngOnInit(): void {
  }

onSubmit() {

    if (this.appointmentForm.valid) {
      const appointmentData: Appointment = this.appointmentForm.value;
      this.appointmentService.createAppointment(appointmentData)
        .subscribe(
          response => {
            console.log('Appointment created successfully:', response);
            // Implement logic for success handling (e.g., redirect, show success message)
          },
          error => {
            console.error('Error creating appointment:', error);
            // Implement error handling (e.g., show error message)
          }
        );
    } else {
      // Form not valid, handle error or validation messages
    }
  }
}

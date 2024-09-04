import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AppointmentService } from 'src/app/services/appointment.service';
import { Appointment } from 'src/app/models/appointment.model';
import { Router } from '@angular/router';

@Component({
  selector: 'app-create-appointment',
  templateUrl: './create-appointment.component.html',
  styleUrls: ['./create-appointment.component.css']
})
export class CreateAppointmentComponent implements OnInit {
  appointmentForm: FormGroup;

  constructor(private fb: FormBuilder, private appointmentService: AppointmentService, private router: Router) {
    this.appointmentForm = this.fb.group({
      date: ['', Validators.required],
      time: ['', Validators.required],
      doorNumber: ['', Validators.required],
      description: ['', Validators.maxLength(200)] // maksimalno 200 karaktera za opis
    });
  }

  ngOnInit(): void {
  }

  onSubmit() {
    if (this.appointmentForm.valid) {
      const datetime = `${this.appointmentForm.value.date}T${this.appointmentForm.value.time}:00`;
      const dateObj = new Date(datetime);

      const appointmentData: Appointment = {
        studentId: '',
        date: dateObj,
        door_number: this.appointmentForm.value.doorNumber,
        description: this.appointmentForm.value.description,
        systematic: false,
        reserved: false,
        faculty_name: '',
        field_of_study: ''
      };

      console.log(appointmentData);

      this.appointmentService.createAppointment(appointmentData)
        .subscribe(
          response => {
            console.log('Appointment created successfully:', response);
            this.router.navigate(['']);
          },
          error => {
            console.error('Error creating appointment:', error);
            // Implement error handling (e.g., show error message)
          }
        );
    } else {
      Object.keys(this.appointmentForm.controls).forEach(field => {
        const control = this.appointmentForm.get(field);
        if (control) {
          if (control instanceof FormGroup) {
            Object.keys(control.controls).forEach(innerField => {
              const innerControl = control.get(innerField);
              if (innerControl) {
                innerControl.markAsTouched({ onlySelf: true });
              }
            });
          } else {
            control.markAsTouched({ onlySelf: true });
          }
        }
      });

    }
  }

  // Helper method to check if a form control has an error
  hasError(controlName: string, errorName: string) {
    return this.appointmentForm.controls[controlName].hasError(errorName);
  }
}

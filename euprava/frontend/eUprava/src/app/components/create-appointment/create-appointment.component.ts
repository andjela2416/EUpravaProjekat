import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AppointmentService } from 'src/app/services/appointment.service';
import { Appointment } from 'src/app/models/appointment.model';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-create-appointment',
  templateUrl: './create-appointment.component.html',
  styleUrls: ['./create-appointment.component.css']
})
export class CreateAppointmentComponent implements OnInit {
  appointmentForm: FormGroup;
  minDate: string;

  constructor(private fb: FormBuilder, private appointmentService: AppointmentService, private router: Router, private authService: AuthService) {
    this.minDate = this.getMinDate();
    this.appointmentForm = this.fb.group({
      date: ['', Validators.required],
      time: ['', Validators.required],
      doorNumber: ['', Validators.required],
      description: ['', Validators.maxLength(200)] // maksimalno 200 karaktera za opis
    });
  }

  ngOnInit(): void {
  }

  getMinDate(): string {
    const today = new Date();
    const day = today.getDate().toString().padStart(2, '0');
    const month = (today.getMonth() + 1).toString().padStart(2, '0');
    const year = today.getFullYear();
    return `${year}-${month}-${day}`;
  }

  onSubmit() {
    if (this.appointmentForm.valid) {
      const datetime = `${this.appointmentForm.value.date}T${this.appointmentForm.value.time}:00`;
      const dateObj = new Date(datetime);

      const loggedUserId = this.authService.getUserId() || '';

      const appointmentData: Appointment = {
        student_id:'',
        doctor_id: loggedUserId,
        date: dateObj,
        door_number: this.appointmentForm.value.doorNumber,
        description: this.appointmentForm.value.description,
        systematic: false,
        reserved: false,
        faculty_name: '',
        field_of_study: ''
      };

      console.log(appointmentData);

      this.appointmentService.createAppointment(appointmentData,loggedUserId)
        .subscribe(
          response => {
            console.log('Appointment created successfully:');
            this.router.navigate(['appointment-management']);
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

import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Appointment } from 'src/app/models/appointment.model';
import { Router } from '@angular/router';
import { AppointmentService } from 'src/app/services/appointment.service';


@Component({
  selector: 'app-systematic-checkup',
  templateUrl: './systematic-checkup.component.html',
  styleUrls: ['./systematic-checkup.component.css']
})
export class SystematicCheckupComponent {
  checkupForm: FormGroup;

  constructor(private fb: FormBuilder, private appointmentService: AppointmentService, private router:Router) {
    this.checkupForm = this.fb.group({
       studentID: [''],
       date: ['', Validators.required],
       doorNumber: ['', Validators.required],
       description: [''],
       faculty_name:[''],
       field_of_study:['']
    });
  }

  onSubmit() {
      if (this.checkupForm.valid) {
        // Spajamo datum i vreme u jedan string u formatu koji očekujete na backendu
        const datetime = `${this.checkupForm.value.date}:00`;
        const dateObj = new Date(datetime);

        // Kreiramo objekat za slanje na backend
        const appointmentData: Appointment = {
          studentId: '',
          date: dateObj,
          door_number: this.checkupForm.value.doorNumber,
          description: this.checkupForm.value.description,
          systematic: true,
          reserved: true,
          faculty_name:this.checkupForm.value.faculty_name,
          field_of_study:this.checkupForm.value.field_of_study
        };
        console.log(appointmentData);

        this.appointmentService.createAppointment(appointmentData)
          .subscribe(
            response => {
              console.log('Appointment created successfully:', response);
              alert('Systematic Checkup Appointment created successfully:')
              this.router.navigate(['']);
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

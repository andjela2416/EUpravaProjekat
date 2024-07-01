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

  constructor(private fb: FormBuilder, private appointmentService: AppointmentService) {
    this.appointmentForm = this.fb.group({
      studentID: [''],
      date: ['', Validators.required],
      time: ['', Validators.required], // Dodato polje za vreme
      doorNumber: ['', Validators.required],
      description: [''],
      systematic: [false],
    });
  }

  ngOnInit(): void {
  }

  onSubmit() {
    if (this.appointmentForm.valid) {
      // Spajamo datum i vreme u jedan string u formatu koji očekujete na backendu
      const datetime = `${this.appointmentForm.value.date}T${this.appointmentForm.value.time}:00`;
      const dateObj = new Date(datetime);

      // Kreiramo objekat za slanje na backend
      const appointmentData: Appointment = {
        studentId: this.appointmentForm.value.studentID,
        date: dateObj,
        door_number: this.appointmentForm.value.doorNumber,
        description: this.appointmentForm.value.description,
        systematic: this.appointmentForm.value.systematic,
      };
      console.log(appointmentData);

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

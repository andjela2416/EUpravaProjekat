import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AppointmentService } from 'src/app/services/appointment.service';

@Component({
  selector: 'app-create-therapy',
  templateUrl: './create-therapy.component.html',
  styleUrls: ['./create-therapy.component.css']
})
export class TherapyCreateComponent implements OnInit {
  therapyForm: FormGroup;

  constructor(
    private fb: FormBuilder,
    private router: Router,
    private appointmentService: AppointmentService
  ) {
    this.therapyForm = this.fb.group({
      studentHealthRecordId: ['', [Validators.required, Validators.minLength(24), Validators.maxLength(24)]],
      diagnosis: ['', Validators.required]
    });
  }

  ngOnInit(): void {
    // Možete ovde inicijalizovati komponentu ako je potrebno
  }

  onSubmit(): void {
    if (this.therapyForm.valid) {
      console.log(this.therapyForm.value);

      this.appointmentService.createTherapy(this.therapyForm.value)
        .subscribe(
          response => {
            console.log('Therapy created successfully', response);
            // Primer preusmeravanja na odgovarajuću stranicu nakon kreiranja terapije
            this.router.navigate(['']); // Preusmeravanje na listu terapija ili drugu stranicu
          },
          error => {
            console.error('Error creating therapy', error);
          }
        );
    }
  }
}

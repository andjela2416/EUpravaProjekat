import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AppointmentService } from 'src/app/services/appointment.service';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-create-therapy',
  templateUrl: './create-therapy.component.html',
  styleUrls: ['./create-therapy.component.css']
})
export class CreateTherapyComponent implements OnInit {
  therapyForm: FormGroup;
  healthRecords: any[] = [];
  therapies: any[] = [];

  constructor(
    private fb: FormBuilder,
    private router: Router,
    private appointmentService: AppointmentService,
    private authService: AuthService
  ) {
    this.therapyForm = this.fb.group({
      studentHealthRecordId: ['', [Validators.required]],
      diagnosis: ['', Validators.required]
    });
  }

  ngOnInit(): void {
    this.loadHealthRecords();
    this.loadTherapies(); // Load therapies on init
  }

  loadTherapies(): void {
    this.appointmentService.getTherapies().subscribe(
      data => {
        this.therapies = data;
      },
      error => {
        console.error('Error loading therapies', error);
      }
    );
  }
  loadStudentsInfo(): void {
    this.healthRecords.forEach((hr) => {
      this.authService.getUser(hr.userId).subscribe(
        (student) => {
          hr.userId = student;
        },
        (error) => {
          console.error('Error loading doctor info:', error);
        }
      );
    });
  }

  loadHealthRecords(): void {
    this.appointmentService.getHealthRecords().subscribe(
      data => {
        this.healthRecords = data;
        this.loadStudentsInfo();
      },
      error => {
        console.error('Error loading health records', error);
      }
    );
  }

  onSubmit(): void {
    if (this.therapyForm.valid) {

      this.appointmentService.createTherapy(this.therapyForm.value)
        .subscribe(
          response => {
            console.log('Therapy created successfully', response);
            alert('Therapy created successfully')
            this.router.navigate(['appointment-management']);
          },
          error => {
            console.error('Error creating therapy', error);
          }
        );
    }
  }
}

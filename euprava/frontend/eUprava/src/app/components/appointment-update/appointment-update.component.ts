import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { ActivatedRoute, Router } from '@angular/router';
import { AppointmentService } from 'src/app/services/appointment.service';
import { DatePipe } from '@angular/common'; // Import DatePipe
import { Appointment } from 'src/app/models/appointment.model';

@Component({
  selector: 'app-appointment-update',
  templateUrl: './appointment-update.component.html',
  styleUrls: ['./appointment-update.component.css'],
  providers: [DatePipe] // Dodaj DatePipe kao providera ovde
})
export class AppointmentUpdateComponent implements OnInit {
  appointmentForm: FormGroup;
  appointmentId: string;
  minDate: string;

  constructor(
    private fb: FormBuilder,
    private http: HttpClient,
    private route: ActivatedRoute,
    private router: Router,
    private appointmentService: AppointmentService,
    private datePipe: DatePipe // Injektuj DatePipe
  ) {
    this.minDate = this.getMinDate();
    this.appointmentForm = this.fb.group({
      date: ['', Validators.required],
      time: ['', Validators.required],
      studentId: [''],
      door_number: ['', Validators.required],
      description: [''],
      systematic: [false],
      faculty_name: [''],
      field_of_study: [''],
      reserved: [false]
    });

    this.appointmentId = '';
  }

  ngOnInit(): void {
    this.route.paramMap.subscribe(params => {
      this.appointmentId = params.get('id') || '';
      this.loadAppointment();
    });
  }

  getMinDate(): string {
    const today = new Date();
    const day = today.getDate().toString().padStart(2, '0');
    const month = (today.getMonth() + 1).toString().padStart(2, '0');
    const year = today.getFullYear();
    return `${year}-${month}-${day}`;
  }

  loadAppointment(): void {
    this.appointmentService.getAppointment(this.appointmentId)
      .subscribe((appointment: any) => {
        // Formatiraj datum pre postavljanja u formu
        const formattedDate = this.datePipe.transform(appointment.date, 'yyyy-MM-dd');
        const formattedTime = this.datePipe.transform(appointment.date, 'HH:mm');
        this.appointmentForm.patchValue({
          date: formattedDate,
          time: formattedTime,
          studentId: appointment.studentId,
          door_number: appointment.door_number,
          description: appointment.description,
          systematic: appointment.systematic,
          faculty_name: appointment.faculty_name,
          field_of_study: appointment.field_of_study,
          reserved: appointment.reserved
        });
      });
  }

  updateAppointment(): void {
    if (this.appointmentForm.valid) {
      const datetime = `${this.appointmentForm.value.date}T${this.appointmentForm.value.time}:00`;
      const dateObj = new Date(datetime);

      const appointmentData: Appointment = {
        student_id: this.appointmentForm.value.studentId,
        doctor_id: this.appointmentForm.value.doctorId,
        date: dateObj,
        door_number: this.appointmentForm.value.door_number,
        description: this.appointmentForm.value.description,
        systematic: this.appointmentForm.value.systematic,
        reserved:this.appointmentForm.value.reserved,
        faculty_name:this.appointmentForm.value.faculty_name,
        field_of_study:this.appointmentForm.value.field_of_study
      };

      this.appointmentService.updateAppointment(this.appointmentId,appointmentData)
        .subscribe(response => {
          console.log('Appointment updated', response);
          this.router.navigate(['/update-appointment-list']);
        }, error => {
          console.error('Error updating appointment', error);
        });
    }
  }

  hasError(controlName: string, errorName: string) {
    return this.appointmentForm.controls[controlName].hasError(errorName);
  }

  goBack() {
    this.router.navigate(['/update-appointment-list']);
  }
}

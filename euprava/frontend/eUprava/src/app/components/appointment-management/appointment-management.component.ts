import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { ToastService } from 'ngx-toast';

@Component({
  selector: 'app-appointment-management',
  templateUrl: './appointment-management.component.html',
  styleUrls: ['./appointment-management.component.css']
})
export class AppointmentManagementComponent {

  constructor(private router: Router,private toast: ToastService,) {}

  createAppointment() {

  this.toast.success('Uspešno ste kreirali termin za zdravstveni pregled!', 'Termin kreiran', {
        timeOut: 3000, // trajanje poruke u ms
        progressBar: true, // prikazivanje progres bara
        closeButton: true // prikazivanje dugmeta za zatvaranje
      });

    this.router.navigate(['/create-appointment']);
  }
}

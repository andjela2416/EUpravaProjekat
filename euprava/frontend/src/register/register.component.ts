import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent {
  registerForm: FormGroup;
  registrationSuccess: boolean = false;
  registrationError: string = '';

  constructor(private fb: FormBuilder, private http: HttpClient) {
    this.registerForm = this.fb.group({
      first_name: ['', Validators.required],
      last_name: ['', Validators.required],
      email: ['', [Validators.required, Validators.email]],
      password: ['', Validators.required],
      phone: ['', Validators.required],
      address: ['', Validators.required],
      user_type: ['', Validators.required]
    });
  }

  onSubmit(): void {
    if (this.registerForm.valid) {
      this.http.post('http://localhost:8080/users/register', this.registerForm.value)
        .subscribe(
          response => {
            console.log(response);
            this.registrationSuccess = true;
            this.registrationError = '';
          },
          error => {
            console.error(error);
            this.registrationSuccess = false;
            this.registrationError = 'Registracija nije uspela. Pokušajte ponovo.';
          }
        );
    }
  }
}

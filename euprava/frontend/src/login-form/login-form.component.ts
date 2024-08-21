import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login-form',
  templateUrl: './login-form.component.html',
  styleUrls: ['./login-form.component.css']
})
export class LoginFormComponent {
  email: string = '';
  password: string = '';
  error: string = '';
  successMessage: string = '';

  constructor(private http: HttpClient, private router: Router) {}

  handleSubmit() {
    this.error = '';
    this.successMessage = '';

    this.http.post('http://localhost:8080/users/login', { email: this.email, password: this.password })
      .subscribe({
        next: (response: any) => {
          localStorage.setItem('token', response.token);
          localStorage.setItem('user_type', response.user_type);
          localStorage.setItem('user_email', this.email);
          this.successMessage = 'Uspešno ste ulogovani.';
        },
        error: (err) => {
          this.error = 'Pogrešni podaci za prijavu.';
        }
      });
  }
}

import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { CommonModule } from '@angular/common';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-navbar',
  standalone: true,
  imports: [CommonModule], // Add CommonModule to imports
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css'] // Correct 'styleUrl' to 'styleUrls'
})
export class NavbarComponent {
  isLoggedIn: boolean = false;
  userType$ = this.authService.userType$;

  constructor(private router: Router, private authService: AuthService) {}

  ngOnInit(): void {
    this.authService.isLoggedIn$.subscribe((loggedIn: boolean) => {
      this.isLoggedIn = loggedIn;
    });
  }

  logout(): void {
    this.authService.logout();
    this.router.navigate(['/login']);
  }
}

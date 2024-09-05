import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { CommonModule } from '@angular/common'; // Import CommonModule for NgIf

@Component({
  selector: 'app-navbar',
  standalone: true,
  imports: [CommonModule], // Add CommonModule to imports
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css'] // Correct 'styleUrl' to 'styleUrls'
})
export class NavbarComponent {
  isLoggedIn: boolean = false;

  constructor(private router: Router) {}

  logout(): void {
    this.router.navigate(['/login']);
  }
}

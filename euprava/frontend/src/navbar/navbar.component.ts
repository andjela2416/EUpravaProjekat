import { Component, OnInit, OnDestroy } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit, OnDestroy {
  userRole: string = '';

  constructor(private router: Router) {}

  ngOnInit(): void {
    document.body.addEventListener('click', this.handleOutsideClick);
    this.userRole = localStorage.getItem('user_type') || '';
  }

  ngOnDestroy(): void {
    document.body.removeEventListener('click', this.handleOutsideClick);
  }

  handleOutsideClick = () => {
    this.userRole = localStorage.getItem('user_type') || '';
  };

  handleLogout(): void {
    localStorage.clear();
    this.router.navigate(['/login']);
  }
}

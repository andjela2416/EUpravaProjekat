import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-home-radnik',
  templateUrl: './home-radnik.component.html',
  styleUrls: ['./home-radnik.component.css']
})
export class HomeRadnikComponent {

  constructor(private router: Router) {}
  
    view() {
      this.router.navigate(['/therapy-list']);
    }
  
   
  
  }
  

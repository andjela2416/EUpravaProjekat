import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private isLoggedInSubject = new BehaviorSubject<boolean>(this.hasToken());
  isLoggedIn$ = this.isLoggedInSubject.asObservable();

  private userTypeSubject = new BehaviorSubject<string | null>(this.getUserType());
  userType$ = this.userTypeSubject.asObservable();

  constructor() {}

  hasToken(): boolean {
    return !!localStorage.getItem('token');
  }

  getUserType(): string | null {
    return localStorage.getItem('userType');
  }

  getUserId(): string | null {
    return localStorage.getItem('user_id');
  }

  login(userType: string): void {
    localStorage.setItem('userType', userType);
    console.log(localStorage.getItem('userType'))
    this.isLoggedInSubject.next(true);
    this.userTypeSubject.next(userType);
    console.log(this.userTypeSubject)
    console.log(this.userType$)
  }

  logout(): void {
    localStorage.removeItem('token');
    this.isLoggedInSubject.next(false);
    localStorage.removeItem('userType');
    this.userTypeSubject.next(null);
    localStorage.removeItem('user_id');
  }
}

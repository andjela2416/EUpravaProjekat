import { Routes } from '@angular/router';
import { HomeComponent } from '../home/home.component';
import { RegisterComponent } from '../register/register.component';
import { DormsComponent } from '../dorms/dorms.component';
import { FoodComponent } from '../food/food.component';
import { HealthCareComponent } from '../healthcare/healthcare.component';

export const appRoutes: Routes = [
  { path: '', component: HomeComponent }, 
  { path: 'register', component: RegisterComponent },
  { path: 'dorms', component: DormsComponent }, 
  { path: 'food', component: FoodComponent }, 
  { path: 'health-care', component: HealthCareComponent }, 
  { path: '**', redirectTo: '', pathMatch: 'full' } 
];

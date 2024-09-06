import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HomeRadnikComponent } from './home-radnik.component';

describe('HomeRadnikComponent', () => {
  let component: HomeRadnikComponent;
  let fixture: ComponentFixture<HomeRadnikComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [HomeRadnikComponent]
    });
    fixture = TestBed.createComponent(HomeRadnikComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

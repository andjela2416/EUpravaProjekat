import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateTherapyComponent } from './create-therapy.component';

describe('CreateTherapyComponent', () => {
  let component: CreateTherapyComponent;
  let fixture: ComponentFixture<CreateTherapyComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [CreateTherapyComponent]
    });
    fixture = TestBed.createComponent(CreateTherapyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

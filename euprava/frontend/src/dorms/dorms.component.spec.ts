import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DormsComponent } from './dorms.component';

describe('DormsComponent', () => {
  let component: DormsComponent;
  let fixture: ComponentFixture<DormsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [DormsComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(DormsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

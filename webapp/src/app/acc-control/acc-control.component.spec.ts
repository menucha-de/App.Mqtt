import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AccControlComponent } from './acc-control.component';

describe('AccControlComponent', () => {
  let component: AccControlComponent;
  let fixture: ComponentFixture<AccControlComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AccControlComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AccControlComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

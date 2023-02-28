import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SideNavComponent } from './side-nav.component';

describe('SideNavComponent', () => {
  let component: SideNavComponent;
  let fixture: ComponentFixture<SideNavComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SideNavComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SideNavComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

//Reference for this test: https://angular.io/guide/testing-components-basics
describe('SideNavComp', () => {
  it('#clicked() should toggle #isOpen', () => {
    const comp = new SideNavComponent();
    expect(comp.isOpen)
      .withContext('closed at first')
      .toBe(false);
    comp.clicked();
    expect(comp.isOpen)
      .withContext('open after click')
      .toBe(true);
    comp.clicked();
    expect(comp.isOpen)
      .withContext('closed after second click')
      .toBe(false);
  });

  it('#clicked() should set #message to "is open"', () => {
    const comp = new SideNavComponent();
    expect(comp.message)
      .withContext('closed at first')
      .toMatch(/is closed/i);
    comp.clicked();
    expect(comp.message)
      .withContext('open after clicked')
      .toMatch(/is open/i);
  });
});

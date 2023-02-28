import { ComponentFixture,TestBed } from '@angular/core/testing';
import { Title } from '@angular/platform-browser';
//import { utils } from 'mocha';
import { FullCalendarComponent } from '../full-calendar/full-calendar.component';
import { CalendarComponent } from './calendar.component';
import { INITIAL_EVENTS } from './event-utils';

let calendar: CalendarComponent;
let fixture: ComponentFixture<CalendarComponent>;
let events: FullCalendarComponent;

describe('CalendarComponent', () => {
  
  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CalendarComponent ],
    }).compileComponents();

    fixture = TestBed.createComponent(CalendarComponent);
    calendar = fixture.componentInstance;
    fixture.detectChanges();
  });
  
  it('should create the calendar', () => {
    expect(calendar).toBeTruthy();
  });

  
  
});

//Reference for this test: https://angular.io/guide/testing-components-basics
describe('Calendar Toggle Weekends', () => {
  it('#clicked() should toggle #hasWeekends', () => {
    const comp = new FullCalendarComponent();
    expect(comp.hasWeekends)
      .withContext('checked at first')
      .toBe(true);
    comp.clicked();
    expect(comp.hasWeekends)
      .withContext('unchecked after click')
      .toBe(false);
    comp.clicked();
    expect(comp.hasWeekends)
      .withContext('checked after second click')
      .toBe(true);
  });

  it('#clicked() should set #message to "has weekends"', () => {
    const comp = new FullCalendarComponent();
    expect(comp.message)
      .withContext('checked at first')
      .toMatch(/has weekends/i);
    comp.clicked();
    expect(comp.message)
      .withContext('unchecked after clicked')
      .toMatch(/has no weekends/i);
  });
});

describe('Calendar Event', () => {
  it('should add a new event', () => {
    spyOn(calendar, 'handleDateSelect');
    expect(calendar.handleDateSelect).toHaveBeenCalled();
  });

  it('should delete an event', () => {
    //spyOn(calendar, 'handleDateClick');
    //expect(calendar.handleDateClick).toHaveBeenCalled();
  });
});

import { ComponentFixture,TestBed } from '@angular/core/testing';
import { Title } from '@angular/platform-browser';
import { FullCalendarComponent } from '../full-calendar/full-calendar.component';
import { CalendarComponent } from './calendar.component';

describe('CalendarComponent', () => {
  let calendar: CalendarComponent;
  let fixture: ComponentFixture<CalendarComponent>;
  let events: FullCalendarComponent;
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

describe('Calendar Toggle Weekends', () => {
  it('#clicked() should toggle #hasWeekends', () => {
    const comp = new FullCalendarComponent();
    expect(comp.hasWeekends)
      .withContext('checked at first')
      .toBe(false);
    comp.clicked();
    expect(comp.hasWeekends)
      .withContext('unchecked after click')
      .toBe(true);
    comp.clicked();
    expect(comp.hasWeekends)
      .withContext('checked after second click')
      .toBe(false);
  });

  it('#clicked() should set #message to "has weekends"', () => {
    const comp = new FullCalendarComponent();
    expect(comp.message)
      .withContext('checked at first')
      .toMatch(/has weekends/i);
    comp.clicked();
    expect(comp.message)
      .withContext('unchecked after clicked')
      .toMatch(/had no weekends/i);
  });
});


describe('Calendar Event', () => {
  it('should create a new event', () => {
    //const title = ${clickInfo.event.title};
    //expect(title).changes;
  });

  it('should delete an event', () => {
    //const title = ${clickInfo.event.title};
    //expect(title).changes;
  });
});

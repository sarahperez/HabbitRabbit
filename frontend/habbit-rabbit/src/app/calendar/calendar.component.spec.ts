import { TestBed } from '@angular/core/testing';
import { Title } from '@angular/platform-browser';
import { CalendarComponent } from './calendar.component';

describe('CalendarComponent', () => {
  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        CalendarComponent
      ],
    }).compileComponents();
  });

  /**
  it('should create the app', () => {
    const fixture = TestBed.createComponent(CalendarComponent);
    const app = fixture.componentInstance;
    expect(app).toBeTruthy();
  });
  */
});

describe('Calendar Event', () => {
  it('should create a new event', () => {
    //const title = ${clickInfo.event.title};
    //expect(title).changes;
  });
});

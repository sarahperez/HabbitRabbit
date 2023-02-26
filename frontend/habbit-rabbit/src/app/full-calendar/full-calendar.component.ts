import { Component } from '@angular/core';

@Component({
  selector: 'app-full-calendar',
  templateUrl: './full-calendar.component.html',
  styleUrls: ['./full-calendar.component.css']
})
export class FullCalendarComponent {
  hasWeekends = true;
  clicked() { this.hasWeekends = !this.hasWeekends; }
  get message() { return `The sidebar is ${this.hasWeekends ? 'has weekends' : 'has no weekends'}`; }
}

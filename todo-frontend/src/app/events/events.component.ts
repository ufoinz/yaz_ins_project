import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { NgFor, DatePipe, NgIf, NgClass } from '@angular/common';
import { Router } from '@angular/router';
import { TodoService } from '../services/todo.service';

@Component({
  selector: 'app-events',
  standalone: true,
  imports: [FormsModule, NgFor, NgIf, DatePipe, NgClass],
  templateUrl: './events.component.html',
  styleUrl: './events.component.scss'
})
export class EventsComponent {
  events: any[] = [];
  newEvent = { name: '', content: '', time: ''};
  editEvent: any = null;
  token: string = '';

  constructor(private todoService: TodoService, private router: Router) {
    this.token = localStorage.getItem('token') || '';

    if (!this.token) {
      this.router.navigate(['/login']);
    }

    this.loadEvents()
  }

  loadEvents() {
    this.todoService.getEvents(this.token).subscribe(res => {
      this.events = res as any[];
    });
  }

  showCreateForm = false;

createEvent() {
  const eventToSend = {
    name: this.newEvent.name,
    content: this.newEvent.content,
    time: new Date(this.newEvent.time).toISOString()
  };

  this.todoService.createEvent(eventToSend, this.token).subscribe({
    next: () => {
      this.loadEvents();
      this.newEvent = { name: '', content: '', time: '' };
      this.showCreateForm = false;
    },
    error: (err) => {
      console.error("Event create error:", err);
    }
  }); 
}

  deleteEvent(id: number) {
    this.todoService.deleteEvent(id, this.token).subscribe(() => {
      this.loadEvents();
    });
  }

  startEdit(event: any) {
    this.editEvent = { ...event, time: this.formatDateForInput(event.time) };
  }

  saveEdit() {
    const eventToSend = {
      ...this.editEvent,
      time: new Date(this.editEvent.time).toISOString()
    };

    this.todoService.updateEvent(this.editEvent.id, eventToSend, this.token).subscribe(() => {
      this.loadEvents();
      this.editEvent = null;
    });
  }

  cancelEdit() {
    this.editEvent = null;
  }

  private formatDateForInput(date: string): string {
    const d = new Date(date);
    return d.toISOString().slice(0, 16);
  }

  logout() {
    localStorage.removeItem('token');
    this.router.navigate(['/login']);
  }

}
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class TodoService {
  private baseUrl = 'http://localhost:8080/api/v1'

  constructor(private http: HttpClient) {}
  
  register(user: { name: string; email: string; password: string }) {
    return this.http.post(`${this.baseUrl}/users/register`,user);
  }

  login(credentials: { email: string; password: string }) {
    return this.http.post<{ token: string }>(`${this.baseUrl}/users/login`, credentials);
  }

  getEvents(token: string) {
    const headers = new HttpHeaders().set('Authorization', `Bearer ${token}`);
    return this.http.get(`${this.baseUrl}/events/`, { headers });
  }

  createEvent(event: {name: string; content: string; time: string }, token: string) {
    const headers = new HttpHeaders().set('Authorization', `Bearer ${token}`);
    return this.http.post(`${this.baseUrl}/events/`, event, { headers });
  }

  updateEvent(id: number, event: { name: string; content: string; time: string }, token: string) {
    const headers = new HttpHeaders().set('Authorization', `Bearer ${token}`);
    return this.http.put(`${this.baseUrl}/events/${id}`, event, { headers });
  }

  deleteEvent(id: number, token: string) {
    const headers = new HttpHeaders().set('Authorization', `Bearer ${token}`);
    return this.http.delete(`${this.baseUrl}/events/${id}`, { headers })
  }
}

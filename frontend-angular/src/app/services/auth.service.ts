import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { tap } from 'rxjs/operators';

export interface LoginResponse {
  token: string;
  username: string;
}

@Injectable({ providedIn: 'root' })
export class AuthService {
  // Go Backend (Port 5002) — Login Only
  private readonly BASE_URL = 'http://localhost:5002/api/auth';

  constructor(private http: HttpClient) {}

  login(username: string, password: string): Observable<LoginResponse> {
    return this.http
      .post<LoginResponse>(`${this.BASE_URL}/login`, { username, password })
      .pipe(
        tap((res) => {
          localStorage.setItem('jwt_token', res.token);
          localStorage.setItem('username', res.username);
        })
      );
  }

  logout(): void {
    localStorage.removeItem('jwt_token');
    localStorage.removeItem('username');
  }

  isLoggedIn(): boolean {
    return !!localStorage.getItem('jwt_token');
  }

  getUsername(): string {
    return localStorage.getItem('username') ?? '';
  }
}

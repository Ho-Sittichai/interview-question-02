import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { AuthService } from '../services/auth.service';

@Component({
  selector: 'app-welcome',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './welcome.component.html',
  styleUrls: ['./welcome.component.css']
})
export class WelcomeComponent implements OnInit {
  username = '';

  constructor(private auth: AuthService, private router: Router) {}

  ngOnInit(): void {
    this.username = this.auth.getUsername();
  }

  logout(): void {
    this.auth.logout();
    this.router.navigate(['/login']);
  }
}

import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { Observable } from 'rxjs';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatListModule } from '@angular/material/list';
import { MatMenuModule } from '@angular/material/menu';
import { AuthService } from '../../core/auth/auth.service';

@Component({
  selector: 'app-dashboard',
  standalone: true,
  imports: [
    CommonModule,
    MatCardModule,
    MatButtonModule,
    MatIconModule,
    MatToolbarModule,
    MatSidenavModule,
    MatListModule,
    MatMenuModule
  ],
  template: `
    <mat-sidenav-container class="sidenav-container">
      <mat-sidenav #drawer class="sidenav" fixedInViewport
          [attr.role]="(isHandset$ | async) ? 'dialog' : 'navigation'"
          [mode]="(isHandset$ | async) ? 'over' : 'side'"
          [opened]="(isHandset$ | async) === false">
        <mat-toolbar>القائمة</mat-toolbar>
        <mat-nav-list>
          <a mat-list-item routerLink="/dashboard" routerLinkActive="active">
            <mat-icon>dashboard</mat-icon>
            <span>لوحة التحكم</span>
          </a>
          <a mat-list-item routerLink="/customers" routerLinkActive="active">
            <mat-icon>people</mat-icon>
            <span>العملاء</span>
          </a>
          <a mat-list-item routerLink="/orders" routerLinkActive="active">
            <mat-icon>shopping_cart</mat-icon>
            <span>الطلبات</span>
          </a>
          <a mat-list-item routerLink="/inventory" routerLinkActive="active">
            <mat-icon>inventory</mat-icon>
            <span>المخزون</span>
          </a>
          <a mat-list-item routerLink="/reports" routerLinkActive="active">
            <mat-icon>assessment</mat-icon>
            <span>التقارير</span>
          </a>
          <a mat-list-item routerLink="/settings" routerLinkActive="active">
            <mat-icon>settings</mat-icon>
            <span>الإعدادات</span>
          </a>
        </mat-nav-list>
      </mat-sidenav>
      <mat-sidenav-content>
        <mat-toolbar color="primary">
          <button
            type="button"
            aria-label="Toggle sidenav"
            mat-icon-button
            (click)="drawer.toggle()"
            *ngIf="isHandset$ | async">
            <mat-icon aria-label="Side nav toggle icon">menu</mat-icon>
          </button>
          <span>ELKHAWAGA ERP</span>
          <span class="spacer"></span>
          <button mat-icon-button [matMenuTriggerFor]="menu">
            <mat-icon>account_circle</mat-icon>
          </button>
          <mat-menu #menu="matMenu">
            <button mat-menu-item (click)="logout()">
              <mat-icon>exit_to_app</mat-icon>
              <span>تسجيل الخروج</span>
            </button>
          </mat-menu>
        </mat-toolbar>
        <div class="sidenav-content">
          <div class="dashboard-container">
            <h1>مرحباً بك في نظام ELKHAWAGA ERP</h1>
            
            <div class="stats-grid">
              <mat-card class="stat-card">
                <mat-card-header>
                  <mat-card-title>إجمالي العملاء</mat-card-title>
                  <mat-card-subtitle>عدد العملاء المسجلين</mat-card-subtitle>
                </mat-card-header>
                <mat-card-content>
                  <div class="stat-number">1,234</div>
                </mat-card-content>
              </mat-card>

              <mat-card class="stat-card">
                <mat-card-header>
                  <mat-card-title>الطلبات الجديدة</mat-card-title>
                  <mat-card-subtitle>الطلبات في انتظار المعالجة</mat-card-subtitle>
                </mat-card-header>
                <mat-card-content>
                  <div class="stat-number">56</div>
                </mat-card-content>
              </mat-card>

              <mat-card class="stat-card">
                <mat-card-header>
                  <mat-card-title>المبيعات الشهرية</mat-card-title>
                  <mat-card-subtitle>إجمالي المبيعات هذا الشهر</mat-card-subtitle>
                </mat-card-header>
                <mat-card-content>
                  <div class="stat-number">$45,678</div>
                </mat-card-content>
              </mat-card>

              <mat-card class="stat-card">
                <mat-card-header>
                  <mat-card-title>المنتجات المتاحة</mat-card-title>
                  <mat-card-subtitle>عدد المنتجات في المخزون</mat-card-subtitle>
                </mat-card-header>
                <mat-card-content>
                  <div class="stat-number">789</div>
                </mat-card-content>
              </mat-card>
            </div>

            <div class="quick-actions">
              <h2>إجراءات سريعة</h2>
              <div class="actions-grid">
                <button mat-raised-button color="primary" routerLink="/customers">
                  <mat-icon>add</mat-icon>
                  إضافة عميل جديد
                </button>
                <button mat-raised-button color="accent" routerLink="/orders">
                  <mat-icon>add_shopping_cart</mat-icon>
                  إنشاء طلب جديد
                </button>
                <button mat-raised-button color="warn" routerLink="/inventory">
                  <mat-icon>inventory_2</mat-icon>
                  إدارة المخزون
                </button>
                <button mat-raised-button routerLink="/reports">
                  <mat-icon>analytics</mat-icon>
                  عرض التقارير
                </button>
              </div>
            </div>
          </div>
        </div>
      </mat-sidenav-content>
    </mat-sidenav-container>
  `,
  styles: [`
    .sidenav-container {
      height: 100vh;
    }

    .sidenav {
      width: 250px;
    }

    .sidenav .mat-toolbar {
      background: inherit;
    }

    .mat-toolbar.mat-primary {
      position: sticky;
      top: 0;
      z-index: 1;
    }

    .spacer {
      flex: 1 1 auto;
    }

    .sidenav-content {
      padding: 20px;
      background-color: #f5f5f5;
      min-height: calc(100vh - 64px);
    }

    .dashboard-container {
      max-width: 1200px;
      margin: 0 auto;
    }

    .dashboard-container h1 {
      color: #333;
      margin-bottom: 30px;
      text-align: center;
    }

    .stats-grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
      gap: 20px;
      margin-bottom: 40px;
    }

    .stat-card {
      text-align: center;
      transition: transform 0.2s;
    }

    .stat-card:hover {
      transform: translateY(-5px);
    }

    .stat-number {
      font-size: 2.5rem;
      font-weight: bold;
      color: #1976d2;
      margin: 20px 0;
    }

    .quick-actions {
      background: white;
      padding: 30px;
      border-radius: 8px;
      box-shadow: 0 2px 8px rgba(0,0,0,0.1);
    }

    .quick-actions h2 {
      color: #333;
      margin-bottom: 20px;
      text-align: center;
    }

    .actions-grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
      gap: 15px;
    }

    .actions-grid button {
      height: 60px;
      font-size: 1rem;
    }

    .actions-grid mat-icon {
      margin-left: 8px;
    }

    .active {
      background-color: rgba(25, 118, 210, 0.1);
      color: #1976d2;
    }

    @media (max-width: 768px) {
      .stats-grid {
        grid-template-columns: 1fr;
      }
      
      .actions-grid {
        grid-template-columns: 1fr;
      }
    }
  `]
})
export class DashboardComponent {
  isHandset$: Observable<boolean> = new Observable<boolean>(observer => observer.next(false)); // Simplified for now

  constructor(
    private authService: AuthService,
    private router: Router
  ) {}

  logout(): void {
    this.authService.logout();
    this.router.navigate(['/login']);
  }
} 
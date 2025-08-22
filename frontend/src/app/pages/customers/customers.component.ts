import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import { Observable } from 'rxjs';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatListModule } from '@angular/material/list';
import { MatTableModule } from '@angular/material/table';
import { MatPaginatorModule } from '@angular/material/paginator';
import { MatSortModule } from '@angular/material/sort';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select';
import { MatChipsModule } from '@angular/material/chips';
import { MatMenuModule } from '@angular/material/menu';
import { AuthService } from '../../core/auth/auth.service';

@Component({
  selector: 'app-customers',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    MatCardModule,
    MatButtonModule,
    MatIconModule,
    MatToolbarModule,
    MatSidenavModule,
    MatListModule,
    MatTableModule,
    MatPaginatorModule,
    MatSortModule,
    MatFormFieldModule,
    MatInputModule,
    MatSelectModule,
    MatChipsModule,
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
          <span>إدارة العملاء</span>
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
          <div class="customers-container">
            <div class="header-actions">
              <h1>إدارة العملاء</h1>
              <div class="actions">
                <button mat-raised-button color="primary" (click)="addCustomer()">
                  <mat-icon>add</mat-icon>
                  إضافة عميل جديد
                </button>
                <button mat-raised-button color="accent" (click)="exportCustomers()">
                  <mat-icon>download</mat-icon>
                  تصدير
                </button>
                <button mat-raised-button (click)="importCustomers()">
                  <mat-icon>upload</mat-icon>
                  استيراد
                </button>
              </div>
            </div>

            <mat-card class="filters-card">
              <mat-card-content>
                <div class="filters-grid">
                  <mat-form-field appearance="outline">
                    <mat-label>البحث</mat-label>
                    <input matInput placeholder="البحث بالاسم أو رقم الهاتف" [(ngModel)]="searchTerm">
                    <mat-icon matSuffix>search</mat-icon>
                  </mat-form-field>

                  <mat-form-field appearance="outline">
                    <mat-label>المحافظة</mat-label>
                    <mat-select [(ngModel)]="selectedGovernorate">
                      <mat-option value="">الكل</mat-option>
                      <mat-option *ngFor="let gov of governorates" [value]="gov.code">
                        {{ gov.name_ar }}
                      </mat-option>
                    </mat-select>
                  </mat-form-field>

                  <mat-form-field appearance="outline">
                    <mat-label>نوع العميل</mat-label>
                    <mat-select [(ngModel)]="selectedType">
                      <mat-option value="">الكل</mat-option>
                      <mat-option *ngFor="let type of customerTypes" [value]="type.id">
                        {{ type.name_ar }}
                      </mat-option>
                    </mat-select>
                  </mat-form-field>

                  <mat-form-field appearance="outline">
                    <mat-label>الفرع</mat-label>
                    <mat-select [(ngModel)]="selectedBranch">
                      <mat-option value="">الكل</mat-option>
                      <mat-option *ngFor="let branch of branches" [value]="branch.id">
                        {{ branch.name_ar }}
                      </mat-option>
                    </mat-select>
                  </mat-form-field>
                </div>
              </mat-card-content>
            </mat-card>

            <mat-card class="table-card">
              <mat-card-content>
                <div class="table-container">
                  <table mat-table [dataSource]="customers" matSort class="customers-table">
                    <ng-container matColumnDef="customer_code">
                      <th mat-header-cell *matHeaderCellDef mat-sort-header>كود العميل</th>
                      <td mat-cell *matCellDef="let customer">{{ customer.customer_code }}</td>
                    </ng-container>

                    <ng-container matColumnDef="name">
                      <th mat-header-cell *matHeaderCellDef mat-sort-header>الاسم</th>
                      <td mat-cell *matCellDef="let customer">{{ customer.name }}</td>
                    </ng-container>

                    <ng-container matColumnDef="phone_primary">
                      <th mat-header-cell *matHeaderCellDef mat-sort-header>رقم الهاتف</th>
                      <td mat-cell *matCellDef="let customer">{{ customer.phone_primary }}</td>
                    </ng-container>

                    <ng-container matColumnDef="governorate">
                      <th mat-header-cell *matHeaderCellDef mat-sort-header>المحافظة</th>
                      <td mat-cell *matCellDef="let customer">{{ customer.governorate_name }}</td>
                    </ng-container>

                    <ng-container matColumnDef="district">
                      <th mat-header-cell *matHeaderCellDef mat-sort-header>المنطقة</th>
                      <td mat-cell *matCellDef="let customer">{{ customer.district_name }}</td>
                    </ng-container>

                    <ng-container matColumnDef="type">
                      <th mat-header-cell *matHeaderCellDef mat-sort-header>النوع</th>
                      <td mat-cell *matCellDef="let customer">
                        <mat-chip *ngIf="customer.customer_type" color="primary" selected>
                          {{ customer.customer_type.name_ar }}
                        </mat-chip>
                      </td>
                    </ng-container>

                    <ng-container matColumnDef="actions">
                      <th mat-header-cell *matHeaderCellDef>الإجراءات</th>
                      <td mat-cell *matCellDef="let customer">
                        <button mat-icon-button [matMenuTriggerFor]="actionMenu">
                          <mat-icon>more_vert</mat-icon>
                        </button>
                        <mat-menu #actionMenu="matMenu">
                          <button mat-menu-item (click)="viewCustomer(customer)">
                            <mat-icon>visibility</mat-icon>
                            <span>عرض</span>
                          </button>
                          <button mat-menu-item (click)="editCustomer(customer)">
                            <mat-icon>edit</mat-icon>
                            <span>تعديل</span>
                          </button>
                          <button mat-menu-item (click)="deleteCustomer(customer)">
                            <mat-icon>delete</mat-icon>
                            <span>حذف</span>
                          </button>
                        </mat-menu>
                      </td>
                    </ng-container>

                    <tr mat-header-row *matHeaderRowDef="displayedColumns"></tr>
                    <tr mat-row *matRowDef="let row; columns: displayedColumns;"></tr>
                  </table>

                  <mat-paginator [pageSizeOptions]="[10, 25, 50, 100]" showFirstLastButtons></mat-paginator>
                </div>
              </mat-card-content>
            </mat-card>
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

    .customers-container {
      max-width: 1400px;
      margin: 0 auto;
    }

    .header-actions {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: 20px;
      flex-wrap: wrap;
      gap: 15px;
    }

    .header-actions h1 {
      color: #333;
      margin: 0;
    }

    .actions {
      display: flex;
      gap: 10px;
      flex-wrap: wrap;
    }

    .filters-card {
      margin-bottom: 20px;
    }

    .filters-grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
      gap: 15px;
    }

    .table-card {
      margin-bottom: 20px;
    }

    .table-container {
      overflow-x: auto;
    }

    .customers-table {
      width: 100%;
    }

    .active {
      background-color: rgba(25, 118, 210, 0.1);
      color: #1976d2;
    }

    @media (max-width: 768px) {
      .header-actions {
        flex-direction: column;
        align-items: stretch;
      }
      
      .actions {
        justify-content: center;
      }
      
      .filters-grid {
        grid-template-columns: 1fr;
      }
    }
  `]
})
export class CustomersComponent implements OnInit {
  isHandset$: Observable<boolean> = new Observable<boolean>(observer => observer.next(false)); // Simplified for now
  searchTerm = '';
  selectedGovernorate = '';
  selectedType = '';
  selectedBranch = '';
  
  customers: any[] = [];
  governorates: any[] = [];
  customerTypes: any[] = [];
  branches: any[] = [];
  
  displayedColumns: string[] = [
    'customer_code', 'name', 'phone_primary', 'governorate', 'district', 'type', 'actions'
  ];

  constructor(
    private authService: AuthService,
    private router: Router
  ) {}

  ngOnInit(): void {
    this.loadCustomers();
    this.loadGovernorates();
    this.loadCustomerTypes();
    this.loadBranches();
  }

  loadCustomers(): void {
    // Mock data for now
    this.customers = [
      {
        id: 1,
        customer_code: 'MAIN-0001',
        name: 'أحمد محمد',
        phone_primary: '+201234567890',
        governorate_name: 'القاهرة',
        district_name: 'المعادي',
        customer_type: { name_ar: 'فرد' }
      },
      {
        id: 2,
        customer_code: 'MAIN-0002',
        name: 'شركة التقنية المتقدمة',
        phone_primary: '+201234567891',
        governorate_name: 'الجيزة',
        district_name: 'الدقي',
        customer_type: { name_ar: 'شركة' }
      }
    ];
  }

  loadGovernorates(): void {
    // Mock data
    this.governorates = [
      { code: 'C', name_ar: 'القاهرة' },
      { code: 'G', name_ar: 'الجيزة' },
      { code: 'ALX', name_ar: 'الإسكندرية' }
    ];
  }

  loadCustomerTypes(): void {
    // Mock data
    this.customerTypes = [
      { id: 1, name_ar: 'فرد' },
      { id: 2, name_ar: 'شركة' },
      { id: 3, name_ar: 'جهة حكومية' }
    ];
  }

  loadBranches(): void {
    // Mock data
    this.branches = [
      { id: 1, name_ar: 'الفرع الرئيسي' },
      { id: 2, name_ar: 'فرع القاهرة' },
      { id: 3, name_ar: 'فرع الجيزة' }
    ];
  }

  addCustomer(): void {
    // TODO: Implement add customer dialog
    console.log('Add customer');
  }

  editCustomer(customer: any): void {
    // TODO: Implement edit customer dialog
    console.log('Edit customer', customer);
  }

  viewCustomer(customer: any): void {
    // TODO: Implement view customer dialog
    console.log('View customer', customer);
  }

  deleteCustomer(customer: any): void {
    // TODO: Implement delete customer confirmation
    console.log('Delete customer', customer);
  }

  exportCustomers(): void {
    // TODO: Implement export functionality
    console.log('Export customers');
  }

  importCustomers(): void {
    // TODO: Implement import functionality
    console.log('Import customers');
  }

  logout(): void {
    this.authService.logout();
    this.router.navigate(['/login']);
  }
} 
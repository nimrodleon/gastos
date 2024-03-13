import { Component } from '@angular/core';
import { DashboardContainerComponent } from '../../../common/containers/dashboard-container/dashboard-container.component';
import { CalendarModule } from 'primeng/calendar';
import { FormsModule } from '@angular/forms';
import { InputTextModule } from 'primeng/inputtext';
import { ButtonModule } from 'primeng/button';
import { IconFieldModule } from 'primeng/iconfield';
import { InputIconModule } from 'primeng/inputicon';
import { TableModule } from 'primeng/table';
import { DropdownModule } from 'primeng/dropdown';
import { Product } from '../../../expenses/interfaces';
import { CategoryFormModalComponent } from '../../components/category-form-modal/category-form-modal.component';

@Component({
  selector: 'app-category-list',
  standalone: true,
  imports: [
    DashboardContainerComponent,
    FormsModule,
    CalendarModule,
    InputTextModule,
    ButtonModule,
    IconFieldModule,
    InputIconModule,
    TableModule,
    DropdownModule,
    CategoryFormModalComponent
  ],
  templateUrl: './category-list.component.html',
  styleUrl: './category-list.component.scss'
})
export class CategoryListComponent {
  rangeDates: Date[] | undefined;
  query: string = "";
  products!: Product[];
  visibleModal: boolean = false;

  public addExpense(): void {
    this.visibleModal = true;
  }

  public changeVisibleModal(value: boolean): void {
    this.visibleModal = value;
  }
}

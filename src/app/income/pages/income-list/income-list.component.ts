import { Component } from '@angular/core';
import { DashboardContainerComponent } from '../../../common/containers/dashboard-container/dashboard-container.component';
import { CalendarModule } from 'primeng/calendar';
import { FormsModule } from '@angular/forms';
import { InputTextModule } from 'primeng/inputtext';
import { ButtonModule } from 'primeng/button';
import { IconFieldModule } from 'primeng/iconfield';
import { InputIconModule } from 'primeng/inputicon';
import { TableModule } from 'primeng/table';
import { Product } from '../../../expenses/interfaces';
import { IncomeFormModalComponent } from '../../components/income-form-modal/income-form-modal.component';

@Component({
  selector: 'app-income-list',
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
    IncomeFormModalComponent
  ],
  templateUrl: './income-list.component.html',
  styleUrl: './income-list.component.scss'
})
export class IncomeListComponent {
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

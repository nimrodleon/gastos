import { Component } from '@angular/core';
import { DashboardContainerComponent } from '../../../common/containers/dashboard-container/dashboard-container.component';
import { CalendarModule } from 'primeng/calendar';
import { FormsModule } from '@angular/forms';
import { InputTextModule } from 'primeng/inputtext';
import { ButtonModule } from 'primeng/button';
import { IconFieldModule } from 'primeng/iconfield';
import { InputIconModule } from 'primeng/inputicon';
import { TableModule } from 'primeng/table';
import { Product } from '../../interfaces';

@Component({
  selector: 'app-expenses-list',
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
  ],
  templateUrl: './expenses-list.component.html',
  styleUrl: './expenses-list.component.scss'
})
export class ExpensesListComponent {
  rangeDates: Date[] | undefined;
  query: string = "";
  products!: Product[];

}

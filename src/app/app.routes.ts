import { Routes } from '@angular/router';
import { BalanceReportComponent } from './balance/pages/balance-report/balance-report.component';
import { ExpensesListComponent } from './expenses/pages/expenses-list/expenses-list.component';
import { IncomeListComponent } from './income/pages/income-list/income-list.component';

export const routes: Routes = [
  { path: "", component: BalanceReportComponent },
  { path: "expenses", component: ExpensesListComponent },
  {path: "income-list", component: IncomeListComponent}
];

import { Component } from '@angular/core';
import { DashboardContainerComponent } from '../../../common/containers/dashboard-container/dashboard-container.component';

@Component({
  selector: 'app-balance-report',
  standalone: true,
  imports: [
    DashboardContainerComponent
  ],
  templateUrl: './balance-report.component.html',
  styleUrl: './balance-report.component.scss'
})
export class BalanceReportComponent {

}

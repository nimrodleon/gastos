import { Component, OnInit } from '@angular/core';
import { CardModule } from 'primeng/card';
import { ChartModule } from 'primeng/chart';
import { DashboardContainerComponent } from '../../../common/containers/dashboard-container/dashboard-container.component';
import { FormsModule } from '@angular/forms';
import { DropdownModule } from 'primeng/dropdown';

@Component({
  selector: 'app-balance-report',
  standalone: true,
  imports: [
    DashboardContainerComponent,
    CardModule,
    ChartModule,
    FormsModule,
    DropdownModule,
  ],
  templateUrl: './balance-report.component.html',
  styleUrl: './balance-report.component.scss'
})
export class BalanceReportComponent implements OnInit {
  data: any;
  options: any;
  years: { label: string, value: number }[] = [];
  selectedYear: number = 2024;

  ngOnInit() {
    // Initialize years array
    this.years = [
      { label: 'TOTAL', value: 0 },
      { label: '2024', value: 2024 },
      { label: '2023', value: 2023 },
      // Add more years as needed
    ];

    // Set initial selected year
    this.selectedYear = this.years[0].value;

    const documentStyle = getComputedStyle(document.documentElement);
    const textColor = documentStyle.getPropertyValue('--text-color');
    const textColorSecondary = documentStyle.getPropertyValue('--text-color-secondary');
    const surfaceBorder = documentStyle.getPropertyValue('--surface-border');

    this.data = {
      labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
      datasets: [
        {
          label: 'My First dataset',
          backgroundColor: documentStyle.getPropertyValue('--blue-500'),
          borderColor: documentStyle.getPropertyValue('--blue-500'),
          data: [65, 59, 80, 81, 56, 55, 40]
        },
        {
          label: 'My Second dataset',
          backgroundColor: documentStyle.getPropertyValue('--pink-500'),
          borderColor: documentStyle.getPropertyValue('--pink-500'),
          data: [28, 48, 40, 19, 86, 27, 90]
        }
      ]
    };

    this.options = {
      maintainAspectRatio: false,
      aspectRatio: 0.8,
      plugins: {
        legend: {
          labels: {
            color: textColor
          }
        }
      },
      scales: {
        x: {
          ticks: {
            color: textColorSecondary,
            font: {
              weight: 500
            }
          },
          grid: {
            color: surfaceBorder,
            drawBorder: false
          }
        },
        y: {
          ticks: {
            color: textColorSecondary
          },
          grid: {
            color: surfaceBorder,
            drawBorder: false
          }
        }

      }
    };
  }

  public onYearChange(event: any) {
    this.selectedYear = event.value;
  }

}

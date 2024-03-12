import { Component } from '@angular/core';
import { HeaderComponent } from '../../components/header/header.component';

@Component({
  selector: 'app-dashboard-container',
  standalone: true,
  imports: [
    HeaderComponent
  ],
  templateUrl: './dashboard-container.component.html'
})
export class DashboardContainerComponent {

}

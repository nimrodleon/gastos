import { Component, EventEmitter, Input, Output } from '@angular/core';
import { DialogModule } from 'primeng/dialog';
import { InputTextModule } from 'primeng/inputtext';
import { InputNumberModule } from 'primeng/inputnumber';
import { InputTextareaModule } from 'primeng/inputtextarea';
import { AutoCompleteModule } from 'primeng/autocomplete';
import { ButtonModule } from 'primeng/button';

@Component({
  selector: 'app-expense-form-modal',
  standalone: true,
  imports: [
    DialogModule,
    InputTextModule,
    InputNumberModule,
    InputTextareaModule,
    AutoCompleteModule,
    ButtonModule,
  ],
  templateUrl: './expense-form-modal.component.html',
  styleUrl: './expense-form-modal.component.scss'
})
export class ExpenseFormModalComponent {
  @Input()
  visible: boolean = false;
  @Output()
  visibilityChanged = new EventEmitter<boolean>();

  public hideModal(): void {
    this.visibilityChanged.emit(false);
  }

}

import { Component, EventEmitter, Input, Output } from '@angular/core';
import { DialogModule } from 'primeng/dialog';
import { InputTextModule } from 'primeng/inputtext';
import { InputTextareaModule } from 'primeng/inputtextarea';
import { ButtonModule } from 'primeng/button';
import { DropdownModule } from 'primeng/dropdown';

@Component({
  selector: 'app-category-form-modal',
  standalone: true,
  imports: [
    DialogModule,
    InputTextModule,
    InputTextareaModule,
    DropdownModule,
    ButtonModule,
  ],
  templateUrl: './category-form-modal.component.html',
  styleUrl: './category-form-modal.component.scss'
})
export class CategoryFormModalComponent {
  @Input()
  visible: boolean = false;
  @Output()
  visibilityChanged = new EventEmitter<boolean>();

  public hideModal(): void {
    this.visibilityChanged.emit(false);
  }
}

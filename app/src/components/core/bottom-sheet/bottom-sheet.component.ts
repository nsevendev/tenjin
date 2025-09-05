import {Component, HostListener, OnDestroy, OnInit, computed, input, model, output} from '@angular/core';
import {BottomSheetService} from './bottom-sheet.service';

@Component({
  selector: 'c-bottom-sheet',
  standalone: true,
  templateUrl: './bottom-sheet.template.html',
})
export class BottomSheetComponent implements OnInit, OnDestroy {
  // state
  readonly open = model<boolean>(false);

  // config
  readonly height = input<string>('85vh');
  readonly closeOnBackdrop = input<boolean>(true);
  readonly showClose = input<boolean>(true);
  readonly title = input<string>('');
  readonly className = input<string>('', { alias: 'class' });
  readonly id = input<string>('');
  readonly showFooter = input<boolean>(false);
  readonly confirmText = input<string>('Valider');
  readonly cancelText = input<string>('Annuler');
  readonly closeOnConfirm = input<boolean>(true);
  readonly closeOnCancel = input<boolean>(true);

  readonly confirmed = output<void>();
  readonly cancelled = output<void>();

  readonly sheetMaxHeight = computed(() => this.height());
  readonly contentMaxHeight = computed(() => `calc(${this.height()} - 56px)`);

  close() {
    this.open.set(false);
  }

  onBackdropClick() {
    if (this.closeOnBackdrop()) this.close();
  }

  @HostListener('document:keydown', ['$event'])
  onKeydown(e: KeyboardEvent) {
    if (!this.open()) return;
    if (e.key === 'Escape') this.close();
  }

  constructor(private readonly sheetSvc: BottomSheetService) {}

  ngOnInit() {
    const key = this.id();
    if (key) {
      this.sheetSvc.register(key, () => this.open(), v => this.open.set(v));
    }
  }

  ngOnDestroy() {
    const key = this.id();
    if (key) this.sheetSvc.unregister(key);
  }

  onConfirm() {
    this.confirmed.emit();
    if (this.closeOnConfirm()) this.close();
  }

  onCancel() {
    this.cancelled.emit();
    if (this.closeOnCancel()) this.close();
  }
}

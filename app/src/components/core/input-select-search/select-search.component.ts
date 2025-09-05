import {Component, ElementRef, HostListener, ViewChild, ViewChildren, QueryList, computed, input, model, signal} from '@angular/core';
import {FormsModule} from '@angular/forms';
import {cn} from '../../../utils/class-name/cn';
import {SelectOption, SelectOptionValue, SelectVariantProps, selectVariants} from '../input-select/select.variant';
import {HericonComponent} from '../icon-svg/hericon/hericon.component';
import {BottomSheetComponent} from '../bottom-sheet/bottom-sheet.component';

type VV<K extends keyof SelectVariantProps> = NonNullable<SelectVariantProps[K]>;

@Component({
  selector: 'c-input-select-search',
  standalone: true,
  imports: [FormsModule, HericonComponent, BottomSheetComponent],
  templateUrl: './select-search.template.html',
})
export class InputSelectSearchComponent {
  @ViewChild('triggerEl') triggerEl?: ElementRef<HTMLElement>;
  @ViewChild('searchInputRef') searchInput?: ElementRef<HTMLInputElement>;
  @ViewChildren('optEl') optEls?: QueryList<ElementRef<HTMLElement>>;
  // Variants
  readonly size    = input<VV<'size'>>('base');
  readonly variant = input<VV<'variant'>>('default');
  readonly radius  = input<VV<'radius'>>('md');

  // Props
  readonly options = input<SelectOption[]>([]);
  readonly placeholder = input<string>('Rechercher...');
  readonly disabled  = input<boolean>(false);
  readonly className = input<string>('', { alias: 'class' });
  readonly panelMode = input<'auto' | 'dropdown' | 'sheet'>('auto');

  // Value model
  readonly value = model<SelectOptionValue | null>(null);

  // Local state
  readonly open = signal(false);
  readonly query = signal('');
  readonly activeIndex = signal<number>(-1);
  private isMobileViewport(): boolean {
    return typeof window !== 'undefined' && window.matchMedia('(max-width: 768px)').matches;
  }
  readonly useSheet = computed(() =>
    this.panelMode() === 'sheet' || (this.panelMode() === 'auto' && this.isMobileViewport())
  );

  private readonly padRight = computed(() => {
    const s = this.size();
    return s === 'sm' ? 'pr-8' : s === 'lg' ? 'pr-12' : 'pr-10';
  });

  readonly triggerClasses = computed(() =>
    cn(
      selectVariants({ size: this.size(), variant: this.variant(), radius: this.radius() }),
      'min-h-11 md:min-h-[2.5rem] flex items-center justify-between',
      this.padRight(),
      this.className(),
    )
  );

  selectedLabel = computed(() => {
    const v = this.value();
    if (v === null || v === '') return '';
    const o = this.options().find(o => o.value === v);
    return o?.label ?? '';
  });

  filtered = computed(() => {
    const q = this.query().toLowerCase();
    if (!q) return this.options();
    return this.options().filter(o => o.label.toLowerCase().includes(q));
  });

  setValue(v: SelectOptionValue) {
    this.value.set(v);
    this.open.set(false);
    this.query.set('');
  }

  clear(e: MouseEvent) {
    e.stopPropagation();
    this.value.set(null);
  }

  toggleOpen() {
    if (this.disabled()) return;
    const next = !this.open();
    this.open.set(next);
    if (next) {
      const len = this.filtered().length;
      this.activeIndex.set(len > 0 ? 0 : -1);
      setTimeout(() => this.searchInput?.nativeElement?.focus(), 0);
      setTimeout(() => this.scrollToActive(), 0);
    }
  }

  @HostListener('document:click', ['$event'])
  onDocumentClick(ev: MouseEvent) {
    if (!this.open()) return;
    if (!this.triggerEl?.nativeElement) return;
    const host = (this.triggerEl.nativeElement.parentElement as HTMLElement) ?? this.triggerEl.nativeElement;
    if (!host.contains(ev.target as Node)) {
      this.open.set(false);
    }
  }

  onKeydown(e: KeyboardEvent) {
    const opts = this.filtered();
    const max = opts.length - 1;
    if (!this.open()) {
      if (e.key === 'ArrowDown' || e.key === 'Enter' || e.key === ' ') {
        e.preventDefault();
        this.toggleOpen();
      }
      return;
    }

    if (e.key === 'Escape' || e.key === 'Tab') {
      this.open.set(false);
      return;
    }

    const cur = this.activeIndex();
    if (e.key === 'ArrowDown') {
      e.preventDefault();
      const next = cur < max ? cur + 1 : 0;
      this.activeIndex.set(next);
      this.scrollToActive();
    } else if (e.key === 'ArrowUp') {
      e.preventDefault();
      const next = cur > 0 ? cur - 1 : max;
      this.activeIndex.set(next);
      this.scrollToActive();
    } else if (e.key === 'Enter') {
      e.preventDefault();
      const idx = cur >= 0 && cur <= max ? cur : 0;
      const o = opts[idx];
      if (o) this.setValue(o.value);
    }
  }

  onPanelKeydown(e: KeyboardEvent) {
    // Delegate to same navigation handler
    this.onKeydown(e);
  }

  private scrollToActive() {
    const idx = this.activeIndex();
    const el = this.optEls?.get(idx)?.nativeElement;
    if (el) {
      el.scrollIntoView({ block: 'nearest' });
    }
  }
}

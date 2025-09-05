import {Component, ElementRef, HostListener, ViewChild, ViewChildren, QueryList, computed, input, model, signal} from '@angular/core';
import {FormsModule} from '@angular/forms';
import {cn} from '../../../utils/class-name/cn';
import {SelectOption, SelectOptionValue, SelectVariantProps, selectVariants} from '../input-select/select.variant';
import {InputCheckboxComponent} from '../input-checkbox/input-checkbox.component';
import {HericonComponent} from '../icon-svg/hericon/hericon.component';
import {BottomSheetComponent} from '../bottom-sheet/bottom-sheet.component';

type VV<K extends keyof SelectVariantProps> = NonNullable<SelectVariantProps[K]>;

@Component({
  selector: 'c-input-select-multi',
  standalone: true,
  imports: [FormsModule, InputCheckboxComponent, HericonComponent, BottomSheetComponent],
  templateUrl: './select-multi.template.html',
})
export class InputSelectMultiComponent {
  @ViewChild('triggerEl') triggerEl?: ElementRef<HTMLElement>;
  @ViewChildren('optEl') optEls?: QueryList<ElementRef<HTMLElement>>;
  constructor(private host: ElementRef<HTMLElement>) {}
  // Variants
  readonly size    = input<VV<'size'>>('base');
  readonly variant = input<VV<'variant'>>('default');
  readonly radius  = input<VV<'radius'>>('md');

  // Props
  readonly options = input<SelectOption[]>([]);
  readonly placeholder = input<string>('Sélectionner...');
  readonly disabled  = input<boolean>(false);
  readonly className = input<string>('', { alias: 'class' });
  readonly panelMode = input<'auto' | 'dropdown' | 'sheet'>('auto');

  // Value model
  readonly values = model<SelectOptionValue[]>([]);

  // Local state
  readonly open = signal(false);
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
      'min-h-11 md:min-h-[2.5rem] flex items-center',
      this.padRight(),
      this.className(),
    )
  );

  selectedOptions = computed(() => {
    const set = new Set(this.values());
    return this.options().filter(o => set.has(o.value));
  });

  isSelected = (v: SelectOptionValue) => this.values().includes(v);

  toggleOpen() {
    if (this.disabled()) return;
    const next = !this.open();
    this.open.set(next);
    if (next) {
      const len = this.options().length;
      this.activeIndex.set(len > 0 ? 0 : -1);
      setTimeout(() => this.scrollToActive(), 0);
    }
  }

  removeValue(v: SelectOptionValue) {
    const cur = this.values();
    if (cur.includes(v)) this.values.set(cur.filter(x => x !== v));
  }

  onToggle(v: SelectOptionValue, next: boolean) {
    const cur = this.values();
    if (next) {
      if (!cur.includes(v)) this.values.set([...cur, v]);
    } else {
      if (cur.includes(v)) this.values.set(cur.filter(x => x !== v));
    }
  }

  @HostListener('document:click', ['$event'])
  onDocumentClick(ev: MouseEvent) {
    if (!this.open()) return;
    if (!this.host.nativeElement.contains(ev.target as Node)) {
      this.open.set(false);
    }
  }

  onKeydown(e: KeyboardEvent) {
    const opts = this.options();
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
    } else if (e.key === 'ArrowUp') {
      e.preventDefault();
      const next = cur > 0 ? cur - 1 : max;
      this.activeIndex.set(next);
      this.scrollToActive();
    } else if (e.key === 'Enter' || e.key === ' ') {
      e.preventDefault();
      const idx = cur >= 0 && cur <= max ? cur : 0;
      const o = opts[idx];
      if (o) this.onToggle(o.value, !this.isSelected(o.value));
    }
  }

  private scrollToActive() {
    const idx = this.activeIndex();
    const el = this.optEls?.get(idx)?.nativeElement;
    if (el) {
      el.scrollIntoView({ block: 'nearest' });
    }
  }
}

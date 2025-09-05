import {Component, computed, input, model} from '@angular/core';
import {FormsModule} from '@angular/forms';
import {cn} from '../../../utils/class-name/cn';
import {SelectOption, SelectOptionValue, SelectVariantProps, selectVariants} from './select.variant';

type VV<K extends keyof SelectVariantProps> = NonNullable<SelectVariantProps[K]>;

@Component({
  selector: 'c-input-select',
  standalone: true,
  imports: [FormsModule],
  templateUrl: './select.template.html',
})
export class InputSelectComponent {
  // Variants
  readonly size    = input<VV<'size'>>('base');
  readonly variant = input<VV<'variant'>>('default');
  readonly radius  = input<VV<'radius'>>('md');

  // Props
  readonly options = input<SelectOption[]>([]);
  readonly placeholder = input<string>('');
  readonly disabled  = input<boolean>(false);
  readonly className = input<string>('', { alias: 'class' });

  // Value model
  readonly value = model<SelectOptionValue | null>(null);

  readonly invalid = computed(() => this.variant() === 'error');

  // Padding right to leave space for chevron depending on size
  private readonly padRight = computed(() => {
    const s = this.size();
    return s === 'sm' ? 'pr-8' : s === 'lg' ? 'pr-12' : 'pr-10';
  });

  readonly classes = computed(() =>
    cn(
      selectVariants({ size: this.size(), variant: this.variant(), radius: this.radius() }),
      'min-h-11 md:min-h-[2.5rem]',
      this.padRight(),
      this.className(),
    )
  );
}

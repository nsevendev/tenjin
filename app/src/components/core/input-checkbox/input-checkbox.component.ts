import {Component, computed, input, model} from '@angular/core';
import {FormsModule} from '@angular/forms';
import {cn} from '../../../utils/class-name/cn';
import {InputCheckboxVariantProps, inputCheckboxVariants} from './input-checkbox.variant';

type VV<K extends keyof InputCheckboxVariantProps> = NonNullable<InputCheckboxVariantProps[K]>;

@Component({
  selector: 'c-input-checkbox',
  standalone: true,
  templateUrl: './input-checkbox.template.html',
  imports: [FormsModule]
})
export class InputCheckboxComponent {
  // Variants
  readonly size   = input<VV<'size'>>('base');
  readonly variant = input<VV<'variant'>>('default');
  readonly radius = input<VV<'radius'>>('md');

  // Props
  readonly disabled  = input<boolean>(false);
  readonly className = input<string>('', { alias: 'class' });
  readonly checked   = model<boolean>(false);

  readonly invalid = computed(() => this.variant() === 'error');

  readonly classes = computed(() =>
    cn(
      inputCheckboxVariants({
        size: this.size(),
        variant: this.variant(),
        radius: this.radius(),
      }),
      this.className(),
    )
  );
}


import {Component, computed, ContentChild, input, model} from '@angular/core';
import {FormsModule} from '@angular/forms';
import {cn} from '../../../utils/class-name/cn';
import {InputType, InputVariantProps, inputVariants} from './input.variant';
import {CInputPrefixDirective} from './input-prefix.directive';
import {CInputSuffixDirective} from './input-suffix.directive';

type VV<K extends keyof InputVariantProps> = NonNullable<InputVariantProps[K]>;

@Component({
  selector: 'c-input',
  standalone: true,
  templateUrl: './input.template.html',
  imports: [
    FormsModule
  ]
})
export class InputComponent {
  // Variants pour style
  readonly size = input<VV<'size'>>('base');
  readonly variant = input<VV<'variant'>>('default');
  readonly radius = input<VV<'radius'>>('md');

  // Props classiques
  readonly type = input<InputType>('text');
  readonly placeholder = input<string>('');
  readonly disabled = input<boolean>(false);
  readonly className = input<string>('', { alias: 'class' });
  readonly value = model<string>('');

  @ContentChild(CInputPrefixDirective) prefix?: CInputPrefixDirective;
  @ContentChild(CInputSuffixDirective) suffix?: CInputSuffixDirective;

  readonly invalid = computed(() => this.variant() === 'error');

  hasPrefix = computed(() => !!this.prefix);
  hasSuffix = computed(() => !!this.suffix);

  private readonly padLeft = computed(() => {
    const s = this.size();
    return s === 'sm' ? 'pl-8' : s === 'lg' ? 'pl-12' : 'pl-10';
  });

  private readonly padRight = computed(() => {
    const s = this.size();
    return s === 'sm' ? 'pr-8' : s === 'lg' ? 'pr-12' : 'pr-10';
  });

  prefixWrapperPad = computed(() => (this.size() === 'sm' ? 'pl-2' : this.size() === 'lg' ? 'pl-3' : 'pl-3'));
  suffixWrapperPad = computed(() => (this.size() === 'sm' ? 'pr-2' : this.size() === 'lg' ? 'pr-3' : 'pr-3'));

  readonly classes = computed(() =>
    cn(
      inputVariants({ size: this.size(), variant: this.variant(), radius: this.radius() }),
      this.className(),
      this.hasPrefix() ? this.padLeft() : '',
      this.hasSuffix() ? this.padRight() : '',
    )
  );
}

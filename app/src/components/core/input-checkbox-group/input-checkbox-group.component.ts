import {Component, computed, input, model} from '@angular/core';
import {cn} from '../../../utils/class-name/cn';
import {InputCheckboxComponent} from '../input-checkbox/input-checkbox.component';

type OptionValue = string | number;
export type CheckboxOption = {
  label: string;
  value: OptionValue;
  disabled?: boolean;
};

@Component({
  selector: 'c-input-checkbox-group',
  standalone: true,
  imports: [InputCheckboxComponent],
  templateUrl: './input-checkbox-group.template.html',
})
export class InputCheckboxGroupComponent {
  // Group props
  readonly options = input<CheckboxOption[]>([]);
  readonly disabled = input<boolean>(false);
  readonly className = input<string>('', { alias: 'class' });

  // Variant passthrough to children
  readonly size   = input<'sm' | 'base' | 'lg'>('base');
  readonly variant = input<'default' | 'error' | 'success' | 'warning'>('default');
  readonly radius = input<'none' | 'sm' | 'md' | 'lg' | 'full'>('md');

  // Layout controls
  readonly display = input<'flex' | 'inline-flex'>('inline-flex');
  readonly direction = input<'row' | 'col'>('row');
  readonly wrap = input<'nowrap' | 'wrap' | 'wrap-reverse'>('wrap');
  readonly gap = input<0|1|2|3|4|5|6|8|10|12|16|20>(3);

  // Selection model (multi)
  readonly checkedValues = model<OptionValue[]>([]);

  readonly wrapperClasses = computed(() =>
    cn(
      this.display(),
      this.direction() === 'row' ? 'flex-row' : 'flex-col',
      this.wrap() === 'wrap' ? 'flex-wrap' : this.wrap() === 'wrap-reverse' ? 'flex-wrap-reverse' : 'flex-nowrap',
      this.gap() === 0 ? 'gap-0' : `gap-${this.gap()}`,
      this.direction() === 'row' ? 'items-center' : 'items-start',
      this.className()
    )
  );

  isSelected = (val: OptionValue) => this.checkedValues().includes(val);

  onToggle(val: OptionValue, next: boolean) {
    const cur = this.checkedValues();
    if (next) {
      if (!cur.includes(val)) this.checkedValues.set([...cur, val]);
    } else {
      if (cur.includes(val)) this.checkedValues.set(cur.filter(v => v !== val));
    }
  }
}

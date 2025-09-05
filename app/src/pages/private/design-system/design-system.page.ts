import {Component, OnDestroy, OnInit, signal} from '@angular/core';
import {LinkTextComponent} from '../../../components/core/link-text/link-text.component';
import {CardComponent} from '../../../components/core/card/card.component';
import {ContainerComponent} from '../../../components/core/container/container.component';
import {TitleComponent} from '../../../components/core/title/title.component';
import {TextBlockquoteComponent, TextComponent, TextMessageComponent} from '../../../components/core/text/text.component';
import {ErrorTextMessageComponent, InfoTextMessageComponent, SuccessTextMessageComponent, WarningTextMessageComponent} from '../../../components/core/text/message.preset';
import {CaptionTextComponent, CodeTextComponent, LeadTextComponent, MutedTextComponent, ParagraphTextComponent} from '../../../components/core/text/text.preset';
import {PrimaryButtonComponent, PrimarySmallButtonComponent, SecondaryButtonComponent, SecondarySmallButtonComponent} from '../../../components/core/button/button.preset';
import {ButtonComponent} from '../../../components/core/button/button.component';
import {BadgeComponent} from '../../../components/core/badge/badge.component';
import {InputComponent} from '../../../components/core/input/input.component';
import {CInputPrefixDirective} from '../../../components/core/input/input-prefix.directive';
import {CInputSuffixDirective} from '../../../components/core/input/input-suffix.directive';
import {HericonComponent} from '../../../components/core/icon-svg/hericon/hericon.component';
import {InputCheckboxComponent} from '../../../components/core/input-checkbox/input-checkbox.component';
import {InputCheckboxGroupComponent} from '../../../components/core/input-checkbox-group/input-checkbox-group.component';
import {InputSelectComponent} from '../../../components/core/input-select/select.component';
import {InputSelectMultiComponent} from '../../../components/core/input-select-multi/select-multi.component';
import {InputSelectSearchComponent} from '../../../components/core/input-select-search/select-search.component';
import {BottomSheetComponent} from '../../../components/core/bottom-sheet/bottom-sheet.component';
import {BottomSheetService} from '../../../components/core/bottom-sheet/bottom-sheet.service';

@Component({
  selector: 'page-design-system',
  templateUrl: './design-system.page.html',
  imports: [
    LinkTextComponent,
    CardComponent,
    ContainerComponent,
    TitleComponent,
    TextBlockquoteComponent,
    ErrorTextMessageComponent,
    WarningTextMessageComponent,
    SuccessTextMessageComponent,
    TextMessageComponent,
    CodeTextComponent,
    MutedTextComponent,
    LeadTextComponent,
    CaptionTextComponent,
    ParagraphTextComponent,
    TextComponent,
    SecondarySmallButtonComponent,
    SecondaryButtonComponent,
    PrimarySmallButtonComponent,
    PrimaryButtonComponent,
    ButtonComponent,
    BadgeComponent,
    InputComponent,
    CInputPrefixDirective,
    CInputSuffixDirective,
    InfoTextMessageComponent,
    HericonComponent,
    InputCheckboxComponent,
    InputCheckboxGroupComponent,
    InputSelectComponent,
    InputSelectMultiComponent,
    InputSelectSearchComponent,
    BottomSheetComponent
  ]
})
export class DesignSystemPage implements OnInit, OnDestroy {
  valueInput = signal('');
  username = signal('');
  filter = signal('');
  valueInputWithValue = signal('je suis un input avec une valeur par defaut');
  // Checkbox models
  acceptTerms = signal(false);
  errorChoice = signal(false);
  bigRound = signal(true);

  // Checkbox group demo
  fruits = [
    { label: 'Pomme', value: 'pomme' },
    { label: 'Banane', value: 'banane' },
    { label: 'Cerise', value: 'cerise' },
    { label: 'Orange', value: 'orange' },
  ];
  selectedFruits = signal<string[]>(['pomme']);

  // Select demo
  fruitOptions = this.fruits;
  selectedFruit = signal<string | null>(null);
  selectedFruitError = signal<string | null>(null);
  selectedFruitsForMulti = signal<string[]>(['cerise']);
  // Large options to demo scroll-into-view
  manyOptions = Array.from({length: 50}, (_, i) => ({ label: `Option ${i+1}`, value: `opt${i+1}` }));
  selectedMany = signal<string[]>([]);
  selectedOneMany = signal<string | null>(null);
  openGenericSheet = signal(false);
  // Footer sheet demo
  openFooterSheet = signal(false);
  lastAction = signal('');
  // Programmatic example (desktop inline, mobile sheet)
  isMobile = signal(false);
  openFiltersSheet = signal(false);
  private mql?: MediaQueryList;

  constructor(private readonly sheetSvc: BottomSheetService) {}

  ngOnInit() {
    if (typeof window !== 'undefined') {
      this.mql = window.matchMedia('(max-width: 768px)');
      const update = () => this.isMobile.set(!!this.mql?.matches);
      this.mql.addEventListener?.('change', update);
      // store listener for removal using closure
      // @ts-ignore
      this._mqlUpdate = update;
      update();
    }
  }

  // holder for the listener reference
  private _mqlUpdate?: () => void;

  ngOnDestroy() {
    if (this.mql && this._mqlUpdate) {
      this.mql.removeEventListener?.('change', this._mqlUpdate);
    }
  }

  openFiltersFromService() {
    this.sheetSvc.open('filters-sheet');
  }

  handlerBadge() {
    console.log('Design SystemPage BADGE clicked');
  }

  handlerButton() {
    console.log('Design SystemPage BUTTON clicked');
  }
}

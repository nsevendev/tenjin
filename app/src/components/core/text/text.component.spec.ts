import { ComponentFixture, TestBed } from '@angular/core/testing';
import { provideZonelessChangeDetection } from '@angular/core';
import { TextComponent, TextMessageComponent, TextBlockquoteComponent } from './text.component';

describe('TextComponent', () => {
  let component: TextComponent;
  let fixture: ComponentFixture<TextComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TextComponent],
      providers: [provideZonelessChangeDetection()]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(TextComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

describe('TextMessageComponent', () => {
  let component: TextMessageComponent;
  let fixture: ComponentFixture<TextMessageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TextMessageComponent],
      providers: [provideZonelessChangeDetection()]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(TextMessageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

describe('TextBlockquoteComponent', () => {
  let component: TextBlockquoteComponent;
  let fixture: ComponentFixture<TextBlockquoteComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TextBlockquoteComponent],
      providers: [provideZonelessChangeDetection()]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(TextBlockquoteComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { provideZonelessChangeDetection } from '@angular/core';
import { InputComponent } from './input.component';

describe('InputComponent', () => {
  let component: InputComponent;
  let fixture: ComponentFixture<InputComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [InputComponent],
      providers: [provideZonelessChangeDetection()]
    })
    .compileComponents();

    fixture = TestBed.createComponent(InputComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should have default values', () => {
    expect(component.size()).toBe('base');
    expect(component.variant()).toBe('default');
    expect(component.radius()).toBe('md');
    expect(component.type()).toBe('text');
    expect(component.placeholder()).toBe('');
    expect(component.disabled()).toBe(false);
    expect(component.value()).toBe('');
  });

  it('should update value when model changes', () => {
    component.value.set('test value');
    expect(component.value()).toBe('test value');
  });

  it('should compute invalid state correctly', () => {
    expect(component.invalid()).toBe(false);
    
    fixture.componentRef.setInput('variant', 'error');
    fixture.detectChanges();
    
    expect(component.invalid()).toBe(true);
  });

  it('should compute padding classes correctly', () => {
    fixture.componentRef.setInput('size', 'sm');
    fixture.detectChanges();
    expect(component.prefixWrapperPad()).toBe('pl-2');
    expect(component.suffixWrapperPad()).toBe('pr-2');

    fixture.componentRef.setInput('size', 'lg');
    fixture.detectChanges();
    expect(component.prefixWrapperPad()).toBe('pl-3');
    expect(component.suffixWrapperPad()).toBe('pr-3');

    fixture.componentRef.setInput('size', 'base');
    fixture.detectChanges();
    expect(component.prefixWrapperPad()).toBe('pl-3');
    expect(component.suffixWrapperPad()).toBe('pr-3');
  });

  it('should apply custom className', () => {
    fixture.componentRef.setInput('class', 'custom-class');
    fixture.detectChanges();
    
    expect(component.classes()).toContain('custom-class');
  });
});
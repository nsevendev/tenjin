import { ComponentFixture, TestBed } from '@angular/core/testing';
import { provideZonelessChangeDetection } from '@angular/core';
import { BadgeComponent } from './badge.component';

describe('BadgeComponent', () => {
  let component: BadgeComponent;
  let fixture: ComponentFixture<BadgeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [BadgeComponent],
      providers: [provideZonelessChangeDetection()]
    })
    .compileComponents();

    fixture = TestBed.createComponent(BadgeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should have default color gray', () => {
    expect(component.color()).toBe('gray');
  });

  it('should have default size base', () => {
    expect(component.size()).toBe('base');
  });

  it('should have default hasHandler false', () => {
    expect(component.hasHandler()).toBe(false);
  });

  it('should emit pressed event on click', () => {
    spyOn(component.pressed, 'emit');
    const event = new MouseEvent('click');

    component.onClick(event);

    expect(component.pressed.emit).toHaveBeenCalledWith(event);
  });

  it('should apply custom className', () => {
    fixture.componentRef.setInput('class', 'custom-class');
    fixture.detectChanges();

    expect(component.classes()).toContain('custom-class');
  });

  it('should apply color variant classes', () => {
    fixture.componentRef.setInput('color', 'blue');
    fixture.detectChanges();

    expect(component.classes()).toContain('bg-blue-300');
    expect(component.classes()).toContain('text-blue-900');
  });

  it('should apply size variant classes', () => {
    fixture.componentRef.setInput('size', 'sm');
    fixture.detectChanges();

    expect(component.classes()).toContain('px-2.5');
    expect(component.classes()).toContain('py-0.5');
  });

  it('should apply interactive classes when hasHandler is true', () => {
    fixture.componentRef.setInput('hasHandler', true);
    fixture.detectChanges();

    expect(component.classes()).toContain('cursor-pointer');
  });
});

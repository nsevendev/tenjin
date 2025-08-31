import { ComponentFixture, TestBed } from '@angular/core/testing';
import { provideZonelessChangeDetection } from '@angular/core';
import { provideRouter } from '@angular/router';
import { LayoutPrivate } from './layout-private';

describe('LayoutPrivate', () => {
  let component: LayoutPrivate;
  let fixture: ComponentFixture<LayoutPrivate>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [LayoutPrivate],
      providers: [provideZonelessChangeDetection(), provideRouter([])]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(LayoutPrivate);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
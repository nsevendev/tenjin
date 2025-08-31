import { ComponentFixture, TestBed } from '@angular/core/testing';
import { provideZonelessChangeDetection } from '@angular/core';
import { provideRouter } from '@angular/router';
import { DesignSystemPage } from './design-system.page';

describe('DesignSystemPage', () => {
  let component: DesignSystemPage;
  let fixture: ComponentFixture<DesignSystemPage>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [DesignSystemPage],
      providers: [provideZonelessChangeDetection(), provideRouter([])]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(DesignSystemPage);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
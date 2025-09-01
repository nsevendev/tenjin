import { ComponentFixture, TestBed } from '@angular/core/testing';
import { provideZonelessChangeDetection } from '@angular/core';
import { provideRouter } from '@angular/router';
import { LayoutPublic } from './layout-public';

describe('LayoutPublic', () => {
  let component: LayoutPublic;
  let fixture: ComponentFixture<LayoutPublic>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [LayoutPublic],
      providers: [provideZonelessChangeDetection(), provideRouter([])]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(LayoutPublic);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
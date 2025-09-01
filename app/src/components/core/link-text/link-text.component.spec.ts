import { ComponentFixture, TestBed } from '@angular/core/testing';
import { provideZonelessChangeDetection } from '@angular/core';
import { LinkTextComponent } from './link-text.component';
import {provideRouter} from '@angular/router';

describe('LinkTextComponent', () => {
  let component: LinkTextComponent;
  let fixture: ComponentFixture<LinkTextComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [LinkTextComponent],
      providers: [provideZonelessChangeDetection(), provideRouter([])]
    })
    .compileComponents();

    fixture = TestBed.createComponent(LinkTextComponent);
    fixture.componentRef.setInput('url', 'https://google.fr');
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

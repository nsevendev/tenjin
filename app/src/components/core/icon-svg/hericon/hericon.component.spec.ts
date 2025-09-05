import {ComponentFixture, TestBed} from '@angular/core/testing';
import {HericonComponent} from './hericon.component';
import {provideZonelessChangeDetection} from '@angular/core';

describe('HericonComponent', () => {
  let component: HericonComponent
  let fixture: ComponentFixture<HericonComponent>

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HericonComponent],
      providers: [provideZonelessChangeDetection()]
    })
      .compileComponents()

    fixture = TestBed.createComponent(HericonComponent);
    component = fixture.componentInstance
    fixture.componentRef.setInput('name', 'calendar');
    fixture.detectChanges()
  })

  it('should create', () => {
    expect(component).toBeTruthy()
  })
})

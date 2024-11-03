import { Overlay, OverlayRef } from '@angular/cdk/overlay';
import { ComponentPortal } from '@angular/cdk/portal';
import { coerceBooleanProperty } from '@angular/cdk/coercion';
import { Directive, ElementRef, HostListener, Injector, Input, OnDestroy, TemplateRef, ViewContainerRef } from '@angular/core';

import { TooltipComponent } from './tooltip.component';
import { TooltipPositions, TOOLTIP_POSITIONS } from './positions';

@Directive({ selector: '[appTooltip]' })
export class TooltipDirective implements OnDestroy {
  @Input() appTooltip!: string | TemplateRef<void>;
  @Input() appTooltipPosition: keyof TooltipPositions = 'auto';
  @Input()
  get appTooltipDisabled() { return this._appTooltipDisabled; }
  set appTooltipDisabled(v: string | boolean) {
    this._appTooltipDisabled = coerceBooleanProperty(v);
  }
  private _appTooltipDisabled = false;

  private _overlayRef?: OverlayRef;

  constructor(
    private readonly _element: ElementRef<HTMLElement>,
    private readonly _overlay: Overlay,
    private readonly _viewContainer: ViewContainerRef,
  ) { }

  ngOnDestroy() {
    this._overlayRef?.dispose();
  }

  @HostListener('mouseenter')
  @HostListener('focus')
  show() {
    if (this._appTooltipDisabled || this._overlayRef?.hasAttached() === true) {
      return;
    }

    this._attach();
  }

  @HostListener('mouseleave')
  @HostListener('blur')
  hide() {
    if (this._overlayRef?.hasAttached() === true) {
      this._overlayRef?.detach();
    }
  }

  private _attach() {
    if (!this._overlayRef) {
      this._overlayRef = this._overlay.create({
        panelClass: ['shadow-lg'],
        positionStrategy: this._overlay
          .position()
          .flexibleConnectedTo(this._element)
          .withPositions(TOOLTIP_POSITIONS[this.appTooltipPosition])
      });
    }

    const injector = Injector.create({
      providers: [
        {
          provide: 'TOOLTIP_DATA',
          useValue: this.appTooltip,
        }
      ]
    });

    this._overlayRef.attach(new ComponentPortal(
      TooltipComponent,
      this._viewContainer,
      injector
    ));
  }
}

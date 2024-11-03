import { Overlay, OverlayRef } from '@angular/cdk/overlay';
import { ComponentPortal } from '@angular/cdk/portal';
import { coerceBooleanProperty } from '@angular/cdk/coercion';
import { Directive, ElementRef, HostListener, Injector, Input, OnDestroy, TemplateRef, ViewContainerRef } from '@angular/core';
import { Subject, takeUntil } from 'rxjs';

import { PopoverComponent } from './popover.component';
import { PopoverPositions, POPOVER_POSITIONS } from './positions';

export type PopoverTrigger =
  'click' |
  'hover';

@Directive({ selector: '[appPopover]' })
export class PopoverDirective implements OnDestroy {
  @Input() appPopover!: string | TemplateRef<void>;
  @Input() appPopoverPosition: keyof PopoverPositions = 'top';
  @Input() appPopoverTriggers: PopoverTrigger | PopoverTrigger[] = 'click';
  @Input() appPopoverBackdropClass: string | string[] = [];
  @Input()
  get appPopoverDisabled() { return this._appPopoverDisabled; }
  set appPopoverDisabled(v: string | boolean) {
    this._appPopoverDisabled = coerceBooleanProperty(v);
  }
  private _appPopoverDisabled = false;

  private _overlayRef?: OverlayRef;
  private readonly _destroy$ = new Subject<void>();

  constructor(
    private readonly _element: ElementRef<HTMLElement>,
    private readonly _overlay: Overlay,
    private readonly _viewContainer: ViewContainerRef,
  ) { }

  ngOnDestroy() {
    this._overlayRef?.dispose();
    this._destroy$.next();
    this._destroy$.complete();
  }

  @HostListener('mouseenter')
  @HostListener('focus')
  onHover() {
    if (!this._hasTrigger('hover')) {
      return;
    }

    if (this._appPopoverDisabled || this._overlayRef?.hasAttached()) {
      return;
    }

    this._attach();
  }

  @HostListener('mouseleave')
  @HostListener('blur')
  onBlur() {
    if (!this._hasTrigger('hover')) {
      return;
    }

    if (!this._overlayRef?.hasAttached()) {
      return;
    }

    this._overlayRef?.detach();
  }

  @HostListener('click')
  onClick() {
    if (!this._hasTrigger('click')) {
      return;
    }

    if (this._overlayRef?.hasAttached()) {
      this._overlayRef?.detach();
      return;
    }

    this._attach();
  }

  private _attach() {
    if (!this._overlayRef) {
      this._overlayRef = this._overlay.create({
        panelClass: ['shadow-lg'],
        hasBackdrop: this._hasTrigger('click'),
        backdropClass: this.appPopoverBackdropClass,
        disposeOnNavigation: true,
        positionStrategy: this._overlay
          .position()
          .flexibleConnectedTo(this._element)
          .withPositions(POPOVER_POSITIONS[this.appPopoverPosition])
      });

      this._overlayRef.backdropClick().pipe(takeUntil(this._destroy$)).subscribe(() => {
        this._overlayRef?.detach();
      });
    }

    const injector = Injector.create({
      providers: [
        {
          provide: 'POPOVER_DATA',
          useValue: this.appPopover,
        }
      ]
    });

    this._overlayRef.attach(new ComponentPortal(
      PopoverComponent,
      this._viewContainer,
      injector
    ));
  }

  private _hasTrigger(trigger: PopoverTrigger) {
    const triggers = Array.isArray(this.appPopoverTriggers)
      ? this.appPopoverTriggers
      : [this.appPopoverTriggers];

    return !!triggers.find(v => v === trigger);
  }
}

import { ConnectedPosition } from '@angular/cdk/overlay';

export interface TooltipPositions {
  readonly auto: ConnectedPosition[];
  readonly top: ConnectedPosition[];
  readonly bottom: ConnectedPosition[];
  readonly left: ConnectedPosition[];
  readonly right: ConnectedPosition[];
  readonly 'top-left': ConnectedPosition[];
  readonly 'top-right': ConnectedPosition[];
  readonly 'bottom-left': ConnectedPosition[];
  readonly 'bottom-right': ConnectedPosition[];
}

export const TOOLTIP_POSITIONS: TooltipPositions = {
  auto: [
    {
      originX: 'center',
      originY: 'top',
      overlayX: 'center',
      overlayY: 'bottom',
      panelClass: ['top', 'mb-2']
    },
    {
      originX: 'center',
      originY: 'bottom',
      overlayX: 'center',
      overlayY: 'top',
      panelClass: ['bottom', 'mt-2']
    }
  ],
  top: [
    {
      originX: 'center',
      originY: 'top',
      overlayX: 'center',
      overlayY: 'bottom',
      panelClass: ['top', 'mb-2']
    }
  ],
  bottom: [
    {
      originX: 'center',
      originY: 'bottom',
      overlayX: 'center',
      overlayY: 'top',
      panelClass: ['bottom', 'mt-2']
    }
  ],
  left: [
    {
      originX: 'start',
      originY: 'center',
      overlayX: 'end',
      overlayY: 'center',
      panelClass: ['left', 'mr-2']
    }
  ],
  right: [
    {
      originX: 'end',
      originY: 'center',
      overlayX: 'start',
      overlayY: 'center',
      panelClass: ['right', 'ml-2']
    }
  ],
  'top-left': [
    {
      originX: 'start',
      originY: 'top',
      overlayX: 'end',
      overlayY: 'bottom',
      panelClass: ['top-left', 'mb-1', 'mr-1']
    }
  ],
  'top-right': [
    {
      originX: 'end',
      originY: 'top',
      overlayX: 'start',
      overlayY: 'bottom',
      panelClass: ['top-right', 'mb-1', 'ml-1']
    }
  ],
  'bottom-left': [
    {
      originX: 'start',
      originY: 'bottom',
      overlayX: 'end',
      overlayY: 'top',
      panelClass: ['bottom-left', 'mt-1', 'mr-1']
    }
  ],
  'bottom-right': [
    {
      originX: 'end',
      originY: 'bottom',
      overlayX: 'start',
      overlayY: 'top',
      panelClass: ['bottom-right', 'mt-1', 'ml-1']
    }
  ]
};

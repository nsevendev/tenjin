import {Injectable, Signal, WritableSignal, signal} from '@angular/core';

type RegistryItem = {
  get: () => boolean;
  set: (v: boolean) => void;
};

@Injectable({ providedIn: 'root' })
export class BottomSheetService {
  private registry = new Map<string, RegistryItem>();

  register(id: string, get: () => boolean, set: (v: boolean) => void) {
    this.registry.set(id, { get, set });
  }

  unregister(id: string) {
    this.registry.delete(id);
  }

  open(id: string) {
    this.registry.get(id)?.set(true);
  }

  close(id: string) {
    this.registry.get(id)?.set(false);
  }

  toggle(id: string) {
    const item = this.registry.get(id);
    if (!item) return;
    item.set(!item.get());
  }
}


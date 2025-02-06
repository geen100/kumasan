declare namespace H {
  namespace service {
    class Platform {
      constructor(options: { apikey: string });
      createDefaultLayers(): any;
    }
  }
  class Map {
    addObject(marker: any) {
      throw new Error("Method not implemented.");
    }
    static Marker: any;
    addEventListener(arg0: string, arg1: () => void) {
      throw new Error("Method not implemented.");
    }
    constructor(element: HTMLElement, defaultLayers: any, options: { zoom: number; center: { lat: number; lng: number } });
    setCenter(center: { lat: number; lng: number }): void;
    setZoom(zoom: number): void;
    getZoom(): number;
    dispose(): void;
  }
  namespace mapevents {
    class MapEvents {
      constructor(map: H.Map);
    }
    class Behavior {
      constructor(mapEvents: MapEvents);
    }
  }
  namespace ui {
    class UI {
      static createDefault(map: H.Map, defaultLayers: any): UI;
    }
  }
} 
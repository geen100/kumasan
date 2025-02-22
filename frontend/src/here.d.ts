declare namespace H {
  namespace service {
    class Platform {
      constructor(options: { apikey: string });
      createDefaultLayers(): any;
    }
  }
  class Map {
    addObject(marker: any): void;
    addEventListener(arg0: string, arg1: () => void): void;
    constructor(element: HTMLElement, defaultLayers: any, options: { zoom: number; center: { lat: number; lng: number } });
    setCenter(center: { lat: number; lng: number }): void;
    setZoom(zoom: number): void;
    getZoom(): number;
    dispose(): void;
  }
  namespace map {
    class Icon {
      constructor(url: string);
    }
    class Marker {
      constructor(coordinates: { lat: number; lng: number }, options?: { icon?: Icon });
    }
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
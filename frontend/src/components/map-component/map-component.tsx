import { useEffect, useRef, useState } from "react";
import "./map.styles.css";
import mapPin from "../../assets/icons/map-pin_icon.png";

function MapComponent() {
  const mapRef = useRef<H.Map | null>(null);
  const initialCenter = { lat: 37.9161, lng: 139.0364 };
  const initialZoom = 10;
  const [currentZoom, setCurrentZoom] = useState(initialZoom);

  useEffect(() => {
    const platform = new H.service.Platform({
      apikey: import.meta.env.VITE_HERE_MAP_API_KEY,
    });

    const defaultLayers = platform.createDefaultLayers();

    const mapElement = document.getElementById("map");
    if (mapElement) {
      const map = new H.Map(mapElement, defaultLayers.vector.normal.map, {
        zoom: initialZoom,
        center: initialCenter,
      });

      mapRef.current = map;

      const mapEvents = new H.mapevents.MapEvents(map);
      new H.mapevents.Behavior(mapEvents);
      H.ui.UI.createDefault(map, defaultLayers);

      map.addEventListener("mapviewchangeend", () => {
        setCurrentZoom(Math.round(map.getZoom()));
      });

      return () => {
        map.dispose();
      };
    }
  }, []);

  const resetMap = () => {
    if (mapRef.current) {
      const map = mapRef.current;

      map.setCenter(initialCenter);
      map.setZoom(initialZoom);
    }
  };

  return (
    <>
      <div className="map-container">
        <div id="map"></div>
      </div>
      <div className="map-control">
        <img onClick={resetMap} src={mapPin} alt="Kumasan" />
      </div>
    </>
  );
}

export default MapComponent;

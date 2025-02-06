import { useEffect, useRef, useState } from "react";
import "./map.styles.css";

function MapComponent() {
  const mapRef = useRef<H.Map | null>(null);
  const [initialCenter, setInitialCenter] = useState<{
    lat: number;
    lng: number;
  } | null>(null);
  const initialZoom = 10;
  const [currentZoom, setCurrentZoom] = useState(initialZoom);

  useEffect(() => {
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(
        (position) => {
          setInitialCenter({
            lat: position.coords.latitude,
            lng: position.coords.longitude,
          });
        },
        (error) => {
          console.error("Error getting location: ", error);
        }
      );
    }
  }, []);

  useEffect(() => {
    if (!initialCenter) return;
    console.log("initialCenter", initialCenter);

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
        setCurrentZoom(map.getZoom());
      });

      return () => {
        map.dispose();
      };
    }
  }, [initialCenter]);

  const resetMap = () => {
    if (mapRef.current) {
      const map = mapRef.current;

      if (initialCenter) {
        map.setCenter(initialCenter);
        map.setZoom(initialZoom);
        setCurrentZoom(initialZoom);
      }
    }
  };

  return (
    <>
      <div className="map-container">
        <div id="map"></div>
      </div>

      <div
        className={`map-control ${
          currentZoom !== initialZoom ? "visible" : ""
        }`}
      >
        <span className="tooltip">リセット</span>
        <span className="text">
          <img
            onClick={resetMap}
            src="https://img.icons8.com/ios-filled/100/recurring-appointment.png"
            alt="Kumasan"
          />
        </span>
      </div>

      {/* <div className="my-location-container">
        <div className="my-location-icon">
          <img
            src="https://img.icons8.com/keek/100/map-pin.png"
            alt="Kumasan"
          />
        </div>
      </div> */}
    </>
  );
}

export default MapComponent;

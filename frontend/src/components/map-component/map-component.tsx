import React, { useEffect } from "react";
import "./styles.css";

function MapComponent() {
  useEffect(() => {
    // Initialize the platform object
    const platform = new H.service.Platform({
      apikey: "z64shwIcaUhARZoFqoPx0eSuGxgcRm8XWELkCDFSnqU",
    });

    const defaultLayers = platform.createDefaultLayers();

    const map = new H.Map(
      document.getElementById("map"),
      defaultLayers.vector.normal.map,
      {
        zoom: 10,
        center: { lat: 37.9161, lng: 139.0364 },
      }
    );

    const mapEvents = new H.mapevents.MapEvents(map);
    const behavior = new H.mapevents.Behavior(mapEvents);
    const ui = H.ui.UI.createDefault(map, defaultLayers);

    return () => {
      map.dispose();
    };
  }, []);

  return <div id="map" className="map-container"></div>;
}

export default MapComponent;

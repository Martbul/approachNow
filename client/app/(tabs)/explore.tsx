import React, { useState, useEffect } from "react";
import { StyleSheet, View } from "react-native";
import MapView, { Marker, UrlTile } from "react-native-maps";
import * as Location from "expo-location";
import { TOMTOM_API_KEY } from "@/tomTomConfig";

interface UserLocation {
  latitude: number;
  longitude: number;
}

const App: React.FC = () => {
  const [location, setLocation] = useState<UserLocation | null>(null);
  const [nearbyUsers, setNearbyUsers] = useState<UserLocation[]>([]);
  const [ws, setWs] = useState<WebSocket | null>(null);

  useEffect(() => {
    // Initialize WebSocket connection when component mounts
    const websocket = new WebSocket("ws://192.168.100.9:9000/ws");

    websocket.onopen = () => {
      console.log("WebSocket connected");
    };

    websocket.onerror = (error) => {
      console.error("WebSocket error:", error);
    };

    websocket.onclose = () => {
      console.log("WebSocket connection closed");
    };

    websocket.onmessage = (message) => {
      try {
        const data = JSON.parse(message.data);
        if (data.nearbyUsers) {
          setNearbyUsers(data.nearbyUsers);
        }
      } catch (error) {
        console.error("Error parsing message data", error);
      }
    };

    setWs(websocket);

    return () => {
      // Clean up WebSocket connection when component unmounts
      websocket.close();
    };
  }, []);

  useEffect(() => {
    const getLocationAndSend = async () => {
      try {
        let { status } = await Location.requestForegroundPermissionsAsync();
        if (status !== "granted") {
          console.log("Permission to access location was denied");
          return;
        }

        let currentLocation = await Location.getCurrentPositionAsync({});
        const coords: UserLocation = {
          latitude: currentLocation.coords.latitude,
          longitude: currentLocation.coords.longitude,
        };
        setLocation(coords);

        // Send location to the server via WebSocket
        if (ws && ws.readyState === WebSocket.OPEN) {
          ws.send(
            JSON.stringify({
              latitude: coords.latitude,
              longitude: coords.longitude,
            })
          );
        }
      } catch (error) {
        console.error("Error fetching location", error);
      }
    };

    // Fetch location immediately on component mount
    getLocationAndSend();

    // Set up an interval to fetch and send the location every 10 seconds
    const locationInterval = setInterval(() => {
      getLocationAndSend();
    }, 10000);

    return () => {
      // Clear the interval when the component unmounts
      clearInterval(locationInterval);
    };
  }, [ws]);

  return (
    <View style={styles.container}>
      {location && (
        <MapView
          style={styles.map}
          initialRegion={{
            latitude: location.latitude,
            longitude: location.longitude,
            latitudeDelta: 0.05,
            longitudeDelta: 0.05,
          }}
        >
          <UrlTile
            urlTemplate={`https://api.tomtom.com/map/1/tile/basic/main/{z}/{x}/{y}.png?key=${TOMTOM_API_KEY}`}
            maximumZ={19}
          />
          <Marker
            coordinate={{
              latitude: location.latitude,
              longitude: location.longitude,
            }}
          />

          {nearbyUsers.map((user, index) => (
            <Marker
              key={index}
              coordinate={{
                latitude: user.latitude,
                longitude: user.longitude,
              }}
              title={`User ${index + 1}`}
            />
          ))}
        </MapView>
      )}
    </View>
  );
};

const styles = StyleSheet.create({
  container: { flex: 1 },
  map: { flex: 1 },
});

export default App;


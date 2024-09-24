// import React, { useState, useEffect } from "react";
// import { StyleSheet, View } from "react-native";
// import MapView, { Marker, UrlTile } from "react-native-maps";
// import * as Location from "expo-location";
// import { TOMTOM_API_KEY } from "@/tomTomConfig";
// import AsyncStorage from "@react-native-async-storage/async-storage";

// interface UserLocation {
//   latitude: number;
//   longitude: number;
// }

// const App: React.FC = () => {
//   const [location, setLocation] = useState<UserLocation | null>(null);
//   const [nearbyUsers, setNearbyUsers] = useState<UserLocation[]>([]);
//   const [ws, setWs] = useState<WebSocket | null>(null);

//   // Helper to get JWT token from AsyncStorage
//   const getToken = async (): Promise<string | null> => {
//     try {
//       const token = await AsyncStorage.getItem("jwtToken");
//       return token;
//     } catch (e) {
//       console.error("Failed to fetch JWT token", e);
//       return null;
//     }
//   };

//   useEffect(() => {
//     const initWebSocket = async () => {
//       const token = await getToken();
//       if (!token) {
//         console.error("No JWT token found");
//         return;
//       }

//       const websocket = new WebSocket("ws://192.168.100.9:9000/ws");

//       websocket.onopen = () => {
//         console.log("WebSocket connected");
//       };

//       websocket.onerror = (error) => {
//         console.error("WebSocket error:", error);
//       };

//       websocket.onclose = () => {
//         console.log("WebSocket connection closed");
//       };

//       websocket.onmessage = (message) => {
//         try {
//           const data = JSON.parse(message.data);
//           if (data.nearbyUsers) {
//             setNearbyUsers(data.nearbyUsers);
//           }
//         } catch (error) {
//           console.error("Error parsing message data", error);
//         }
//       };

//       setWs(websocket);

//       return () => {
//         websocket.close();
//       };
//     };

//     // Initialize WebSocket connection when component mounts
//     initWebSocket();
//   }, []);

//   useEffect(() => {
//     const getLocationAndSend = async () => {
//       try {
//         let { status } = await Location.requestForegroundPermissionsAsync();
//         if (status !== "granted") {
//           console.log("Permission to access location was denied");
//           return;
//         }

//         let currentLocation = await Location.getCurrentPositionAsync({});
//         const coords: UserLocation = {
//           latitude: currentLocation.coords.latitude,
//           longitude: currentLocation.coords.longitude,
//         };
//         setLocation(coords);

//         // Send location to the server via WebSocket
//         if (ws && ws.readyState === WebSocket.OPEN) {
//           const token = await getToken();
//           if (!token) {
//             console.error("No JWT token found");
//             return;
//           }

//           ws.send(
//             JSON.stringify({
//               latitude: coords.latitude,
//               longitude: coords.longitude,
//               token: token, // Add JWT token here
//             })
//           );
//         }

//         // Optionally send location to REST API as well
//         await sendLocationToApi(coords);
//       } catch (error) {
//         console.error("Error fetching location", error);
//       }
//     };

//     const sendLocationToApi = async (coords: UserLocation) => {
//       const token = await getToken();
//       if (!token) {
//         console.error("No JWT token found for API call");
//         return;
//       }

//       try {
//         const response = await fetch(
//           "http://192.168.100.9:9000/api/update-location",
//           {
//             method: "POST",
//             headers: {
//               "Content-Type": "application/json",
//               Authorization: `Bearer ${token}`, // Send JWT token in the Authorization header
//             },
//             body: JSON.stringify({
//               latitude: coords.latitude,
//               longitude: coords.longitude,
//             }),
//           }
//         );

//         if (!response.ok) {
//           throw new Error(`Failed to send location: ${response.status}`);
//         }

//         console.log("Location sent successfully");
//       } catch (error) {
//         console.error("Error sending location to API:", error);
//       }
//     };

//     // Fetch location immediately on component mount
//     getLocationAndSend();

//     // Set up an interval to fetch and send the location every 10 seconds
//     const locationInterval = setInterval(() => {
//       getLocationAndSend();
//     }, 10000);

//     return () => {
//       clearInterval(locationInterval);
//     };
//   }, [ws]);

//   // useEffect(() => {
//   //   // Initialize WebSocket connection when component mounts
//   //   const websocket = new WebSocket("ws://192.168.100.9:9000/ws");

//   //   websocket.onopen = () => {
//   //     console.log("WebSocket connected");
//   //   };

//   //   websocket.onerror = (error) => {
//   //     console.error("WebSocket error:", error);
//   //   };

//   //   websocket.onclose = () => {
//   //     console.log("WebSocket connection closed");
//   //   };

//   //   websocket.onmessage = (message) => {
//   //     try {
//   //       const data = JSON.parse(message.data);
//   //       if (data.nearbyUsers) {
//   //         setNearbyUsers(data.nearbyUsers);
//   //       }
//   //     } catch (error) {
//   //       console.error("Error parsing message data", error);
//   //     }
//   //   };

//   //   setWs(websocket);

//   //   return () => {
//   //     // Clean up WebSocket connection when component unmounts
//   //     websocket.close();
//   //   };
//   // }, []);

//   // useEffect(() => {
//   //   const getLocationAndSend = async () => {
//   //     try {
//   //       let { status } = await Location.requestForegroundPermissionsAsync();
//   //       if (status !== "granted") {
//   //         console.log("Permission to access location was denied");
//   //         return;
//   //       }

//   //       let currentLocation = await Location.getCurrentPositionAsync({});
//   //       const coords: UserLocation = {
//   //         latitude: currentLocation.coords.latitude,
//   //         longitude: currentLocation.coords.longitude,
//   //       };
//   //       setLocation(coords);

//   //       // Send location to the server via WebSocket
//   //       if (ws && ws.readyState === WebSocket.OPEN) {
//   //         ws.send(
//   //           JSON.stringify({
//   //             latitude: coords.latitude,
//   //             longitude: coords.longitude,
//   //           })
//   //         );
//   //       }
//   //     } catch (error) {
//   //       console.error("Error fetching location", error);
//   //     }
//   //   };

//   //   // Fetch location immediately on component mount
//   //   getLocationAndSend();

//   //   // Set up an interval to fetch and send the location every 10 seconds
//   //   const locationInterval = setInterval(() => {
//   //     getLocationAndSend();
//   //   }, 10000);

//   //   return () => {
//   //     // Clear the interval when the component unmounts
//   //     clearInterval(locationInterval);
//   //   };
//   // }, [ws]);

//   return (
//     <View style={styles.container}>
//       {location && (
//         <MapView
//           style={styles.map}
//           initialRegion={{
//             latitude: location.latitude,
//             longitude: location.longitude,
//             latitudeDelta: 0.05,
//             longitudeDelta: 0.05,
//           }}
//         >
//           <UrlTile
//             urlTemplate={`https://api.tomtom.com/map/1/tile/basic/main/{z}/{x}/{y}.png?key=${TOMTOM_API_KEY}`}
//             maximumZ={19}
//           />
//           <Marker
//             coordinate={{
//               latitude: location.latitude,
//               longitude: location.longitude,
//             }}
//           />

//           {nearbyUsers.map((user, index) => (
//             <Marker
//               key={index}
//               coordinate={{
//                 latitude: user.latitude,
//                 longitude: user.longitude,
//               }}
//               title={`User ${index + 1}`}
//             />
//           ))}
//         </MapView>
//       )}
//     </View>
//   );
// };

// const styles = StyleSheet.create({
//   container: { flex: 1 },
//   map: { flex: 1 },
// });

// export default App;



























// import React, { useState, useEffect, useContext } from "react";
// import { StyleSheet, View } from "react-native";
// import MapView, { Marker, UrlTile } from "react-native-maps";
// import * as Location from "expo-location";
// import { TOMTOM_API_KEY } from "@/tomTomConfig";
// import { AuthContext } from "@/contexts/AuthContext";


// interface UserLocation {
//   latitude: number;
//   longitude: number;
// }

// const App: React.FC = () => {
//     const { getToken, isLoading, user, setUser } = useContext(AuthContext);
//   const [location, setLocation] = useState<UserLocation | null>(null);
//   const [nearbyUsers, setNearbyUsers] = useState<UserLocation[]>([]);
//   const [ws, setWs] = useState<WebSocket | null>(null);

//   useEffect(() => {
//     const initWebSocket = async () => {
//       const token = getToken();
//       if (!token) {
//         console.error("No JWT token available");
//         return;
//       }

//       const websocket = new WebSocket(
//         `ws://192.168.100.9:9000/ws?token=${token}`
//       );

//       websocket.onopen = () => {
//         console.log("WebSocket connected");
//       };

//       websocket.onerror = (error) => {
//         console.error("WebSocket error:", error);
//       };

//       websocket.onclose = () => {
//         console.log("WebSocket connection closed");
//       };

//       websocket.onmessage = (message) => {
//         try {
//           const data = JSON.parse(message.data);
//           if (data.nearbyUsers) {
//             setNearbyUsers(data.nearbyUsers);
//           }
//         } catch (error) {
//           console.error("Error parsing message data", error);
//         }
//       };

//       setWs(websocket);

//       return () => {
//         websocket.close();
//       };
//     };

//     // Initialize WebSocket connection when component mounts
//     initWebSocket();
//   }, []);

//   useEffect(() => {
//     const getLocationAndSend = async () => {
//       try {
//         let { status } = await Location.requestForegroundPermissionsAsync();
//         if (status !== "granted") {
//           console.log("Permission to access location was denied");
//           return;
//         }

//         let currentLocation = await Location.getCurrentPositionAsync({});
//         const coords: UserLocation = {
//           latitude: currentLocation.coords.latitude,
//           longitude: currentLocation.coords.longitude,
//         };
//         setLocation(coords);

//         // Send location to the server via WebSocket
//         if (ws && ws.readyState === WebSocket.OPEN) {
//           ws.send(
//             JSON.stringify({
//               latitude: coords.latitude,
//               longitude: coords.longitude,
//             })
//           );
//         }

//         // Optionally send location to REST API as well
//         await sendLocationToApi(coords);
//       } catch (error) {
//         console.error("Error fetching location", error);
//       }
//     };

//     const sendLocationToApi = async (coords: UserLocation) => {
//       const token = getToken();
//       if (!token) {
//         console.error("No JWT token found for API call");
//         return;
//       }

//       try {
//         const response = await fetch(
//           //! fix
//           "http://192.168.100.9:9000/api/---------------------",
//           {
//             method: "POST",
//             headers: {
//               "Content-Type": "application/json",
//               Authorization: `Bearer ${token}`, // Send JWT token in the Authorization header
//             },
//             body: JSON.stringify({
//               latitude: coords.latitude,
//               longitude: coords.longitude,
//             }),
//           }
//         );

//         if (!response.ok) {
//           throw new Error(`Failed to send jwt token: ${response.status}`);
//         }

//         console.log("Location sent successfully");
//       } catch (error) {
//         console.error("Error sending jwt token:", error);
//       }
//     };

//     // Fetch location immediately on component mount
//     getLocationAndSend();

//     // Set up an interval to fetch and send the location every 10 seconds
//     const locationInterval = setInterval(() => {
//       getLocationAndSend();
//     }, 10000);

//     return () => {
//       clearInterval(locationInterval);
//     };
//   }, [ws]);

//   return (
//     <View style={styles.container}>
//       {location && (
//         <MapView
//           style={styles.map}
//           initialRegion={{
//             latitude: location.latitude,
//             longitude: location.longitude,
//             latitudeDelta: 0.05,
//             longitudeDelta: 0.05,
//           }}
//         >
//           <UrlTile
//             urlTemplate={`https://api.tomtom.com/map/1/tile/basic/main/{z}/{x}/{y}.png?key=${TOMTOM_API_KEY}`}
//             maximumZ={19}
//           />
//           <Marker
//             coordinate={{
//               latitude: location.latitude,
//               longitude: location.longitude,
//             }}
//           />

//           {nearbyUsers.map((user, index) => (
//             <Marker
//               key={index}
//               coordinate={{
//                 latitude: user.latitude,
//                 longitude: user.longitude,
//               }}
//               title={`User ${index + 1}`}
//             />
//           ))}
//         </MapView>
//       )}
//     </View>
//   );
// };

// const styles = StyleSheet.create({
//   container: { flex: 1 },
//   map: { flex: 1 },
// });

// export default App;








// import React, { useState, useEffect, useContext } from "react";
// import { StyleSheet, View } from "react-native";
// import MapView, { Marker, UrlTile } from "react-native-maps";
// import * as Location from "expo-location";
// import { TOMTOM_API_KEY } from "@/tomTomConfig";
// import { AuthContext } from "@/contexts/AuthContext";

// interface UserLocation {
//   latitude: number;
//   longitude: number;
// }

// const App: React.FC = () => {
//   const { getToken, user } = useContext(AuthContext); // Ensure you get the token and user from AuthContext
//   const [location, setLocation] = useState<UserLocation | null>(null);
//   const [nearbyUsers, setNearbyUsers] = useState<UserLocation[]>([]);
//   const [ws, setWs] = useState<WebSocket | null>(null);

//   // Initialize WebSocket connection
//   useEffect(() => {
//     const initWebSocket = async () => {
//       const token = getToken(); // Get the token from the context
//       if (!token) {
//         console.error("No JWT token available for WebSocket");
//         return;
//       }

//       const websocket = new WebSocket(
//         `ws://192.168.100.9:9000/ws?token=${token}` // Pass token as query param
//       );

//       websocket.onopen = () => {
//         console.log("WebSocket connected");
//       };

//       websocket.onerror = (error) => {
//         console.error("WebSocket error:", error);
//       };

//       websocket.onclose = () => {
//         console.log("WebSocket connection closed");
//       };

//       websocket.onmessage = (message) => {
//         try {
//           const data = JSON.parse(message.data);
//           if (data.nearbyUsers) {
//             setNearbyUsers(data.nearbyUsers); // Update nearby users
//           }
//         } catch (error) {
//           console.error("Error parsing message data", error);
//         }
//       };

//       setWs(websocket); // Save the WebSocket in state

//       return () => {
//         websocket.close(); // Close WebSocket on cleanup
//       };
//     };

//     // Initialize WebSocket connection when component mounts
//     initWebSocket();
//   }, [getToken]); // Re-run only if token changes

//   // Fetch location and send it via WebSocket or API every 10 seconds
//   useEffect(() => {
//     const getLocationAndSend = async () => {
//       try {
//         let { status } = await Location.requestForegroundPermissionsAsync();
//         if (status !== "granted") {
//           console.log("Permission to access location was denied");
//           return;
//         }

//         let currentLocation = await Location.getCurrentPositionAsync({});
//         const coords: UserLocation = {
//           latitude: currentLocation.coords.latitude,
//           longitude: currentLocation.coords.longitude,
//         };
//         setLocation(coords); // Save location

//         // Send location to the server via WebSocket
//         if (ws && ws.readyState === WebSocket.OPEN) {
//           ws.send(
//             JSON.stringify({
//               latitude: coords.latitude,
//               longitude: coords.longitude,
//             })
//           );
//         }

//         // Optionally send location to REST API
//         await sendLocationToApi(coords);
//       } catch (error) {
//         console.error("Error fetching location", error);
//       }
//     };

//     const sendLocationToApi = async (coords: UserLocation) => {
//       const token = getToken();
//       if (!token) {
//         console.error("No JWT token found for API call");
//         return;
//       }

//       try {
//         const response = await fetch(
//           "http://192.168.100.9:9000/api/location", // Your backend endpoint
//           {
//             method: "POST",
//             headers: {
//               "Content-Type": "application/json",
//               Authorization: `Bearer ${token}`, // Send JWT token in the Authorization header
//             },
//             body: JSON.stringify({
//               latitude: coords.latitude,
//               longitude: coords.longitude,
//             }),
//           }
//         );

//         if (!response.ok) {
//           throw new Error(`Failed to send jwt token: ${response.status}`);
//         }

//         console.log("Location sent successfully");
//       } catch (error) {
//         console.error("Error sending jwt token:", error);
//       }
//     };

//     // Fetch location immediately on component mount
//     getLocationAndSend();

//     // Set up an interval to fetch and send the location every 10 seconds
//     const locationInterval = setInterval(() => {
//       getLocationAndSend();
//     }, 10000);

//     return () => {
//       clearInterval(locationInterval); // Clear the interval when component unmounts
//     };
//   }, [ws, getToken]); // Ensure effect depends on WebSocket and token

//   return (
//     <View style={styles.container}>
//       {location && (
//         <MapView
//           style={styles.map}
//           initialRegion={{
//             latitude: location.latitude,
//             longitude: location.longitude,
//             latitudeDelta: 0.05,
//             longitudeDelta: 0.05,
//           }}
//         >
//           <UrlTile
//             urlTemplate={`https://api.tomtom.com/map/1/tile/basic/main/{z}/{x}/{y}.png?key=${TOMTOM_API_KEY}`}
//             maximumZ={19}
//           />
//           <Marker
//             coordinate={{
//               latitude: location.latitude,
//               longitude: location.longitude,
//             }}
//           />

//           {nearbyUsers.map((user, index) => (
//             <Marker
//               key={index}
//               coordinate={{
//                 latitude: user.latitude,
//                 longitude: user.longitude,
//               }}
//               title={`User ${index + 1}`}
//             />
//           ))}
//         </MapView>
//       )}
//     </View>
//   );
// };

// const styles = StyleSheet.create({
//   container: { flex: 1 },
//   map: { flex: 1 },
// });

// export default App;






















// import React, { useState, useEffect, useContext } from "react";
// import { StyleSheet, View } from "react-native";
// import MapView, { Marker, UrlTile } from "react-native-maps";
// import * as Location from "expo-location";
// import { TOMTOM_API_KEY } from "@/tomTomConfig";
// import { AuthContext } from "@/contexts/AuthContext";

// interface UserLocation {
//   latitude: number;
//   longitude: number;
// }

// const App: React.FC = () => {
//   const { getToken, user } = useContext(AuthContext); // Getting token and user context
//   const [location, setLocation] = useState<UserLocation | null>(null);
//   const [nearbyUsers, setNearbyUsers] = useState<UserLocation[]>([]);
//   const [ws, setWs] = useState<WebSocket | null>(null);

//   // Initialize WebSocket connection
//   useEffect(() => {
//     const initWebSocket = async () => {
//       const token = getToken(); // Get the token from the Auth context
//       if (!token) {
//         console.error("No JWT token available for WebSocket");
//         return;
//       }
//       console.log("GOT TOKEN", token)

//       const websocket = new WebSocket(
//         `ws://192.168.100.9:9000/ws?token=${token}` // Pass token as query param
//       );

//       websocket.onopen = () => {
//         console.log("WebSocket connected");
//       };

//       websocket.onerror = (error) => {
//         console.error("WebSocket error:", error);
//       };

//       websocket.onclose = () => {
//         console.log("WebSocket connection closed");
//       };

//       websocket.onmessage = (message) => {
//         try {
//           const data = JSON.parse(message.data);
//           if (data.nearbyUsers) {
//             setNearbyUsers(data.nearbyUsers); // Update nearby users if available
//           }
//         } catch (error) {
//           console.error("Error parsing message data", error);
//         }
//       };

//       setWs(websocket); // Save WebSocket instance in state

//       return () => {
//         websocket.close(); // Close WebSocket when component unmounts
//       };
//     };

//     // Initialize WebSocket connection when component mounts
//     initWebSocket();
//   }, [getToken]); // Re-run if the token changes

//   // Fetch and send location via WebSocket
//   useEffect(() => {
//     const getLocationAndSend = async () => {
//       try {
//         let { status } = await Location.requestForegroundPermissionsAsync();
//         if (status !== "granted") {
//           console.log("Permission to access location was denied");
//           return;
//         }

//         let currentLocation = await Location.getCurrentPositionAsync({});
//         const coords: UserLocation = {
//           latitude: currentLocation.coords.latitude,
//           longitude: currentLocation.coords.longitude,
//         };
//         setLocation(coords); // Save current location in state

//         // Send location to the server via WebSocket
//         if (ws && ws.readyState === WebSocket.OPEN) {
//           ws.send(
//             JSON.stringify({
//               latitude: coords.latitude,
//               longitude: coords.longitude,
//             })
//           );
//         }
//       } catch (error) {
//         console.error("Error fetching location", error);
//       }
//     };

//     // Fetch location immediately when the component mounts
//     getLocationAndSend();

//     // Set up an interval to fetch and send the location every 10 seconds
//     const locationInterval = setInterval(() => {
//       getLocationAndSend();
//     }, 10000);

//     return () => {
//       clearInterval(locationInterval); // Clear interval when component unmounts
//     };
//   }, [ws]); // Effect depends on WebSocket

//   return (
//     <View style={styles.container}>
//       {location && (
//         <MapView
//           style={styles.map}
//           initialRegion={{
//             latitude: location.latitude,
//             longitude: location.longitude,
//             latitudeDelta: 0.05,
//             longitudeDelta: 0.05,
//           }}
//         >
//           <UrlTile
//             urlTemplate={`https://api.tomtom.com/map/1/tile/basic/main/{z}/{x}/{y}.png?key=${TOMTOM_API_KEY}`}
//             maximumZ={19}
//           />
//           <Marker
//             coordinate={{
//               latitude: location.latitude,
//               longitude: location.longitude,
//             }}
//           />

//           {nearbyUsers.map((user, index) => (
//             <Marker
//               key={index}
//               coordinate={{
//                 latitude: user.latitude,
//                 longitude: user.longitude,
//               }}
//               title={`User ${index + 1}`}
//             />
//           ))}
//         </MapView>
//       )}
//     </View>
//   );
// };

// const styles = StyleSheet.create({
//   container: { flex: 1 },
//   map: { flex: 1 },
// });

// export default App;














import React, { useState, useEffect, useContext } from "react";
import { StyleSheet, View } from "react-native";
import MapView, { Marker, UrlTile } from "react-native-maps";
import * as Location from "expo-location";
import { TOMTOM_API_KEY } from "@/tomTomConfig";
import { AuthContext } from "@/contexts/AuthContext";

interface UserLocation {
  latitude: number;
  longitude: number;
}

const App: React.FC = () => {
  const { getToken, user } = useContext(AuthContext); // Getting token and user context
  const [location, setLocation] = useState<UserLocation | null>(null);
  const [nearbyUsers, setNearbyUsers] = useState<UserLocation[]>([]);
  const [ws, setWs] = useState<WebSocket | null>(null);

  useEffect(() => {
    const initWebSocket = async () => {
      const token = getToken(); // Get the token from the Auth context
      if (!token) {
        console.error("No JWT token available for WebSocket");
        return;
      }

      const websocket = new WebSocket(
        `ws://192.168.100.9:9000/ws?token=${token}` // Pass token as query param
      );

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
            setNearbyUsers(data.nearbyUsers); // Update nearby users if available
          }
        } catch (error) {
          console.error("Error parsing message data", error);
        }
      };

      setWs(websocket); // Save WebSocket instance in state

      return () => {
        websocket.close(); // Close WebSocket when component unmounts
      };
    };

    initWebSocket();
  }, [getToken]);

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
        setLocation(coords); // Save current location in state

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

    getLocationAndSend();

    const locationInterval = setInterval(() => {
      getLocationAndSend();
    }, 10000);

    return () => {
      clearInterval(locationInterval); // Clear interval when component unmounts
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

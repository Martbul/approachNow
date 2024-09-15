import React, { useState } from "react";
import { View, Button, Text, StyleSheet } from "react-native";
import { useRouter } from "expo-router";
import { grpcClient } from "../grpc/grpcClient";

export default function HomeScreen() {
  const router = useRouter();
  const [isConnected, setIsConnected] = useState(false);

  const handleConnect = async () => {
    try {
      await grpcClient.connectToHost(); // Establish connection with another phone
      setIsConnected(true);
      router.push("/camera"); // Navigate to CameraScreen
    } catch (error) {
      console.error("Connection failed", error);
    }
  };

  return (
    <View style={styles.container}>
      <Text style={styles.text}>Connect to Another Phone</Text>
      <Button
        title={isConnected ? "Connected" : "Connect"}
        onPress={handleConnect}
        disabled={isConnected}
      />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
  },
  text: {
    fontSize: 20,
    marginBottom: 20,
  },
});

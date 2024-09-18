import { StatusBar } from "expo-status-bar";
import { Image, ScrollView, Text, View } from "react-native";
import { Redirect, router } from "expo-router";
import { SafeAreaView } from "react-native-safe-area-context";
import { images } from "../constants";
import CustomButton from "../components/CustomButton";
import { AuthContext } from "../contexts/AuthContext";
import { useContext, useEffect, useState } from "react";
import { LoadingSpinner } from "@/components/LoadingSpinner";


// Define the type for the context (optional based on your AuthContext implementation)
interface AuthContextType {
  user: object | null;
  isLoading: boolean;
}

export default function App() {
  const { isLoading, user, getToken } = useContext(AuthContext);
  const [token,setToken] = useState("")

  useEffect(() => {
    const token = getToken();
    setToken(token)

  },[])

  // If the user data is still being fetched, show the loading spinner
  if (isLoading) {
    return <LoadingSpinner />;
  }

  // If the user is logged in, redirect to the dashboard
  if (!isLoading && token !== null) return <Redirect href="/(tabs)/home" />;

  // Render the login screen if the user is not logged in
  return (
    <SafeAreaView style={{ flex: 1, backgroundColor: "#161622" }}>
      <ScrollView contentContainerStyle={{ flexGrow: 1 }}>
        <View
          style={{
            flex: 1,
            justifyContent: "center",
            alignItems: "center",
            paddingHorizontal: 16,
          }}
        >
          <Text
            style={{
              fontSize: 14,
              fontFamily: "pregular",
              color: "#d1d5db",
              marginTop: 28,
              textAlign: "center",
            }}
          >
            Find you true love without the need of awkwardness
          </Text>

          {/* Make sure the CustomButton component has proper props for styling and behavior */}
          <CustomButton
            title="Register"
            handlePress={() => router.push("/(auth)/register")}
            containerStyles={{
              width: "100%",
              marginTop: 28,
              backgroundColor: "red",
            }}
          />
        </View>
      </ScrollView>

      <StatusBar backgroundColor="#161622" style="light" />
    </SafeAreaView>
  );
}

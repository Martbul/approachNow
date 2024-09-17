import { ActivityIndicator, StyleSheet, View } from "react-native";

export const LoadingSpinner = ({ size = "large", color = "#000000", style }) => (
  <View style={[styles.container, style]}>
    <ActivityIndicator size={size} color={color} />
  </View>
);
const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
  },
});

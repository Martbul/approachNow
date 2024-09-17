import {
  TouchableOpacity,
  Text,
  GestureResponderEvent,
  StyleProp,
  ViewStyle,
  TextStyle,
} from "react-native";
import React from "react";

// Define types for the props
interface CustomButtonProps {
  title: string; // The text displayed on the button
  handlePress: (event: GestureResponderEvent) => void; // Function to handle the button press event
  containerStyles?: string | StyleProp<ViewStyle>; // Optional custom container styles (string or StyleProp for inline styles)
  textStyles?: string | StyleProp<TextStyle>; // Optional custom text styles (string or StyleProp for inline styles)
  isLoading?: boolean; // Optional boolean to show a loading state
}

const CustomButton: React.FC<CustomButtonProps> = ({
  title,
  handlePress,
  containerStyles = "",
  textStyles = "",
  isLoading = false,
}) => {
  return (
    <TouchableOpacity
      onPress={handlePress}
      activeOpacity={0.7}
      className={`bg-red-500 rounded-xl min-h-[62px] justify-center items-center ${containerStyles} ${
        isLoading ? "opacity-50" : ""
      }`}
      disabled={isLoading}
    >
      <Text className={`text-primary font-psemibold text-lg ${textStyles}`}>
        {title}
      </Text>
    </TouchableOpacity>
  );
};

export default CustomButton;

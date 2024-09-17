import { View, Text, TextInput,TouchableOpacity, Image } from 'react-native'
import React, { useState } from 'react'
import { icons } from '@/constants';


const FormField = ({title, value, placeholder, otherStyles,handleChangeText,...props}) => {

    const [showPassword, setShowPassword] = useState(false)
    return (
      <View className={`space-y-2 ${otherStyles}`}>
        <Text className="text-base text-gray-100 font-pmedium">{title}</Text>

        <View className="w-full h-16 px-4 bg-tertiary rounded-2xl border-2 border-tertiary focus:border-secondary flex flex-row items-center">
          <TextInput
            className="flex-1 text-white font-psemibold text-base"
            value={value}
            placeholder={placeholder}
            placeholderTextColor="#7B7B8B"
            onChangeText={handleChangeText}
            secureTextEntry={title === "password" && !showPassword}
            {...props}
          />

          {title === "password" && (
            <TouchableOpacity onPress={() => setShowPassword(!showPassword)}>
              <Image
                source={!showPassword ? icons.eye : icons.eyeHide}
                className="w-6 h-6"
                resizeMode="contain"
              />
            </TouchableOpacity>
          )}
        </View>
      </View>
    );
}

export default FormField

// import React, { useState } from "react";
// import {
//   View,
//   Text,
//   TextInput,
//   TouchableOpacity,
//   Image,
//   TextInputProps,
//   ViewStyle,
// } from "react-native";
// import { icons } from "@/constants";

// // Define the types for the props
// interface FormFieldProps extends TextInputProps {
//   title: string;
//   value: string;
//   placeholder?: string;
//   otherStyles?: ViewStyle | string; // You may need to adjust this depending on how you're handling styles
//   handleChangeText: (text: string) => void;
// }

// const FormField: React.FC<FormFieldProps> = ({
//   title,
//   value,
//   placeholder,
//   otherStyles,
//   handleChangeText,
//   ...props
// }) => {
//   const [showPassword, setShowPassword] = useState(false);

//   return (
//     <View
//       style={[
//         { marginVertical: 8 },
//         typeof otherStyles === "string" ? {} : otherStyles,
//       ]}
//     >
//       <Text style={{ fontSize: 16, color: "#e0e0e0", fontFamily: "pmedium" }}>
//         {title}
//       </Text>

//       <View
//         style={{
//           width: "100%",
//           height: 64,
//           paddingHorizontal: 16,
//           backgroundColor: "#3a3a3a",
//           borderRadius: 16,
//           borderColor: "#3a3a3a",
//           borderWidth: 2,
//           flexDirection: "row",
//           alignItems: "center",
//         }}
//       >
//         <TextInput
//           style={{
//             flex: 1,
//             color: "#fff",
//             fontFamily: "psemibold",
//             fontSize: 16,
//           }}
//           value={value}
//           placeholder={placeholder}
//           placeholderTextColor="#7B7B8B"
//           onChangeText={handleChangeText}
//           secureTextEntry={title === "password" && !showPassword}
//           {...props}
//         />

//         {title === "password" && (
//           <TouchableOpacity onPress={() => setShowPassword(!showPassword)}>
//             <Image
//               source={!showPassword ? icons.eye : icons.eyeHide}
//               style={{ width: 24, height: 24 }}
//               resizeMode="contain"
//             />
//           </TouchableOpacity>
//         )}
//       </View>
//     </View>
//   );
// };

// export default FormField;

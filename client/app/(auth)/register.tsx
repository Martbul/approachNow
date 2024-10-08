import { View, Text, ScrollView, Image, Alert } from "react-native";
import React, { useContext, useState } from "react";
import { SafeAreaView } from "react-native-safe-area-context";
import { images } from "../../constants";
import FormField from "../../components/FormField";
import CustomButton from "../../components/CustomButton";
import { Link, router } from "expo-router";
import { AuthContext } from "../../contexts/AuthContext";

const SignUp = () => {
  const { registerUser, isLoading, user, setUser } = useContext(AuthContext);

  const [form, setForm] = useState({ username: "", email: "", password: "" });
  const [isSubmiting, setIsSubmiting] = useState(false);
  const submit = async () => {
    if (!form.username || !form.email || !form.password) {
      Alert.alert("Error", "Please fill in all fields");
      return;
    }
    setIsSubmiting(true);

    try {
      const result = await registerUser(
        form.username,
        form.email,
        form.password
      );

      console.log("result", result);
      setUser(result);
      router.replace("/(tabs)/home");
    } catch (error) {
      Alert.alert("Error", error.message);
    } finally {
      setIsSubmiting(false);
    }
  };
  return (
    <SafeAreaView className="bg-primary h-full">
      <ScrollView>
        <View className="w-full justify-center min-h-[83vh] px-4 my-6">
          <Image
            source={images.logo}
            className="w-[115px] h-[35px]"
            resizeMode="contain"
          />

          <Text className="text-2xl text-white semi-bold mt-10 font-psemibold">
            Sign up to Snit
          </Text>

          <FormField
            placeholder={"johnDoe"}
            title="username"
            value={form.username}
            handleChangeText={(e: string) => setForm({ ...form, username: e })}
            otherStyles="mt-10"
          ></FormField>
          <FormField
            placeholder={"johnDoe@gmail.com"}
            title="email"
            value={form.email}
            handleChangeText={(e: string) => setForm({ ...form, email: e })}
            otherStyles="mt-7"
            keyboardType="email-address"
          ></FormField>

          <FormField
            placeholder={"mysecretpass"}
            title="password"
            value={form.password}
            handleChangeText={(e: string) => setForm({ ...form, password: e })}
            otherStyles="mt-7"
          ></FormField>

          <CustomButton
            title="Sign In"
            handlePress={submit}
            containerStyles="mt-7"
            isLoading={isSubmiting}
          />
          <View className="justify-center pt-5 flex-row gap-2">
            <Text className="text-lg text-gray-100 font-pregular">
              Have an account already?
            </Text>
            <Link
              href="/(auth)/login"
              className="text-lg font-psemibold text-secondary"
            >
              Sign In
            </Link>
          </View>
        </View>
      </ScrollView>
    </SafeAreaView>
  );
};

export default SignUp;

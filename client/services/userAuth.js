// //! server side jwt token
// import { baseUrl, postRequest } from "../utils/request";

// // Store the token in memory or handle it via global state management
// let token = null;

// export const registerUser = async (username, email, password) => {
//   try {
//     const response = await postRequest(
//       `${baseUrl}/auth/register`,
//       JSON.stringify({ username, email, password })
//     );

//     if (response.error) {
//       console.error("Registration error", response.error);
//       throw new Error(response.error);
//     }

//     // Save the token in memory
//     token = response.token; // Assuming the token is part of the response

//     // Optionally, handle post-registration logic here
//     console.log("Registration successful", response);
//     return response;
//   } catch (error) {
//     console.error("Error during registration", error);
//     throw error;
//   }
// };

// export const loginUser = async (email, password) => {
//   try {
//     const response = await postRequest(
//       `${baseUrl}/auth/login`,
//       JSON.stringify({ email, password })
//     );

//     if (response.error) {
//       console.error("Login error", response.error);
//       throw new Error(response.error);
//     }

//     // Save the token in memory
//     token = response.token; // Assuming the token is part of the response

//     console.log("Login successful", response);
//     return response;
//   } catch (error) {
//     console.error("Error during login", error);
//     throw error;
//   }
// };

// export const logoutUser = async () => {
//   try {
//     // Optionally, notify the server to invalidate the token
//     await postRequest(`${baseUrl}/auth/logout`, JSON.stringify({ token }));

//     // Clear the token from memory
//     token = null;

//     console.log("Logout successful");
//   } catch (error) {
//     console.error("Error during logout", error);
//     throw error;
//   }
// };

// export const getToken = () => token;

// //! client side jwt token
// // import { baseUrl, postRequest } from "../utils/request";
// // import AsyncStorage from '@react-native-async-storage/async-storage';

// // export const registerUser = async (username, email, password) => {

// //     const response = await postRequest(
// //         `${baseUrl}/auth/register`,
// //         JSON.stringify({ username, email, password })
// //     );

// //     if (response.error) {
// //         console.log("error", response);
// //         throw new Error(response);
// //     }

// //     await AsyncStorage.setItem('jwtToken', JSON.stringify(response));
// // return response
// // }

// // export const loginUser = async (email, password) => {

// //     const response = await postRequest(
// //       `${baseUrl}/auth/login`,
// //       JSON.stringify({ email, password })
// //     );

// //     if (response.error) {
// //         throw new Error(response);
// //     }

// //     await AsyncStorage.setItem("jwtToken", JSON.stringify(response));
// //     return response
// // }

// // export const logoutUser = async() => {
// //     await AsyncStorage.removeItem("jwtToken");
// //     return null;
// // }

import { createContext, useState } from "react";
import { baseUrl, postRequest } from "../utils/request";

// Create the AuthContext
export const AuthContext = createContext();

export const AuthContextProvider = ({ children }) => {
  const [user, setUser] = useState(null); // User data
  const [token, setToken] = useState(null); // JWT token stored in memory
  const [isLoading, setIsLoading] = useState(false);

  // Register a new user
  const registerUser = async (username, email, password) => {
    try {
      setIsLoading(true);
      const response = await postRequest(
        `${baseUrl}/auth/register`,
        JSON.stringify({ username, email, password })
      );

      if (response.error) {
        console.error("Registration error", response.error);
        throw new Error(response.error);
      }

      // Save the token and user info in memory
      setToken(response.token);
      setUser({ username, email }); // You can adjust the user data returned from the server if needed

      console.log("Registration successful", response);
      return response;
    } catch (error) {
      console.error("Error during registration", error);
      throw error;
    } finally {
      setIsLoading(false);
    }
  };

  // Log in an existing user
  const loginUser = async (email, password) => {
    try {
      setIsLoading(true);
      const response = await postRequest(
        `${baseUrl}/auth/login`,
        JSON.stringify({ email, password })
      );

      if (response.error) {
        console.error("Login error", response.error);
        throw new Error(response.error);
      }

      // Save the token and user info in memory
      setToken(response.token);
      setUser({ email }); // Store other relevant user data if needed

      console.log("Login successful", response);
      return response;
    } catch (error) {
      console.error("Error during login", error);
      throw error;
    } finally {
      setIsLoading(false);
    }
  };

  // Log out the user
  const logoutUser = async () => {
    try {
      setIsLoading(true);
      // Optionally, notify the server to invalidate the token
      await postRequest(`${baseUrl}/auth/logout`, JSON.stringify({ token }));

      // Clear the token and user data from memory
      setToken(null);
      setUser(null);

      console.log("Logout successful");
    } catch (error) {
      console.error("Error during logout", error);
      throw error;
    } finally {
      setIsLoading(false);
    }
  };

  // Get the token
  const getToken = () => token;

  return (
    <AuthContext.Provider
      value={{
        user,
        token,
        isLoading,
        registerUser,
        loginUser,
        logoutUser,
        getToken,
        setUser
      }}
    >
      {children}
    </AuthContext.Provider>
  );
};

// import { createContext, useEffect, useState } from "react";
// import axios from "axios"; // For making requests to your server

// export const AuthContext = createContext();

// export const AuthContextProvider = ({ children }) => {
//   const [user, setUser] = useState(null); // User state to track login
//   const [isLoading, setIsLoading] = useState(false); // Loading state

//   // Function to check if the user is authenticated (session valid)
//   const checkAuthStatus = async () => {
//     setIsLoading(true);
//     try {
//       // Call your server to check if the user is authenticated
//       const response = await axios.get("/api/auth/status", {
//         withCredentials: true,
//       });
//       if (response.data.isAuthenticated) {
//         setUser(response.data.user); // Set the user data if authenticated
//       } else {
//         setUser(null); // Not authenticated
//       }
//     } catch (error) {
//       console.error("Failed to check auth status:", error);
//       setUser(null); // Not authenticated
//     } finally {
//       setIsLoading(false);
//     }
//   };

//   // Login function
//   const login = async (email, password) => {
//     setIsLoading(true);
//     try {
//       const response = await axios.post(
//         "/api/auth/login",
//         { email, password },
//         { withCredentials: true }
//       );
//       if (response.data.success) {
//         setUser(response.data.user); // Set the user data after successful login
//       }
//     } catch (error) {
//       console.error("Login failed:", error);
//       setUser(null);
//     } finally {
//       setIsLoading(false);
//     }
//   };

//   // Logout function
//   const logout = async () => {
//     setIsLoading(true);
//     try {
//       // Call the server to log out (destroy session or invalidate JWT)
//       await axios.post("/api/auth/logout", {}, { withCredentials: true });
//       setUser(null); // Clear the user data
//     } catch (error) {
//       console.error("Logout failed:", error);
//     } finally {
//       setIsLoading(false);
//     }
//   };

//   // Load auth status on component mount (useEffect)
//   useEffect(() => {
//     checkAuthStatus(); // Check if the user is already logged in on the server
//   }, []);

//   return (
//     <AuthContext.Provider value={{ user, isLoading, login, logout }}>
//       {children}
//     </AuthContext.Provider>
//   );
// };

// import { createContext, useEffect, useState } from "react";
// import AsyncStorage from '@react-native-async-storage/async-storage';

// export const AuthContext = createContext();

// export const AuthContextProvider = ({ children }) => {

//   const [user, setUser] = useState(null);
//   const [isLoading,setIsLoading] = useState(false);

//   const loadUser = async () => {
//     try {
//       setIsLoading(true);
//       const storedUser = await AsyncStorage.getItem("jwtToken");

//       if (storedUser!==null && !user) {
//         setUser(JSON.parse(storedUser));
//       }
//     } catch (error) {
//       console.error("Failed to load user", error);
//     }
//     finally {
//       setIsLoading(false);
//     }
//   };

//   useEffect(() => {
//     loadUser();

//   }, []);

//   return (
//     <AuthContext.Provider value={{ user,isLoading,setUser}}>
//       {children}
//     </AuthContext.Provider>

//   );
// };

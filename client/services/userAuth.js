import { baseUrl, postRequest } from "../utils/request";
import AsyncStorage from '@react-native-async-storage/async-storage';


export const registerUser = async (username, email, password) => {
 
    const response = await postRequest(
        `${baseUrl}/users/signup`,
        JSON.stringify({ username, email, password })
    );


    if (response.error) {
        console.log("error", response);
        throw new Error(response);
    }
   

    await AsyncStorage.setItem('user', JSON.stringify(response));
return response
}


export const loginUser = async (email, password) => {

    const response = await postRequest(
        `${baseUrl}/users/login`,
        JSON.stringify({ email, password })
    );
    
    if (response.error) {
        throw new Error(response);
    }

    await AsyncStorage.setItem('user', JSON.stringify(response));
    return response
}


export const logoutUser = async() => {
    await AsyncStorage.removeItem('user');
    return null;
}

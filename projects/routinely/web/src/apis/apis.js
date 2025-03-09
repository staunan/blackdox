import axios from "axios";
import { config } from "config/api_url.js";

let responseHeader = {
    "Access-Control-Allow-Headers": "Origin, X-Requested-With, Content-Type, Accept",
    "Access-Control-Allow-Credentials": true,
    "Access-Control-Allow-Origin": "http://localhost:5173",
    'Content-Type': 'application/json'
};
let reqConfig = {
    maxRedirects: 0,
    headers: responseHeader,
    withCredentials: true,
    credentials: 'include'
};

export async function createAccount(data){
    let response = await axios.post(config.api_base_url + 'create_account', data, reqConfig);
    return response.data;
}

export async function logoutUser(data){
    let response = await axios.post(config.api_base_url + 'logout', data, reqConfig);
    return response.data;
}

export async function createRoutine(data){
    let response = await axios.post(config.api_base_url + 'create_routine', data, reqConfig);
    return response.data;
}

export async function getAllRoutines(data){
    let response = await axios.post(config.api_base_url + 'all_routines', data, reqConfig);
    return response.data;
}

export async function getRoutineDetails(data){
    let response = await axios.post(config.api_base_url + 'routine_details', data, reqConfig);
    return response.data;
}

export async function markRoutineAsDone(data){
    let response = await axios.post(config.api_base_url + 'mark_routine_as_done', data, reqConfig);
    return response.data;
}

export async function markRoutineAsNotDone(data){
    let response = await axios.post(config.api_base_url + 'mark_routine_as_not_done', data, reqConfig);
    return response.data;
}

export async function getProgress(){
    let response = await axios.get(config.api_base_url + 'progress');
    return response.data;
}

export function getUserDefaultImage() {
    return config.api_base_url + "images/user_default_image.jpg";
}

export async function uploadProfilePhoto(data) {
    let formData = new FormData();
    formData.append('file', data.file, data.file.name);

    let req_config = {
        headers: {
            'Content-Type': `multipart/form-data; boundary=${formData._boundary}`,
            "Access-Control-Allow-Origin": true,
            "Accept-Language": 'en-US,en;q=0.8'
        },
        withCredentials: true,
        credentials: 'include',
        maxRedirects: 0,
    }
 
    let response = await axios.post(config.api_base_url + 'upload_photo', formData, req_config);
    return response.data;
}
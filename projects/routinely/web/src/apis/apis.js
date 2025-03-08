import axios from "axios";
import {config} from "config/api_url.js";

export async function createAccount(data){
    let header_config = {
        headers: {
            'Content-Type': 'application/json',
            "Access-Control-Allow-Origin": true,
            maxRedirects: 0,
            maxRate: [100 * 1024],
        },
    };
    let response = await axios.post(config.api_base_url + 'create_account', data, header_config);
    return response.data;
}

export async function createRoutine(data){
    let header_config = {
        headers: {
            'Content-Type': 'application/json',
            "Access-Control-Allow-Origin": true,
            maxRedirects: 0,
            maxRate: [100 * 1024],
        },
    };
    let response = await axios.post(config.api_base_url + 'create_routine', data, header_config);
    return response.data;
}

export async function getAllRoutines(data){
    let header_config = {
        headers: {
            'Content-Type': 'application/json',
            "Access-Control-Allow-Origin": true,
            maxRedirects: 0,
            maxRate: [100 * 1024],
        },
    };
    let response = await axios.post(config.api_base_url + 'all_routines', data, header_config);
    return response.data;
}

export async function getRoutineDetails(data){
    let header_config = {
        headers: {
            'Content-Type': 'application/json',
            "Access-Control-Allow-Origin": true,
            maxRedirects: 0,
            maxRate: [100 * 1024],
        },
    };
    let response = await axios.post(config.api_base_url + 'routine_details', data, header_config);
    return response.data;
}

export async function markRoutineAsDone(data){
    let header_config = {
        headers: {
            'Content-Type': 'application/json',
            "Access-Control-Allow-Origin": true,
            maxRedirects: 0,
            maxRate: [100 * 1024],
        },
    };
    let response = await axios.post(config.api_base_url + 'mark_routine_as_done', data, header_config);
    return response.data;
}

export async function markRoutineAsNotDone(data){
    let header_config = {
        headers: {
            'Content-Type': 'application/json',
            "Access-Control-Allow-Origin": true,
            maxRedirects: 0,
            maxRate: [100 * 1024],
        },
    };
    let response = await axios.post(config.api_base_url + 'mark_routine_as_not_done', data, header_config);
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

    let header_config = {
        headers: {
            'Accept': 'application/json',
            'Content-Type': `multipart/form-data; boundary=${formData._boundary}`,
            "Access-Control-Allow-Origin": true,
            'Accept-Language': 'en-US,en;q=0.8',
            maxRedirects: 0,
            maxRate: [100 * 1024],
        },
    };
    let response = await axios.post(config.api_base_url + 'upload_photo', formData, header_config);
    return response.data;
}
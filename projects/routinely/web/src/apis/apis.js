import axios from "axios";
import {config} from "config/api_url.js";

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
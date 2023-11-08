import axios from "axios";
import {config} from "../config.js";

export async function createTask(data){
    let header_config = {
        headers: {
            'Content-Type': 'application/json',
            "Access-Control-Allow-Origin": true,
            maxRedirects: 0,
            maxRate: [100 * 1024],
        },
    };
    let response = await axios.post(config.api_base_url + 'tasks/create_task', data, header_config);
    return response.data;
}

export async function updateTask(data){
    let header_config = {
        headers: {
            'Content-Type': 'application/json',
            "Access-Control-Allow-Origin": true,
            maxRedirects: 0,
            maxRate: [100 * 1024],
        },
    };
    let response = await axios.post(config.api_base_url + 'tasks/update_task', data, header_config);
    return response.data;
}

export async function getTasks(data){
    let header_config = {
        headers: {
            'Content-Type': 'application/json',
            "Access-Control-Allow-Origin": true,
            maxRedirects: 0,
            maxRate: [100 * 1024],
        },
    };
    let response = await axios.get(config.api_base_url + 'tasks/get_tasks', data, header_config);
    return response.data;
}

export async function deleteTask(data){
    let header_config = {
        headers: {
            'Content-Type': 'application/json',
            "Access-Control-Allow-Origin": true,
            maxRedirects: 0,
            maxRate: [100 * 1024],
        },
    };
    let response = await axios.post(config.api_base_url + 'tasks/delete_task', data, header_config);
    return response.data;
}
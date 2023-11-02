import axios from "axios";
import {config} from "../config.js";

export async function newSpending(data){
    let header_config = {
        headers: {
            'Content-Type': 'application/json',
            "Access-Control-Allow-Origin": true,
            maxRedirects: 0,
            maxRate: [100 * 1024],
        },
    };
    let response = await axios.post(config.api_base_url + 'spending/new_spending', data, header_config);
    return response.data;
}

export async function updateBookmark(data){
    let header_config = {
        headers: {
            'Content-Type': 'application/json',
            "Access-Control-Allow-Origin": true,
            maxRedirects: 0,
            maxRate: [100 * 1024],
        },
    };
    let response = await axios.post(config.api_base_url + 'bookmark/update_bookmark', data, header_config);
    return response.data;
}

export async function getSpendings(data){
    let header_config = {
        headers: {
            'Content-Type': 'application/json',
            "Access-Control-Allow-Origin": true,
            maxRedirects: 0,
            maxRate: [100 * 1024],
        },
    };
    let response = await axios.get(config.api_base_url + 'spending/get_spendings', data, header_config);
    return response.data;
}

export async function deleteSpending(data){
    let header_config = {
        headers: {
            'Content-Type': 'application/json',
            "Access-Control-Allow-Origin": true,
            maxRedirects: 0,
            maxRate: [100 * 1024],
        },
    };
    let response = await axios.post(config.api_base_url + 'spending/delete_spending', data, header_config);
    return response.data;
}
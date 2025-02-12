import axios from "axios";
import {config} from "../config.js";

export async function createBookmark(data){
    let header_config = {
        headers: {
            'Content-Type': 'application/json',
            "Access-Control-Allow-Origin": true,
            maxRedirects: 0,
            maxRate: [100 * 1024],
        },
    };
    let response = await axios.post(config.api_base_url + 'bookmark/create_bookmark', data, header_config);
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

export async function searchBookmark(data){
    let header_config = {
        headers: {
            'Content-Type': 'application/json',
            "Access-Control-Allow-Origin": true,
            maxRedirects: 0,
            maxRate: [100 * 1024],
        },
    };
    let response = await axios.post(config.api_base_url + 'bookmark/search_bookmark', data, header_config);
    return response.data;
}

export async function getBookmarks(data){
    let header_config = {
        headers: {
            'Content-Type': 'application/json',
            "Access-Control-Allow-Origin": true,
            maxRedirects: 0,
            maxRate: [100 * 1024],
        },
    };
    let response = await axios.get(config.api_base_url + 'bookmark/get_bookmarks', data, header_config);
    return response.data;
}

export async function deleteBookmark(data){
    let header_config = {
        headers: {
            'Content-Type': 'application/json',
            "Access-Control-Allow-Origin": true,
            maxRedirects: 0,
            maxRate: [100 * 1024],
        },
    };
    let response = await axios.post(config.api_base_url + 'bookmark/delete_bookmark', data, header_config);
    return response.data;
}
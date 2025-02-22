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
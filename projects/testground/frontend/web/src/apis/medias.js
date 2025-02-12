import axios from "axios";
import {config} from "../config.js";

export async function uploadPicture(data){
    let header_config = {
        headers: {
            'Content-Type': 'application/json',
            "Access-Control-Allow-Origin": true,
            maxRedirects: 0,
            maxRate: [100 * 1024],
        },
    };
    let response = await axios.post(config.api_base_url + 'media/upload_picture', data, header_config);
    return response.data;
}
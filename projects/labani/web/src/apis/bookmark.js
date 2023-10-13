import axios from "axios";
import config from "../config.js";

export async function createBookmark(data){
    let header_config = {
        headers: {
            'Content-Type': 'application/json',
            maxRedirects: 0,
            maxRate: [100 * 1024],
        },
    };
    let {data} = await axios.post(config.api_base_url + 'bookmark/create_bookmark', data, header_config);
    return data;
}
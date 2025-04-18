import { getUser } from "apis/apis.js";
import { TodayDayName } from "lib/js/datetime.js";
import { get, writable } from "svelte/store";

export const user_details = writable(null);
    
export let store = {
    getUser: async function () {
        let response = await getUser();
        if (!response.HasError) {
            user_details.set(response.Data)
        }
    },
    logout: function () {
        user_details.set(null);
    }
};
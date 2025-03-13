import { getUser, getAllRoutines } from "apis/apis.js";

import { writable, derived } from "svelte/store";

export const user_details = writable(null);
export const routines = writable([]);
export const inboxes = derived(routines, ($routines) => {
    console.log("routines", $routines);
    return routines;
});
    
export let store = {
    getUser: async function () {
        let response = await getUser();
        if (!response.HasError) {
            user_details.set(response.Data)
        }
    },
    getRoutines: async function () {
        try {
            let all_routines_response = await getAllRoutines({ user_id: 1 });
            if (!all_routines_response.HasError) { 
                routines.set(all_routines_response.Data);
            } else {
                routines.set([]);
            }
        } catch (err) {
            console.log(err);
            routines.set([]);
        }
        
        
    },
};
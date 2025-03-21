import { getUser, getAllRoutines } from "apis/apis.js";

import { get, writable, derived } from "svelte/store";

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
                this.isRoutineDirty = false;
            } else {
                this.isRoutineDirty = false;
                routines.set([]);
            }
        } catch (err) {
            console.log(err);
            routines.set([]);
        }
    },
    addRoutine: function (routine) {
        let all_routines = get(routines);
        if (all_routines.length > 0) {
            routines.set([...all_routines, routine]);
        }
    },
    updateRoutine: function (routine) {
        let all_routines = get(routines);
        if (all_routines.length > 0) { 
            let temp_arr = all_routines.map((r) => {
                if (r.ID == routine.ID) {
                    return routine;
                } else {
                    return r;
                }
            });
            routines.set(temp_arr);
        }
    },
    moveRoutineToTrash: function (routine) {
        let all_routines = get(routines);
        if (all_routines.length > 0) { 
            let temp_arr = all_routines.map((r) => {
                if (r.ID == routine.ID) {
                    return { ...r, IsTrash: 1 };
                } else {
                    return r;
                }
            });
            routines.set(temp_arr);
        }
    },
    restoreRoutineFromTrash: function (routine) {
        let all_routines = get(routines);
        if (all_routines.length > 0) { 
            let temp_arr = all_routines.map((r) => {
                if (r.ID == routine.ID) {
                    return { ...r, IsTrash: 0 };
                } else {
                    return r;
                }
            });
            routines.set(temp_arr);
        }
    }
};
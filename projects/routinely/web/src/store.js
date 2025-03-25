import { getUser, getAllRoutines, getProgress } from "apis/apis.js";
import { TodayDayName } from "lib/js/datetime.js";
import { get, writable } from "svelte/store";

export const user_details = writable(null);
export const routines = writable([]);
export const progress = writable([]);
export const inboxes = writable([]);

let routine_initialized = false;
let progress_initialized = false;
let initializing_inbox = false;
    
export let store = {
    init: async function () {
        await this.getRoutines();
        await this.getProgressData();
    },
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
            routine_initialized = true;
        } catch (err) {
            console.log(err);
            routines.set([]);
        }
    },
    addRoutine: function (routine) {
        let all_routines = get(routines);
        if (all_routines.length == 0) {
            routines.set([routine]);
        } else if (all_routines.length > 0) {
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
    },
    deleteRoutineForever: function (routine) {
        let all_routines = get(routines);
        if (all_routines.length > 0) { 
            let temp_arr = all_routines.filter((r) => {
                if (r.ID == routine.ID) {
                    return false;
                } else {
                    return true;
                }
            });
            routines.set(temp_arr);
        }
    },
    getProgressData: async function() {
        let progress_response = await getProgress({ user_id: 1 });
        if (progress_response.HasError) {
            progress.set([]);
        }
        if (progress_response.Data == null) {
            progress.set([]);
        } else {
            progress.set(progress_response.Data);
        }
        progress_initialized = true;
    },
    addEntry: function (entry) {
        let all_progress = get(progress);
        if (all_progress.length == 0) {
            progress.set([entry]);
        } else if (all_progress.length > 0) {
            progress.set([...all_progress, entry]);
        }
    },
    removeEntry: function (entry) {
        let all_progress = get(progress);
        if (all_progress.length > 0) { 
            let temp_arr = all_progress.filter((p) => {
                if (p.ID == entry.ID) {
                    return false;
                } else {
                    return true;
                }
            });
            progress.set(temp_arr);
        }
    },
    initializeInbox: async function () {
		let inbox_items = generateInboxItems(get(routines));
		inbox_items = applyProgressToInboxItems(inbox_items, get(progress));
        inboxes.set(inbox_items);
	}
};

routines.subscribe((r) => {
    store.initializeInbox();
});
progress.subscribe((p) => {
    store.initializeInbox();
});

function generateInboxItems(routines) {
    let inbox_items = [];
    for (let i = 0; i < routines.length; i++) {
        if (routines[i].Mode == "Daily") {
            if (validateDailyRoutine(routines[i])) {
                inbox_items.push(routines[i]);
            }
        }
    }
    return inbox_items;
}
function applyProgressToInboxItems(inbox_items, progress_items) {
    inbox_items.forEach((routine) => {
        let entry = null;
        if (progress_items) {
            for (let i = 0; i < progress_items.length; i++) {
                if (progress_items[i].RoutineID === routine.ID) {
                    entry = progress_items[i];
                    break;
                }
            }
        }
        if (entry) {
            routine.Done = true;
            routine.DoneData = entry;
        } else {
            routine.Done = false;
            routine.DoneData = null;
        }
    });
    return inbox_items;
}
function validateDailyRoutine(routine) {
    if (routine.IsTrash == 1) {
        return false;
    }
    if (routine.Status != "active") {
        return false;
    }
    let routine_days = routine.DailyBasisDays.split(",");
    if (!routine_days.includes(TodayDayName())) {
        return false;
    }
    return true;
}
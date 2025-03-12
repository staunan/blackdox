import { getUser } from "apis/apis.js";
    
export let store = {
    initialized: false,
    user_details: null,
    initialize: async function () {
        // get user details --
        let response = await getUser();
        if (!response.HasError) {
            this.user_details = response.Data;
            this.initialized = true;
            console.log(this);
        }
        return true;
    },
    setUserDetails: function (user) {
        this.user_details = user;
    }
};
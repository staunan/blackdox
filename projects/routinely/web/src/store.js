export let store = {
    initialized: false,
    user_details: null,
    initialize: function () {
        // get user details --
        this.initialized = true;
    },
    setUserDetails: function (user) {
        this.user_details = user;
    }
};
export function Now(){
    let current_datetime = new Date();
    let now = "";
    let date = current_datetime.getDate();
    if (date < 10) {
        date = "0" + date;
    }
    let month = current_datetime.getMonth() + 1;
    if (month < 10) {
        month = "0" + month;
    }
    let year = current_datetime.getFullYear();
    let hour = current_datetime.getHours();
    if (hour < 10) {
        hour = "0" + hour;
    }
    let minute = current_datetime.getMinutes();
    if (minute < 10) {
        minute = "0" + minute;
    }
    let second = current_datetime.getSeconds();
    if (second < 10) {
        second = "0" + second;
    }
    now = date+"-"+month+"-"+year+" "+hour+":"+minute+":"+second
    return now;
}
export function TodayDate(){
    let current_datetime = new Date();
    let now = "";
    let date = current_datetime.getDate();
    if (date < 10) {
        date = "0" + date;
    }
    let month = current_datetime.getMonth() + 1;
    if (month < 10) {
        month = "0" + month;
    }
    let year = current_datetime.getFullYear();
    now = year+"-"+month+"-"+date;
    return now;
}
export function YesterdayDate() {
    let current_datetime = new Date();
    current_datetime.setDate(current_datetime.getDate() - 1);
    let now = "";
    let date = current_datetime.getDate();
    if (date < 10) {
        date = "0" + date;
    }
    let month = current_datetime.getMonth() + 1;
    if (month < 10) {
        month = "0" + month;
    }
    let year = current_datetime.getFullYear();
    now = year+"-"+month+"-"+date;
    return now;
}
export function TodayDayName() {
    let days = ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"];
    let current_datetime = new Date();
    return days[current_datetime.getDay()];
}
export function ParseDateToRoutineHistory(date) {
    let date_part = date.substr(0, 10);
    if (date_part == TodayDate()) {
        return "<b>Today</b> at <b>" + ParseTimeToHumanReadableFormat(date.substr(11, 5)) + "</b>";
    } else if (date_part == YesterdayDate()) {
        return "<b>Yesterday</b> at " + ParseTimeToHumanReadableFormat(date.substr(11, 5)) + "</b>";
    } else {
        return date_part;
    }
    // console.log(date_part);
    // let d = convertMySQLDateTimeToJSDateTime(date);
    // return date;
}
function convertMySQLDateTimeToJSDateTime(date) {
    // Split timestamp into [ Y, M, D, h, m, s ]
    var t = "2010-06-09 13:12:01".split(/[- :]/);

    // Apply each element to the Date function
    var d = new Date(Date.UTC(t[0], t[1]-1, t[2], t[3], t[4], t[5]));

    return d;
}
export function ParseTimeToHumanReadableFormat(time) {
    let time_arr = time.split(":");
    let zone = "";
    let hour = 0;
    if (Number(time_arr[0]) < 12) {
        zone = "AM";
        hour = Number(time_arr[0]);
    } else {
        zone = "PM";
        hour = Number(time_arr[0]) - 12;
    }
    return hour + ":" + time_arr[1] + " " + zone;
}
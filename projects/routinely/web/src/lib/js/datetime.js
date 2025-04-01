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
}
export function ConvertMySQLDateTimeToJSDateTime(date) {
    // Split timestamp into [ Y, M, D, h, m, s ]
    var t = date.split(/[- :]/);
    // Apply each element to the Date function
    var d = new Date(t[0], t[1] - 1, t[2], t[3], t[4], t[5]);
    return d;
}
export function ConvertJSDateToMySQLDate(selected_date) {
    let now = "";
    let date = selected_date.getDate();
    if (date < 10) {
        date = "0" + date;
    }
    let month = selected_date.getMonth() + 1;
    if (month < 10) {
        month = "0" + month;
    }
    let year = selected_date.getFullYear();
    now = year+"-"+month+"-"+date;
    return now;
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
export function ParseDateToHumanReadableFormat(date) {
    if (!date) {
        return "";
    }
    let today = new Date();
    const months = [
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
    ];
    const weekdays = [
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	];
    let yesterday = new Date();
    yesterday.setDate(yesterday.getDate() - 1);
    let the_day_before_yesterday = new Date();
    the_day_before_yesterday.setDate(the_day_before_yesterday.getDate() - 2);

    if (today.getDate() == date.getDate() && today.getMonth() == date.getMonth() && today.getFullYear() == date.getFullYear()) {
        return "Today";
    } else if ((yesterday.getDate() == date.getDate() && yesterday.getMonth() == date.getMonth() && yesterday.getFullYear() == date.getFullYear())) {
        return "Yesterday";
    } else if ((the_day_before_yesterday.getDate() == date.getDate() && the_day_before_yesterday.getMonth() == date.getMonth() && the_day_before_yesterday.getFullYear() == date.getFullYear())) {
        return "The day before yesterday"
    } else {
        return weekdays[date.getDay()] + ", " + date.getDate() + " " + (months[date.getMonth() + 1]) + ", " + date.getFullYear();    
    }
}
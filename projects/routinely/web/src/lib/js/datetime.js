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
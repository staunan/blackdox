<div class="days_selector">
    <FormLabel label={label}></FormLabel>
    <div class="days">
        {#if days.length > 0}
            {#each days as day}
                <!-- svelte-ignore a11y_click_events_have_key_events -->
                 <!-- svelte-ignore a11y-no-static-element-interactions -->
                <div class="day" class:selected={ day.id == selectedDay.id } on:click={() => dayClickedHandler(day)} title={day.title}>{day.short}</div>
            {/each}
        {/if}
    </div>
</div>
<script>
import FormLabel from 'components/form/FormLabel.svelte';
export let label = "Input Label";
export let selectedItem = "";
let selectedDay = null;
let days = [
    {id: 1, short: "Sun", title: "Sunday"},
    {id: 2, short: "Mon", title: "Monday"},
    {id: 3, short: "Tue", title: "Tuesday"},
    {id: 4, short: "Wed", title: "Wednesday"},
    {id: 5, short: "Thu", title: "Thursday"},
    {id: 6, short: "Fri", title: "Friday"},
    {id: 7, short: "Sat", title: "Saturday"},
    {id: 8, short: "Any", title: "Anyday"},
];
$: {
    if(selectedItem){
        selectedDay = days.filter((day)=>day.short == selectedItem)[0];
    }
}
function dayClickedHandler(day){
    selectedDay = day;
}
</script>
<style>
.days{
    padding-top: 10px;
    display: flex;
    align-items: center;
    justify-content: left;
}
.day{
    display: flex;
    justify-content: center;
    align-items: center;
    width: 50px;
    height: 50px;
    cursor: pointer;
    border-radius: 50%;
    border: 1px solid #ccc;
    background-color: #fff;
    margin-right: 20px;
    margin-bottom: 20px;
    font-family: monospace;
    font-size: 18px;
    font-weight: bold;
}
.day:hover{
    background-color: #ccc;
}
.day.selected{
    color: #fff;
    box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
    background: #0F2027;  /* fallback for old browsers */
    background: -webkit-linear-gradient(to right, #2C5364, #203A43, #0F2027);  /* Chrome 10-25, Safari 5.1-6 */
    background: linear-gradient(to right, #2C5364, #203A43, #0F2027); /* W3C, IE 10+/ Edge, Firefox 16+, Chrome 26+, Opera 12+, Safari 7+ */
}
</style>
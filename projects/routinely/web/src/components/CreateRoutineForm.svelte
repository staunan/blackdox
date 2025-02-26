<script>
  import { onMount } from "svelte";
  import Card from "components/Card.svelte";
  import TextBox from "components/form/TextBox.svelte";
  import TextArea from "components/form/TextArea.svelte";
  import Dropdown from "components/form/Dropdown.svelte";
  import DaysSelector from "components/form/DaysSelector.svelte";
  import WeekDaySelector from "components/form/WeekDaySelector.svelte";
  import DayPicker from "components/form/DayPicker.svelte";
  import DateMonthPicker from "components/form/DateMonthPicker.svelte";
  import TimePicker from "components/form/TimePicker.svelte";
  import SubmitButton from "components/form/SubmitButton.svelte";
  import LinkButton from "components/form/LinkButton.svelte";
  import FormErrorMessage from "components/form/FormErrorMessage.svelte";
  import RoutineModeDisplayString from "components/RoutineModeDisplayString.svelte";
  import ArrowDown from "components/svg/ArrowDown.svelte";
  import SuccessModal from "components/modals/SuccessModal.svelte";
  import { createRoutine } from "apis/apis.js";

  // Form Settings Variable --
  let advanceSettings = false;
  let routine_title_label = "";
  let node_routine_mode_display_string;
  let routine_mode_display_string_has_error;
  let routine_mode_display_string_error_message;

  // Dropdown Data Variable --
  let all_routine_modes = [
    { label: "Daily", value: "Daily" },
    { label: "Weekly", value: "Weekly" },
    { label: "Monthly", value: "Monthly" },
    { label: "Yearly", value: "Yearly" },
  ];

  // Form Input Variables --
  let routine_title = "";
  let routineTitleHasError = false;
  let routineTitleErrorMessage = "";
  let routine_details = "";
  let routineDetailsHasError = false;
  let routineDetailsErrorMessage = "";
  let selected_routine_mode = null;
  let routineModeHasError = false;
  let routineModeErrorMessage = "";
  let selected_days = [];
  let selected_week_day = "";
  let selected_month_day = "";
  let selected_month_and_date = "";
  let selectedTime = "";

  $: {
    if (advanceSettings === true) {
      routine_title_label = "Routine Title";
    } else {
      routine_title_label = "";
    }
  }

  onMount(() => {
    // Set Default Values --
    routine_title = "";
    routine_details = "";
    selected_routine_mode = all_routine_modes[0]; // Daily
    selected_days = ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"];
    selected_week_day = "";
    selected_month_day = 0;
    selected_month_and_date = "00-00";
    selectedTime = "00:00";
  });

  function titleChangedHandler(event) {
    routine_title = event.detail;
    if (routine_title && routineTitleHasError) {
      routineTitleHasError = false;
      routineTitleErrorMessage = "";
    }
  }
  function routineDetailsChangeHandler(event) {
    routine_details = event.detail;
    if (routine_details && routineDetailsHasError) {
      routineDetailsHasError = false;
      routineDetailsErrorMessage = "";
    }
  }
  function routineModeChangeHandler(event) {
    selected_routine_mode = event.detail;
    if (selected_routine_mode && routineModeHasError) {
      routineModeHasError = false;
      routineModeErrorMessage = "";
    }
    if (selected_routine_mode.value == "Daily") {
      selected_days = ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"];
      selectedTime = "00:00";
      generateRoutineModeDisplayString();
    } else if (selected_routine_mode.value == "Weekly") {
      selected_week_day = "";
      selectedTime = "00:00";
      generateRoutineModeDisplayString();
    } else if (selected_routine_mode.value == "Monthly") {
      selected_month_day = 0;
      selectedTime = "00:00";
      generateRoutineModeDisplayString();
    } else if (selected_routine_mode.value == "Yearly") {
      selected_month_and_date = "00-00";
      selectedTime = "00:00";
      generateRoutineModeDisplayString();
    }
  }
  function dailyDaysChangedHandler(event) {
    selected_days = event.detail;
    generateRoutineModeDisplayString();
  }
  function weekDayChangedHandler(event) {
    selected_week_day = event.detail;
    generateRoutineModeDisplayString();
  }
  function monthDayChangedHandler(event) {
    selected_month_day = event.detail;
    generateRoutineModeDisplayString();
  }
  function yearMonthDateChangedHandler(event) {
    selected_month_and_date = event.detail;
    generateRoutineModeDisplayString();
  }
  function routineTimeChangedHandler(event) {
    selectedTime = event.detail;
    generateRoutineModeDisplayString();
  }
  function generateRoutineModeDisplayString() {
    let str = "";
    if (selected_routine_mode.value == "Daily") {
      let days_str = "";
      if (selected_days.length === 7) {
        days_str = "<b>Everyday</b> ";
      } else if (selected_days.length === 0) {
        routine_mode_display_string_has_error = true;
        routine_mode_display_string_error_message =
          "Please select at least one day";
        return;
      } else {
        days_str = "Every <b>" + selected_days.join(", ") + "</b> ";
      }
      if (selectedTime === "00:00") {
        str = days_str + "at <b>Anytime</b>";
      } else {
        str = days_str + "at <b>" + selectedTime + "</b>";
      }
      routine_mode_display_string_has_error = false;
    } else if (selected_routine_mode.value == "Weekly") {
      let days_str = "";
      if (selected_week_day) {
        days_str = "Every <b>" + selected_week_day + "</b> ";
      } else {
        days_str = "Anyday of the week ";
      }
      if (selectedTime === "00:00") {
        str = days_str + "at <b>Anytime</b>";
      } else {
        str = days_str + "at <b>" + selectedTime + "</b>";
      }
      routine_mode_display_string_has_error = false;
    } else if (selected_routine_mode.value == "Monthly") {
      let days_str = "";
      if (selected_month_day) {
        let monthday = parseInt(selected_month_day);
        if (monthday == 0) {
          days_str = "Anyday of the month ";
        } else if (monthday == 1) {
          days_str = "Every 1<sup>st</sup> of the month ";
        } else if (monthday == 2) {
          days_str = "Every 2<sup>nd</sup> of the month ";
        } else if (monthday == 3) {
          days_str = "Every 3<sup>rd</sup> of the month ";
        } else if (monthday == 32) {
          days_str = "Every end of the month ";
        } else if (monthday == 33) {
          days_str = "Yestereday of every end of the month ";
        } else {
          days_str =
            "Every " + selected_month_day + "<sup>th</sup> of the month ";
        }
      } else {
        days_str = "Anyday of the month ";
      }
      if (selectedTime === "00:00") {
        str = days_str + "at <b>Anytime</b>";
      } else {
        str = days_str + "at <b>" + selectedTime + "</b>";
      }
      routine_mode_display_string_has_error = false;
    } else if (selected_routine_mode.value == "Yearly") {
      let Months = [
        { label: "Anyday of the Year", value: 0 },
        { label: "January", value: 1 },
        { label: "February", value: 2 },
        { label: "March", value: 3 },
        { label: "April", value: 4 },
        { label: "May", value: 5 },
        { label: "June", value: 6 },
        { label: "July", value: 7 },
        { label: "August", value: 8 },
        { label: "September", value: 9 },
        { label: "October", value: 10 },
        { label: "November", value: 11 },
        { label: "December", value: 12 },
      ];
      let days_str = "";
      if (selected_month_and_date) {
        let monthday = selected_month_and_date.split("-");
        let m = parseInt(monthday[0]);
        let d = parseInt(monthday[1]);
        if (m === 0) {
          days_str = "Anyday of the year ";
        } else {
          let month_str = Months.filter((tm) => tm.value === m)[0].label;
          if (d == 0) {
            days_str = "Anyday of the year";
          } else if (d == 1) {
            days_str = "Every 1<sup>st</sup> of " + month_str + " ";
          } else if (d == 2) {
            days_str = "Every 2<sup>nd</sup> of " + month_str + " ";
          } else if (d == 3) {
            days_str = "Every 3<sup>rd</sup> of " + month_str + " ";
          } else {
            days_str = "Every " + d + "<sup>th</sup> of " + month_str + " ";
          }
        }
      } else {
        days_str = "Anyday of the year ";
      }
      if (selectedTime === "00:00") {
        str = days_str + "at <b>Anytime</b>";
      } else {
        str = days_str + "at <b>" + selectedTime + "</b>";
      }
      routine_mode_display_string_has_error = false;
    }
    node_routine_mode_display_string = str;
  }
  function showAdvanceSettingHandler() {
    advanceSettings = !advanceSettings;
    setTimeout(() => generateRoutineModeDisplayString(), 200);
  }
  function validateRoutineForm() {
    let routineObj = {};
    let hasError = false;
    // Routine Title --
    if (routine_title == "") {
      hasError = true;
      routineTitleHasError = true;
      routineTitleErrorMessage = "Please provide a title";
    } else {
      routineTitleHasError = false;
      routineTitleErrorMessage = "";
      routineObj.title = routine_title;
    }
    // Routine Description --
    routineObj.description = routine_details;
    // Routine Mode --
    if (selected_routine_mode == null) {
      hasError = true;
      routineModeHasError = true;
      routineModeErrorMessage = "Please select an execution mode";
    } else {
      routineModeHasError = false;
      routineModeErrorMessage = "";
      routineObj.mode = selected_routine_mode.value;
      if (routineObj.mode === "Daily") {
        // Routine Days --
        if (selected_days && selected_days.length > 0) {
          routine_mode_display_string_has_error = false;
          routine_mode_display_string_error_message = "";
          routineObj.days = selected_days.join(",");
        } else if (selected_days && selected_days.length === 0) {
          hasError = true;
          routine_mode_display_string_has_error = true;
          routine_mode_display_string_error_message =
            "Please select at least on day above";
        } else {
          hasError = true;
        }
        // Routine Time --
        if (selectedTime) {
          routineObj.time = selectedTime;
        }
      } else if (routineObj.mode === "Weekly") {
        // Routine WeekDay --
        if (selected_week_day) {
          routine_mode_display_string_has_error = false;
          routine_mode_display_string_error_message = "";
          routineObj.weekday = selected_week_day;
        } else {
          routine_mode_display_string_has_error = false;
          routine_mode_display_string_error_message = "";
          routineObj.weekday = ""; // Empty string means anyday
        }
        // Routine Time --
        if (selectedTime) {
          routineObj.time = selectedTime;
        }
      } else if (routineObj.mode === "Monthly") {
        // Routine Month Day --
        if (selected_month_day) {
          routine_mode_display_string_has_error = false;
          routine_mode_display_string_error_message = "";
          routineObj.monthday = selected_month_day;
        } else {
          routine_mode_display_string_has_error = false;
          routine_mode_display_string_error_message = "";
          routineObj.monthday = 0; // 0 Means anyday
        }
        // Routine Time --
        if (selectedTime) {
          routineObj.time = selectedTime;
        }
      } else if (routineObj.mode === "Yearly") {
        // Routine Yearly Month and Date --
        if (selected_month_and_date) {
          routine_mode_display_string_has_error = false;
          routine_mode_display_string_error_message = "";
          routineObj.yearlymonthdate = selected_month_and_date;
        } else {
          routine_mode_display_string_has_error = false;
          routine_mode_display_string_error_message = "";
          routineObj.yearlymonthdate = "00-00";
        }
        // Routine Time --
        if (selectedTime) {
          routineObj.time = selectedTime;
        }
      }
    }
    console.log(routineObj);
    if (hasError) {
      return false;
    } else {
      return routineObj;
    }
  }
  async function createRoutineHandler(event) {
    let routineObj = validateRoutineForm();
    if (routineObj === false) {
      return;
    }
    let formData = {
      routine_title: routine_title,
    };
    try {
      let response = await createRoutine(formData);
      console.log(response);
    } catch (error) {
      console.log(error);
    }
  }
</script>

<Card>
  <div class="card_body">
    <h1 class="center mb10 form_heading">Create Routine</h1>
    <TextBox
      label={routine_title_label}
      placeholder="Write a title for your routine..."
      on:change={titleChangedHandler}
      value={routine_title}
      hasError={routineTitleHasError}
      errorMessage={routineTitleErrorMessage}
    ></TextBox>
    {#if advanceSettings}
      <div class="advanceSettings">
        <div class="form_gap"></div>
        <Dropdown
          label="Routine Mode"
          placeholder="Select a execution mode of this routine "
          items={all_routine_modes}
          currentitem={selected_routine_mode}
          on:change={routineModeChangeHandler}
          hasError={routineModeHasError}
          errorMessage={routineModeErrorMessage}
        />
        {#if selected_routine_mode && selected_routine_mode.value === "Daily"}
          <div class="form_gap"></div>
          <div class="form_row">
            <div class="form_column">
              <DaysSelector
                value={selected_days}
                on:change={dailyDaysChangedHandler}
                label="Select Days"
              ></DaysSelector>
            </div>
            <div class="form_column">
              <TimePicker
                value={selectedTime}
                on:change={routineTimeChangedHandler}
                label="Choose a time (24 Hour Format)"
                format="24Hours"
              ></TimePicker>
            </div>
          </div>
        {:else if selected_routine_mode && selected_routine_mode.value === "Weekly"}
          <div class="form_gap"></div>
          <div class="form_row">
            <div class="form_column">
              <WeekDaySelector
                value={selected_week_day}
                on:change={weekDayChangedHandler}
                label="Select a day"
              ></WeekDaySelector>
            </div>
            <div class="form_column">
              <TimePicker
                value={selectedTime}
                on:change={routineTimeChangedHandler}
                label="Choose a time (24 Hour Format)"
                format="24Hours"
              ></TimePicker>
            </div>
          </div>
        {:else if selected_routine_mode && selected_routine_mode.value === "Monthly"}
          <div class="form_gap"></div>
          <div class="form_row">
            <div class="form_column">
              <DayPicker
                value={selected_month_day}
                on:change={monthDayChangedHandler}
                label="Choose a Day"
              ></DayPicker>
            </div>
            <div class="form_column">
              <TimePicker
                value={selectedTime}
                on:change={routineTimeChangedHandler}
                label="Choose a time (24 Hour Format)"
                format="24Hours"
              ></TimePicker>
            </div>
          </div>
        {:else if selected_routine_mode && selected_routine_mode.value === "Yearly"}
          <div class="form_gap"></div>
          <div class="form_row">
            <div class="form_column">
              <DateMonthPicker
                value={selected_month_and_date}
                on:change={yearMonthDateChangedHandler}
                label="Choose a date (Month-Day)"
              ></DateMonthPicker>
            </div>
            <div class="form_column">
              <TimePicker
                value={selectedTime}
                on:change={routineTimeChangedHandler}
                label="Choose a time (24 Hour Format)"
                format="24Hours"
              ></TimePicker>
            </div>
          </div>
        {/if}
        {#if !routine_mode_display_string_has_error}
          <RoutineModeDisplayString message={node_routine_mode_display_string}
          ></RoutineModeDisplayString>
        {:else}
          <div class="mode_error_message">
            <FormErrorMessage
              message={routine_mode_display_string_error_message}
            ></FormErrorMessage>
          </div>
        {/if}

        <div class="form_gap"></div>
        <TextArea
          label="Routine Description [Optional]"
          placeholder="Write a description for your routine (Optional)..."
          on:change={routineDetailsChangeHandler}
          value={routine_details}
          hasError={routineDetailsHasError}
          errorMessage={routineDetailsErrorMessage}
        ></TextArea>

        <div class="advance_settings_hide">
          <LinkButton
            label="Hide Advance Settings"
            on:tap={(advanceSettings = !advanceSettings)}
          ></LinkButton>
        </div>
      </div>
    {:else}
      <div class="advance_settings_trigger">
        <LinkButton
          label="Show Advance Settings"
          on:tap={showAdvanceSettingHandler}
        ></LinkButton>
        <div class="advance_settings_trigger_icon">
          <ArrowDown></ArrowDown>
        </div>
      </div>
    {/if}

    <div class="center mt10 submit_button_container">
      <SubmitButton title="Create Routine" on:tap={createRoutineHandler}
      ></SubmitButton>
    </div>
  </div>
</Card>

<SuccessModal active={true}></SuccessModal>

<style>
  .form_heading {
    font-family: monospace;
    font-size: 24px;
    font-weight: bold;
    margin-top: 10px;
    padding-bottom: 40px;
  }
  .form_gap {
    width: 100%;
    height: 20px;
  }
  .card_body {
    padding: 20px;
  }
  .center {
    display: flex;
    justify-content: center;
    align-items: center;
  }
  .advance_settings_trigger {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    padding-top: 20px;
    position: relative;
  }
  .advance_settings_trigger_icon {
    position: relative;
    bottom: 10px;
    cursor: pointer;
  }
  .advance_settings_hide {
    padding-top: 40px;
    padding-bottom: 10px;
    display: flex;
  }
  .submit_button_container {
    margin-bottom: 10px;
    display: flex;
    justify-content: center;
    align-items: center;
    padding-bottom: 10px;
  }
  .form_row {
    display: flex;
    flex-wrap: nowrap;
    box-sizing: border-box;
    justify-content: space-between;
  }
  .form_column {
    flex-basis: 45%;
  }
  .mb10 {
    margin-bottom: 10px;
  }
  .mt10 {
    margin-top: 10px;
  }
  .mode_error_message {
    padding-top: 5px;
    padding-bottom: 10px;
  }
</style>

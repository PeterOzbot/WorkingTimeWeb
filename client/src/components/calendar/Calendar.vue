<template>
  <div class="calendar">
    <v-card class="calendar-month" elevation="3">
      <div class="calendar-month-header">
        <MonthHeaderComponent :currentMonth="currentMonth" :currentYear="currentYear" />
      </div>
      <WeekdaysComponent />
    </v-card>
    <v-card class="calendar-days" elevation="5">
      <ol class="days-grid">
        <DayComponent v-for="(day, dayIndex) in workingDays" :key="dayIndex" :day="day" />
      </ol>
    </v-card>
  </div>
</template>

<script lang="ts" setup>
import WeekdaysComponent from "./Weekdays.vue";
import MonthHeaderComponent from "./MonthHeader.vue";
import DayComponent from "./Day.vue";
import WorkingDay from "@/models/workingDay";
import { watch, ref } from "vue";

const props = defineProps({
  generatedWorkingDays: { type: Array<WorkingDay>, required: true },
  currentMonth: { type: Number, required: true },
  currentYear: { type: Number, required: true },
});

let workingDays = ref(new Array<WorkingDay>());

// when generated days change prepare calendar
watch(() => props.generatedWorkingDays, () => { prepareCalendar(); });

function prepareCalendar() {
  workingDays.value = [
    ...getDaysInPreviousMonth(),
    ...props.generatedWorkingDays,
    ...getDaysInNextMonth()
  ];
}

function getDaysInPreviousMonth() {
  const days: WorkingDay[] = [];
  // create first of the month
  const date = new Date(props.currentYear, props.currentMonth, 1);

  // get which day of the week
  const weekDay = date.getDay();

  // how many days we need from previous month
  let previousMonthDays = weekDay ? weekDay - 1 : 6;

  // get the first day we need from previous month
  date.setDate(date.getDate() - previousMonthDays);

  // add days if there are any left
  while (previousMonthDays > 0) {
    // add day
    days.push(new WorkingDay(new Date(date), 0, true));

    // increment day
    date.setDate(date.getDate() + 1);
    // reduce remaining days
    previousMonthDays--;
  }

  return days;
}

function getDaysInNextMonth() {
  const days: WorkingDay[] = [];

  // create last day of the month
  const date = new Date(props.currentYear, props.currentMonth + 1, 0);

  // get which day of the week
  const weekDay = date.getDay();

  // how many days we need from next month
  let nextMonthDays = weekDay ? 7 - weekDay : 0;

  // get the first day we need from previous month
  date.setDate(date.getDate() + 1);

  // add days if there are any left
  while (nextMonthDays > 0) {
    // add day
    days.push(new WorkingDay(new Date(date), 0, true));

    // increment day
    date.setDate(date.getDate() + 1);
    // reduce remaining days
    nextMonthDays--;
  }

  return days;
}
</script>

<style scoped>
.calendar-month {
  position: relative;
  margin: 10px;
}

.calendar-month-header {
  display: flex;
  justify-content: center;
  background-color: #94cb94;
  padding: 10px;
}

.days-grid {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  height: 100%;
  position: relative;
  grid-column-gap: 1px;
  grid-row-gap: 1px;
  background-color: #dbdbdb;
}

.calendar-days {
  margin: 10px;
}

ol,
li {
  padding: 0 !important;
  margin: 0;
  list-style: none;
}
</style>

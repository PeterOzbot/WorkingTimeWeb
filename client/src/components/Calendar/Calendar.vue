<template>
  <div class="calendar">
    <v-card class="calendar-month" elevation="3">
      <div class="calendar-month-header">
        <MonthHeaderComponent
          :currentMonth="currentMonth"
          :currentYear="currentYear"
        />
      </div>
      <WeekdaysComponent />
    </v-card>
    <v-card class="calendar-days" elevation="5">
      <ol class="days-grid">
        <DayComponent
          v-for="(day, dayIndex) in workingDays"
          :key="dayIndex"
          :day="day"
        />
      </ol>
    </v-card>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { Component, Prop } from "vue-property-decorator";
import WeekdaysComponent from "./Weekdays.vue";
import MonthHeaderComponent from "./MonthHeader.vue";
import DayComponent from "./Day.vue";
import WorkingDay from "@/models/workingDay";

@Component({
  components: { WeekdaysComponent, MonthHeaderComponent, DayComponent }
})
export default class CalendarView extends Vue {
  @Prop()
  public generatedWorkingDays: WorkingDay[];
  @Prop()
  public currentMonth: number;
  @Prop()
  public currentYear: number;

  public workingDays: WorkingDay[] = [];

  created() {
    // when generated days change prepare calendar
    this.$watch("generatedWorkingDays", () => this.prepareCalendar());
  }

  private prepareCalendar() {
    this.workingDays = [
      ...this.getDaysInPreviousMonth(),
      ...this.generatedWorkingDays,
      ...this.getDaysInNextMonth()
    ];
  }

  private getDaysInPreviousMonth() {
    const days: WorkingDay[] = [];
    // create first of the month
    const date = new Date(this.currentYear, this.currentMonth, 1);

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
      // reduce ramining days
      previousMonthDays--;
    }

    return days;
  }

  private getDaysInNextMonth() {
    const days: WorkingDay[] = [];

    // create last day of the month
    const date = new Date(this.currentYear, this.currentMonth + 1, 0);

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
      // reduce ramining days
      nextMonthDays--;
    }

    return days;
  }
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

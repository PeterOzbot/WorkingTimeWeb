<template>
  <div class="calendar">
    <v-card class="calendar-month" elevation="5">
      <div class="calendar-month-header">
        <MonthHeaderComponent v-bind:currentMonth="currentMonth" />
      </div>
      <WeekdaysComponent />
    </v-card>
    <v-card class="calendar-days" elevation="5">
      <ol class="days-grid">
        <DayComponenet v-for="day in workingDays" :key="day.date" :day="day" />
      </ol>
    </v-card>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import { Component, Prop } from "vue-property-decorator";
import WeekdaysComponent from "@/components/Calendar/Weekdays.vue";
import MonthHeaderComponent from "@/components/Calendar/MonthHeader.vue";
import DayComponent from "@/components/Calendar/Day.vue";
import WorkingDay from "@/models/workingDay";

@Component({
  components: { WeekdaysComponent, MonthHeaderComponent, DayComponent }
})
export default class CalendarView extends Vue {
  public currentMonth: Date = new Date();
  @Prop()
  public workingDays: WorkingDay[] = [];
}
</script>

<style scoped>
.calendar-month {
  position: relative;
  background-color: #fafaf9;
  margin: 10px;
}
.calendar-month-header {
  display: flex;
  justify-content: center;
  background-color: #fafaf9;
  padding: 10px;
  border-color: gainsboro;
  border-width: 0px 0px 1px 0px;
  border-style: solid;
}
.days-grid {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  height: 100%;
  position: relative;
  grid-column-gap: 1px;
  grid-row-gap: 1px;
  border-top: solid 1px #cfd7e3;
}
</style>

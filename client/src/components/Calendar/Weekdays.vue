<template>
  <ol class="day-of-week">
    <li v-for="weekday in Weekdays" :key="weekday">
      <div class="weekday">
        {{ weekday }}
      </div>
    </li>
  </ol>
</template>

<script lang="ts">
import Vue from "vue";
import { Component } from "vue-property-decorator";

@Component
export default class WeekdaysComponent extends Vue {
  public Weekdays = [
    "Monday",
    "Tuesday",
    "Wednesday",
    "Thursday",
    "Friday",
    "Saturday",
    "Sunday"
  ];
}
</script>

<style scoped>
ol,
li {
  padding: 0;
  margin: 0;
  list-style: none;
}

.day-of-week {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  padding: 10px;
}

.day-of-week > * {
  text-align: center;
  padding-right: 5px;
}
</style>

<template>
  <li
    class="calendar-day"
    :class="{
      holiday: isHoliday
    }"
  >
    <div
      class="day-container"
      :class="{ 'irrelevantMonth-day': day.isIrrelevantMonth }"
    >
      <div class="day-hours" :class="{ 'not-working-day': day.hours == 0 }">
        <InputComponent
          v-if="!day.isIrrelevantMonth"
          :model.sync="day.hours"
          after="h"
        />
      </div>
      <div class="day-number">
        <span>{{ dayLabel }}</span>
      </div>
    </div>
  </li>
</template>

<script lang="ts">
import Vue from "vue";
import { Component, Prop } from "vue-property-decorator";
import WorkingDay from "@/models/workingDay";
import InputComponent from "@/components/Controls/Input.vue";

@Component({
  components: { InputComponent }
})
export default class DayComponent extends Vue {
  @Prop()
  day: WorkingDay;

  get dayLabel(): string {
    if (!this.day.isIrrelevantMonth) {
      return `${this.day.date.getDate()}.`;
    } else {
      return `${this.day.date.getDate()}.${this.day.date.getMonth() + 1}`;
    }
  }

  get isHoliday(): boolean {
    return this.day.date.getDay() == 1;
  }
}
</script>

<style scoped lang="scss">
.calendar-day {
  min-height: 120px;
}

.irrelevantMonth-day {
  opacity: 0.5;

  & > * {
    color: #9d9d9d;
  }
}

.day-container {
  display: flex;
  flex-flow: column;
  height: 100%;
  background-color: white;

  .day-hours {
    flex: 1 0 auto;
    font-size: 20px;
  }

  .not-working-day {
    opacity: 0.5;
  }

  .day-number {
    flex: 0 0 auto;
    background-color: #f4f4f4;
    text-align: right;
    padding-bottom: 5px;
    padding-right: 20px;
    padding-top: 5px;

    span {
      font-size: 12px;
      font-weight: 500;
    }
  }
}

ol,
li {
  padding: 0;
  margin: 0;
  list-style: none;
}
</style>

<template>
  <li class="calendar-day" :class="{ holiday: isHoliday() }">
    <div class="day-container" :class="{ 'irrelevantMonth-day': day.isIrrelevantMonth }">
      <div class="day-hours" :class="{ 'not-working-day': day.hours == 0 }">
        <InputComponent v-if="!day.isIrrelevantMonth" v-model:value="day.hours" after="h" />
      </div>
      <div class="day-number">
        <span>{{ dayLabel() }}</span>
      </div>
    </div>
  </li>
</template>

<script lang="ts" setup>
import type WorkingDay from "@/models/workingDay";
import InputComponent from "@/components/controls/Input.vue";
import type { PropType } from "vue";

const props = defineProps({
  day: { type: Object as PropType<WorkingDay>, required: true },
});

//const day = reactive(props.day);

function dayLabel(): string {
  if (!props.day.isIrrelevantMonth) {
    return `${props.day.date.getDate()}.`;
  } else {
    return `${props.day.date.getDate()}.${props.day.date.getMonth() + 1}`;
  }
}

function isHoliday(): boolean {
  return props.day.date.getDay() == 1;
}

</script>

<style scoped lang="scss">
.calendar-day {
  min-height: 120px;
}

.irrelevantMonth-day {
  opacity: 0.5;

  &>* {
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

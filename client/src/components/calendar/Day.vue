<template>
  <li class="calendar-day" :class="{ holiday: isHoliday() }">
    <div class="day-container" :class="{ 'irrelevantMonth-day': day.isIrrelevantMonth }">
      <div class="day-hours" :class="{ 'not-working-day': day.hours == 0 }" @click="focusInputElement()">
        <InputComponent ref="inputElement" v-if="!day.isIrrelevantMonth" v-model="day.hours" after="h" />
      </div>
      <div class="day-footer">
        <div class="day-group">
          <GroupInput v-if="!day.isIrrelevantMonth && day.hours > 0" v-model="day.groupId" />
        </div>
        <div class="day-number">
          <span>{{ dayLabel() }}</span>
        </div>
      </div>
    </div>
  </li>
</template>

<script lang="ts" setup>
import type EditableWorkingDay from "@/models/editableWorkingDay";
import InputComponent from "@/components/controls/HoursInput.vue";
import GroupInput from "@/components/controls/GroupInput.vue";
import { ref, type PropType } from "vue";

const props = defineProps({
  day: { type: Object as PropType<EditableWorkingDay>, required: true },
});

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

const inputElement = ref<InstanceType<typeof InputComponent>>();
function focusInputElement() {
  inputElement.value?.focus();
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
    display: flex;
    flex-wrap: wrap;
    align-content: center;
    flex: 1 0 auto;
    font-size: 20px;
  }

  .not-working-day {
    opacity: 0.5;
  }

  .day-footer {
    background-color: #f4f4f4;
    padding-bottom: 5px;
    padding-right: 20px;
    padding-top: 5px;
    display: flex;
  }

  .day-number {
    flex: 1 0 auto;
    text-align: right;

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

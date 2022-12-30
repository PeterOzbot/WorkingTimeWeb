<template>
  <div class="calendar-date-indicator">{{ formatCurrentMonth() }}</div>
</template>

<script lang="ts" setup>

import moment from "moment";

const props = defineProps({
  currentMonth: { type: Number, required: true },
  currentYear: { type: Number, required: true },
});

function formatCurrentMonth(): string {
  return moment(new Date(props.currentYear ?? 0, props.currentMonth, 1)).format(
    "MMMM YYYY"
  );
}

</script>

<style scoped>
.calendar-date-indicator {
  font-size: 28px;
  padding: 10px;
}
</style>

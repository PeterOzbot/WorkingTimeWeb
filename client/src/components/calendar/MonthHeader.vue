<template>
  <div class="calendar-month-header">
    <div class="calendar-date-indicator">{{ formatCurrentMonth() }}</div>

    <v-card class="download-button" @click="emit('download')" elevation="6">
      Download Excel
    </v-card>
  </div>
</template>

<script lang="ts" setup>

import moment from "moment";

const emit = defineEmits(['download'])

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

.calendar-month-header {
  display: flex;
  justify-content: center;
  background-color: #94cb94;
  padding: 10px;
}

.download-button {
  position: absolute;
  right: 0;
  margin-right: 20px;
  padding: 10px;
  background-color: #67b567;
  font-size: 25px;
}
</style>

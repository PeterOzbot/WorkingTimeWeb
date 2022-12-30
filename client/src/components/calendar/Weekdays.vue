<template>
  <ol class="week">
    <li v-for="weekday in Weekdays" :key="weekday">
      <div class="weekday">
        {{ weekday }}
      </div>
    </li>
  </ol>
</template>

<script lang="ts" setup>

const Weekdays = [
  "Monday",
  "Tuesday",
  "Wednesday",
  "Thursday",
  "Friday",
  "Saturday",
  "Sunday"
];

</script>

<style scoped>
ol,
li {
  padding: 0;
  margin: 0;
  list-style: none;
}

.week {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
}

.week li:last-child .weekday {
  border: 0px;
}

.weekday {
  text-align: center;
  padding: 15px;
  background-color: #67b567;
  border-right: 1px solid #dbdbdb;
}
</style>

<template>
    <div class="months-container">
        <v-card v-for="(month, monthIndex) in Months" :key="monthIndex" :class="{ 'selected': month.selected }"
            class="month-container" @click="selectMonth(month)">
            {{ month.name }}
        </v-card>
    </div>
</template>

<script lang="ts" setup>
import Month from "@/models/month";
import { ref } from "vue";

const emit = defineEmits(['update:modelValue'])

const Months = ref(Array<Month>());
initializeMonths();

function initializeMonths() {
    Months.value.push(new Month("January", 1));
    Months.value.push(new Month("February", 2));
    Months.value.push(new Month("March", 3));
    Months.value.push(new Month("April", 4));
    Months.value.push(new Month("May", 5));
    Months.value.push(new Month("June", 6));
    Months.value.push(new Month("July", 7));
    Months.value.push(new Month("August", 8));
    Months.value.push(new Month("September", 9));
    Months.value.push(new Month("October", 10));
    Months.value.push(new Month("November", 11));
    Months.value.push(new Month("December", 12));

    selectMonth(Months.value[new Date().getMonth()]);
}

function selectMonth(month: Month) {
    Months.value.forEach(month => month.selected = false);
    month.selected = !month.selected;
    emit('update:modelValue', month.id)
}

</script>

<style scoped lang="scss">
.months-container {
    font-size: 30px;
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    max-height: 450px;
    min-width: 400px;
    overflow: auto;

    .month-container {
        flex: 0 1 170px;
        margin: 10px;
        padding: 8px;
        background-color: #dddddd;
        text-align: center;

        border-color: #dddddd;
        border-width: 2px;
    }

    .month-container.selected {
        border-color: #67b567;
        border-width: 2px;
    }
}
</style>
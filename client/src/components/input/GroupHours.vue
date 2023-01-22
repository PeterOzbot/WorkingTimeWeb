<template>
    <div class="group-hours-container">
        <div class="hours-container">
            <div class="column">
                <div class="input-label">A - Hours:</div>
            </div>
            <div class="column">
                <InputComponent class="hours-input" v-model="a_hours" after="h" />
            </div>
        </div>
        <div class="hours-container">
            <div class="column">
                <div class="input-label">B - Hours:</div>
            </div>
            <div class="column">
                <InputComponent class="hours-input" v-model="b_hours" after="h" />
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import InputComponent from "@/components/controls/HoursInput.vue";
import { computed, ref } from "vue";

const emit = defineEmits(['update:modelValue'])
// const props = defineProps({
//     modelValue: {
//         type: Array<Number>,
//         default: () => []
//     }
// })

const a_hoursRef = ref(0)
const b_hoursRef = ref(0)

const a_hours = computed({
    get() {
        return a_hoursRef.value
    },
    set(val) {
        a_hoursRef.value = val
        emit("update:modelValue", [a_hours.value, b_hours.value]);
    }
})

const b_hours = computed({
    get() {
        return b_hoursRef.value
    },
    set(val) {
        b_hoursRef.value = val
        emit("update:modelValue", [a_hours.value, b_hours.value]);
    }
})

</script>

<style scoped lang="scss">
.hours-container {
    flex: 0 0 auto;

    width: 100%;
    flex-wrap: wrap;
    display: flex;
    flex-flow: row;

    background-color: #67b567;

    .input-label {
        margin: 20px;
        margin-right: 50px;
        font-size: 30px;
        text-align: right;
    }

    .hours-input {
        font-size: 30px;
        margin-left: 60px;
        margin-right: 60px;
    }
}
</style>
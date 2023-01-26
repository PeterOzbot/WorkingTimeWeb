<template>
    <div class="group-hours-container">
        <v-card class="add-group-button" @click="addGroup()" elevation="6">
            Add
        </v-card>
        <div class="hours-container" v-for="group in props.modelValue">
            <GroupInput class="input-label" v-model="group.groupId" />
            <HoursInput class="hours-input" v-model="group.hours" after="h" />
        </div>
    </div>
</template>

<script lang="ts" setup>
import HoursInput from "@/components/controls/HoursInput.vue";
import GroupInput from "@/components/controls/GroupInput.vue";
import Group from "@/models/group";

const props = defineProps({
    modelValue: {
        type: Array<Group>,
        default: []
    }
});

function addGroup() {
    const newGroupId = (props.modelValue.length + 1).toString();
    props.modelValue.push(new Group(newGroupId, 0));
}

</script>

<style scoped lang="scss">
.group-hours-container {
    display: flex;
    flex-flow: column;
    background-color: #67b567;
}

.add-group-button {
    width: 50px;
    margin: 5px;
    background-color: #67b667;
    text-align: center;
}

.hours-container {
    width: 100%;
    display: flex;
    flex-flow: row;

    .input-label {
        width: 100%;
        margin: 20px;
        font-size: 30px;
        text-align: right;
    }

    .hours-input {
        width: 100%;
        font-size: 30px;
        margin: 20px;
    }
}
</style>
import type EditableWorkingDay from "./editableWorkingDay";

export default class CreateRequest {
    public days: EditableWorkingDay[];
    public groupId: string;

    constructor(groupId: string, days: EditableWorkingDay[]) {
        this.groupId = groupId;
        this.days = [...days];
    }
}

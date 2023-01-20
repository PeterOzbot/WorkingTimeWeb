import WorkingDay from "./workingDay";

export default class EditableWorkingDay extends WorkingDay {
    public groupId: string;
    public isIrrelevantMonth: boolean;

    constructor(date: Date, hours: number, isIrrelevantMonth: boolean) {
        super(date, hours)
        this.date = date;
        this.hours = hours;
        this.isIrrelevantMonth = isIrrelevantMonth;
        this.groupId = "Group 1";
    }
}

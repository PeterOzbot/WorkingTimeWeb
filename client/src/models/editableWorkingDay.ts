import WorkingDay from "./workingDay";

export default class EditableWorkingDay extends WorkingDay {
    public isIrrelevantMonth: boolean;

    constructor(date: Date, hours: number, isIrrelevantMonth: boolean, groupId: string) {
        super(date, hours, groupId)
        this.date = date;
        this.hours = hours;
        this.isIrrelevantMonth = isIrrelevantMonth;
    }
}

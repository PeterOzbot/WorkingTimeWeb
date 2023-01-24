export default class Group {
    public groupId: string;
    public hours: number;

    constructor(groupId: string, hours: number) {
        this.groupId = groupId;
        this.hours = hours;
    }
}
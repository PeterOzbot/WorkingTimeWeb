export default class WorkingDay {
  public date: Date;
  public hours: number;
  public groupId: string;

  constructor(date: Date, hours: number, groupId: string) {
    this.date = date;
    this.hours = hours;
    this.groupId = groupId;
  }
}

export default class WorkingDay {
  public date: Date;
  public hours: number;
  public isIrrelevantMonth: boolean;

  constructor(date: Date, hours: number, isIrrelevantMonth: boolean) {
    this.date = date;
    this.hours = hours;
    this.isIrrelevantMonth = isIrrelevantMonth;
  }
}

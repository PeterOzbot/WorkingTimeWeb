import { Injectable } from "@dvolper/tsdi";

@Injectable
export class NotificationService {
  public show = false;
  public currentMessage = "";
  public timeout = 2000;

  public showNotification(message: string) {
    this.show = true;
    this.currentMessage = message;
  }

  public close() {
    this.show = false;
  }
}

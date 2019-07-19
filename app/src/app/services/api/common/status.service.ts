import { Injectable } from "@angular/core";
import { BehaviorSubject, Subscription } from "rxjs";
import { Status } from "./status";

@Injectable()
export class StatusService {

    private subject: BehaviorSubject<Status | undefined>

    constructor() {
        this.subject = new BehaviorSubject<Status | undefined>(undefined)
    }

    update(status: Status): void {
        this.subject.next(status)
    }

    onChanged(fn: (status: Status | undefined) => void): Subscription {
        return this.subject.subscribe(fn)
    }

}
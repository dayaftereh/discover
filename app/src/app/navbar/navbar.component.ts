import { Component, OnInit, OnDestroy } from "@angular/core";
import { MenuItem } from "primeng/components/common/menuitem";
import { StatusService } from "../services/api/common/status.service";
import { Status } from "../services/api/common/status";
import { Subscription } from "rxjs";
import { SelectItem } from "primeng/components/common/selectitem";

@Component({
    selector: 'app-navbar',
    templateUrl: './navbar.component.html'
})
export class NavbarComponent implements OnInit, OnDestroy {

    items: MenuItem[] | undefined

    status: Status | undefined

    private subscription: Subscription | undefined

    constructor(private readonly statusService: StatusService) {

    }

    ngOnInit(): void {
        this.subscription = this.statusService.onChanged((status: Status | undefined) => {
            this.status = status
            if (status) {
                this.createMenu(status)
            }
        })
    }

    private createMenu(status: Status): void {
        this.items = [
            { label: status.name },
            { separator: true },
            { label: 'Settings' },
            { separator: true },
            {
                label: 'Admin',
                routerLink: ['admin'],
                items: [
                    {
                        label: 'Star-Systems',
                        routerLink: ['admin', 'star-systems'],
                    }
                ]
            }
        ]
    }

    ngOnDestroy(): void {
        if (this.subscription) {
            this.subscription.unsubscribe()
        }
    }

}

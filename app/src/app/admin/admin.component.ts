import { Component, OnInit } from "@angular/core";
import { MenuItem } from "primeng/components/common/menuitem";

@Component({
    templateUrl: './admin.component.html'
})
export class AdminComponent implements OnInit {

    items: MenuItem[]

    constructor() {
        this.items = []
    }

    ngOnInit(): void {
        this.items.push({
            label: 'Star-Systems'
        })
    }

}

import { Component } from "@angular/core";
import { MenuItem } from "primeng/components/common/menuitem";

@Component({
    selector: 'app-navbar',
    templateUrl: './navbar.component.html'
})
export class NavbarComponent {

    items: MenuItem[] | undefined

    constructor() { }

}
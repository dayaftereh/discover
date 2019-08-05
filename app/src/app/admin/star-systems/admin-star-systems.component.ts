import { Component, OnInit } from "@angular/core";
import { AdminService } from "src/app/services/api/admin/admin.service";

@Component({
    templateUrl: './admin-star-systems.component.html'
})
export class AdminStarSystemsComponent implements OnInit {

    starSystems: string[]

    constructor(private readonly adminService: AdminService) {
        this.starSystems = []
    }

    async ngOnInit(): Promise<void> {
        this.starSystems = await this.adminService.allStarSystems()
    }

}
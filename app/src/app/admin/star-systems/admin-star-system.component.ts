import { Component, OnInit } from "@angular/core";
import { StarSystem } from "src/app/services/api/admin/star-system";
import { AdminService } from "src/app/services/api/admin/admin.service";

@Component({
    templateUrl: "./admin-star-system.component.ts"
})
export class AdminStarSystemComponent implements OnInit {

    starSystem: StarSystem | undefined

    constructor(private readonly adminService: AdminService) {

    }

    async ngOnInit(): Promise<void> { 
        
    }

}
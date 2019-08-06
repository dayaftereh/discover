import { Component, OnInit } from "@angular/core";
import { ActivatedRoute, ParamMap } from "@angular/router";
import { first } from 'rxjs/operators';
import { AdminService } from "src/app/services/api/admin/admin.service";
import { StarSystem } from "src/app/services/api/admin/star-system";

@Component({
    templateUrl: "./admin-star-system.component.html"
})
export class AdminStarSystemComponent implements OnInit {

    starSystem: StarSystem | undefined

    constructor(private readonly adminService: AdminService,
        private readonly activatedRoute: ActivatedRoute) {

    }

    async ngOnInit(): Promise<void> {
        const paramMap: ParamMap = await this.activatedRoute.paramMap.pipe(first()).toPromise()
        const name: string | null = paramMap.get("name")

        if (name) {
            this.starSystem = await this.adminService.starSystem(name)
            console.log(this.starSystem)
        }
    }

    color(x: number): string {
        if (!x) {
            return `#FFFFFF`
        }
        return `#${x.toString(16)}`
    }

    solarMass2KG(x: number): number {
        if (!x) {
            return 0.0
        }
        return x * 1.98847e30
    }

    au2KM(x: number): number {
        if (!x) {
            return 0.0
        }
        // 1 AU = 149597870700 m
        return x * 149597870.700
    }

}
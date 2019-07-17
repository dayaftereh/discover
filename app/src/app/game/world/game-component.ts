export interface GameComponent {

    init(): void
    update(delta: number): void
    dispose(): void

}
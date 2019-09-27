import {FastifyInstance} from 'fastify';

export class AppModule {

    static footRoot(fastify: FastifyInstance, options: any, done: any): void {
        const app = new AppModule(fastify);
        app.setProviders();
        app.onInit();
        app.setRoute();
        app.onChange();
        done();
    }

    constructor(
        private fastify: FastifyInstance,
    ) {
    }

    /**
     * Set Providers
     */
    setProviders() {
    }

    /**
     * Init
     */
    onInit() {
    }

    /**
     * Set Route
     */
    setRoute() {
    }

    /**
     * On Event Change
     */
    onChange() {
    }
}

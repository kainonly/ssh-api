import { FastifyInstance } from 'fastify';
import { api } from './router/api';
import { ClientService } from './common/client.service';

export class AppModule {
  private client: ClientService;

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
    this.client = new ClientService();
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
    api(this.fastify, this.client);
  }

  /**
   * On Event Change
   */
  onChange() {
  }
}

import { FastifyInstance } from 'fastify';
import { join } from 'path';
import { api } from './router/api';
import { ClientService } from './common/client.service';
import { ConfigService } from './common/config.service';

export class AppModule {
  private config: ConfigService;
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
    this.config = new ConfigService(join(__dirname, 'config.json'));
    this.client = new ClientService();
  }

  /**
   * Init
   */
  onInit() {
    const configs: any = this.config.get();
    for (const key in configs) {
      if (configs.hasOwnProperty(key)) {
        const data = configs[key];
        data['privateKey'] = Buffer.from(data['privateKey'], 'base64');
        this.client.put(key, data);
      }
    }
  }

  /**
   * Set Route
   */
  setRoute() {
    api(this.fastify, this.client, this.config);
  }

  /**
   * On Event Change
   */
  onChange() {
  }
}

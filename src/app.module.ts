import { FastifyInstance } from 'fastify';
import { join } from 'path';
import { ClientService } from './common/client.service';
import { ConfigService } from './common/config.service';
import { api } from './router/api';

export class AppModule {
  private config: ConfigService;
  private client: ClientService;

  static footRoot(fastify: FastifyInstance, options: any, done: any): void {
    const app = new AppModule(fastify);
    app.setProviders();
    app.onInit();
    app.setRoute();
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
    this.config = new ConfigService(join(__dirname, 'data', 'config.json'));
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
        this.client.put(key, {
          host: data.host,
          port: data.port,
          username: data.username,
          password: data.password,
          privateKey: Buffer.from(data.privateKey, 'base64'),
          passphrase: data.passphrase,
        });
        this.client.tunnel(key, data.tunnels);
      }
    }
  }

  /**
   * Set Route
   */
  setRoute() {
    api(this.fastify, this.client, this.config);
  }
}

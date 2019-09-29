import { ClientOption } from '../types/client-option';
import { Client, ConnectConfig } from 'ssh2';
import { Runtime } from '../types/runtime';
import { Stream } from 'stream';

export class ClientService {
  private clientOption: ClientOption | {} = {};
  private runtime: Runtime | {} = {};

  /**
   * Connection
   * @param config
   */
  private connection(config: ConnectConfig): Promise<Client> {
    return new Promise((resolve, reject) => {
      let client = new Client();
      client.connect(config);
      client.on('ready', () => {
        resolve(client);
      });
      client.on('error', error => {
        reject(error);
      });
      client.on('close', () => {
        client.removeAllListeners();
        client = undefined;
      });
    });
  }

  /**
   * Connect testing ssh client
   * @param config
   */
  async testing(config: ConnectConfig): Promise<boolean> {
    try {
      const client = await this.connection(config);
      client.destroy();
      return true;
    } catch (e) {
      console.error(e.message);
      return false;
    }
  }

  /**
   * Put a ssh client
   * @param identity
   * @param config
   */
  put(identity: string, config: ConnectConfig) {
    return Reflect.set(this.clientOption, identity, config);
  }

  /**
   * Remote exec
   * @param identity
   * @param bash
   */
  async exec(identity: string, bash: string): Promise<Stream> {
    return new Promise(async (resolve, reject) => {
      try {
        if (!this.clientOption.hasOwnProperty(identity)) {
          reject('client not exists');
        }
        const client = this.runtime[identity] = await this.connection(this.clientOption[identity]);
        client.exec(bash, (err, channel) => {
          resolve(channel);
        });
      } catch (e) {
        reject(e.message);
      }
    });
  }
}

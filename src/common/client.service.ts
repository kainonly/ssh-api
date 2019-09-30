import { Stream } from 'stream';
import { Client, ConnectConfig } from 'ssh2';
import { Server } from 'net';

export class ClientService {
  private clientOption: Map<string, ConnectConfig> = new Map<string, ConnectConfig>();
  private clientRuntime: Map<string, Client> = new Map<string, Client>();
  private clientStatus: Map<string, boolean> = new Map<string, boolean>();
  private serverOption: Map<string, any[]> = new Map<string, any[]>();
  private serverRuntime: Map<string, Server> = new Map<string, Server>();

  /**
   * Get Client Option
   */
  getClientOption() {
    return this.clientOption;
  }

  /**
   * Connect testing ssh client
   * @param config
   */
  async testing(config: ConnectConfig): Promise<any> {
    return new Promise((resolve, reject) => {
      let client = new Client();
      client.connect(config);
      client.on('ready', () => {
        resolve('ok');
        client.destroy();
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
   * Connection
   * @param identity
   */
  private connection(identity: string): Promise<Client> {
    return new Promise((resolve, reject) => {
      if (!this.clientOption.has(identity)) {
        reject('client not exists');
      }
      const client = new Client();
      client.connect(this.clientOption.get(identity));
      client.on('ready', () => {
        this.clientStatus.set(identity, true);
        resolve(client);
      });
      client.on('error', error => {
        reject(error);
      });
      client.on('close', () => {
        client.removeAllListeners();
        this.clientStatus.set(identity, false);
      });
      this.clientRuntime.set(identity, client);
    });
  }

  /**
   * Get ssh client
   * @param identity
   */
  get(identity: string) {
    if (!this.clientOption.has(identity)) {
      return null;
    }
    const option = this.clientOption.get(identity);
    return {
      identity,
      host: option.host,
      port: option.port,
      username: option.username,
      connected: this.clientStatus.get(identity),
    };
  }

  /**
   * Put a ssh client
   * @param identity
   * @param config
   * @return boolean
   */
  put(identity: string, config: ConnectConfig): boolean {
    try {
      this.close(identity);
      this.clientOption.set(identity, config);
      this.clientStatus.set(identity, false);
      return true;
    } catch (e) {
      return false;
    }
  }

  /**
   * Remote exec
   * @param identity
   * @param bash
   */
  exec(identity: string, bash: string): Promise<Stream> {
    return new Promise(async (resolve, reject) => {
      try {
        if (!this.clientOption.has(identity)) {
          reject('client not exists');
        }
        let client: Client;
        if (!this.clientRuntime.has(identity)) {
          client = await this.connection(identity);
        } else {
          client = this.clientRuntime.get(identity);
        }
        client.exec(bash, (err, channel) => {
          this.clientStatus.set(identity, true);
          resolve(channel);
        });
      } catch (e) {
        reject(e.message);
      }
    });
  }

  /**
   * Close ssh Client
   * @param identity
   */
  close(identity: string) {
    if (this.clientRuntime.has(identity)) {
      this.clientRuntime.get(identity).destroy();
    }
    return this.clientRuntime.delete(identity);
  }

  /**
   * Delete ssh client
   * @param identity
   */
  delete(identity: string): boolean {
    return (
      this.close(identity) &&
      this.clientOption.delete(identity) &&
      this.clientStatus.delete(identity)
    );
  }

  tunnel() {

  }
}

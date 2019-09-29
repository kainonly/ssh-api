import { ClientOption } from '../types/client-option';
import { Client, ConnectConfig } from 'ssh2';

export class ClientService {
  private client: ClientOption;

  testing(config: ConnectConfig): Promise<any> {
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
}

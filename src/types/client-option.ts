import { ConnectConfig } from 'ssh2';

export interface ClientOption {
  identity: string;
  config: ConnectConfig;
}

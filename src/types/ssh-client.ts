import { ConnectConfig } from 'ssh2';

export interface SshClient {
  identity: string;
  config: ConnectConfig;
}

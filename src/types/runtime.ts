import { Client } from 'ssh2';

export interface Runtime {
  [identity: string]: Client;
}

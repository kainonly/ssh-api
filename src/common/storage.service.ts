import { FastifyInstance } from 'fastify';
import * as PouchDB from 'pouchdb';

export class StorageService {
  private database: PouchDB.Database;

  constructor(
    private readonly fastify: FastifyInstance,
  ) {
    this.database = fastify.pouchdb;
  }

  /**
   * Get Data
   * @param key
   */
  async get(key: string) {
    try {
      return await this.database.get(key);
    } catch (e) {
      if (e.message === 'missing') {
        return null;
      } else {
        return false;
      }
    }
  }

  /**
   * Add Data
   * @param key
   * @param value
   */
  async add(key: string, value: any) {
    try {
      const doc = await this.database.get(key);
      const data = Object.assign({
        _id: key,
        _rev: doc._rev,
      }, value);
      return await this.database.put(data);
    } catch (e) {
      if (e.message === 'missing') {
        return await this.database.put(Object.assign({
          _id: key,
        }, value));
      } else {
        return false;
      }
    }
  }
}

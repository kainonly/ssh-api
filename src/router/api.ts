import { FastifyInstance } from 'fastify';
import { ClientService } from '../common/client.service';
import { ConfigService } from '../common/config.service';

const api = (fastify: FastifyInstance, client: ClientService, config: ConfigService) => {
  /**
   * temporary
   */
  function temporary() {
    const data = {};
    const serverOption = client.getServerOption();
    client.getClientOption().forEach((value, key) => {
      data[key] = {
        host: value.host,
        port: value.port,
        username: value.username,
        password: value.password,
        privateKey: value.privateKey.toString('base64'),
        passphrase: value.passphrase,
        tunnels: !serverOption.has(key) ? [] : serverOption.get(key),
      };
    });
    config.set(data);
  }

  /**
   * get all identity
   */
  fastify.post('/all', async (request, reply) => {
    reply.send({
      error: 0,
      data: [...client.getClientOption().keys()],
    });
  });
  /**
   * Testing a ssh client
   */
  fastify.post('/testing', {
    schema: {
      body: {
        required: ['host', 'port', 'username'],
        properties: {
          host: {
            type: 'string',
          },
          port: {
            type: 'number',
          },
          username: {
            type: 'string',
          },
          password: {
            type: 'string',
          },
          private_key: {
            type: 'string',
          },
          passphrase: {
            type: 'string',
          },
        },
        oneOf: [
          {
            required: ['password'],
          },
          {
            required: ['private_key'],
          },
        ],
      },
    },
  }, async (request, reply) => {
    try {
      const body = request.body;
      if (body.private_key) {
        body.private_key = Buffer.from(body.private_key, 'base64');
      }
      const result = await client.testing({
        host: body.host,
        port: body.port,
        username: body.username,
        password: body.password,
        privateKey: body.private_key,
        passphrase: body.passphrase,
      });
      reply.send({
        error: 0,
        msg: result,
      });
    } catch (e) {
      reply.send({
        error: 1,
        msg: e.message,
      });
    }
  });
  /**
   * Lists all ssh client status
   */
  fastify.post('/lists', {
    schema: {
      body: {
        required: ['identity'],
        properties: {
          identity: {
            type: 'array',
          },
        },
      },
    },
  }, (request, reply) => {
    const body = request.body;
    reply.send({
      error: 0,
      data: body.identity.map(v => client.get(v)),
    });
  });
  /**
   * Get a ssh client status
   */
  fastify.post('/get', {
    schema: {
      body: {
        required: ['identity'],
        properties: {
          identity: {
            type: 'string',
          },
        },
      },
    },
  }, (request, reply) => {
    const body = request.body;
    reply.send({
      error: 0,
      data: client.get(body.identity),
    });
  });
  /**
   * Put a ssh client
   */
  fastify.post('/put', {
    schema: {
      body: {
        required: ['identity', 'host', 'port', 'username'],
        properties: {
          identity: {
            type: 'string',
          },
          host: {
            type: 'string',
          },
          port: {
            type: 'number',
          },
          username: {
            type: 'string',
          },
          password: {
            type: 'string',
          },
          private_key: {
            type: 'string',
          },
          passphrase: {
            type: 'string',
          },
        },
        oneOf: [
          {
            required: ['password'],
          },
          {
            required: ['private_key'],
          },
        ],
      },
    },
  }, async (request, reply) => {
    try {
      const body = request.body;
      if (body.private_key) {
        body.private_key = Buffer.from(body.private_key, 'base64');
      }
      const result: boolean = client.put(body.identity, {
        host: body.host,
        port: body.port,
        username: body.username,
        password: body.password,
        privateKey: body.private_key,
        passphrase: body.passphrase,
      });
      temporary();
      reply.send(result ? {
        error: 0,
        msg: 'ok',
      } : {
        error: 1,
        msg: 'failed',
      });
    } catch (e) {
      reply.send({
        error: 1,
        msg: e.message,
      });
    }
  });
  /**
   * A ssh client to exec
   */
  fastify.post('/exec', {
    schema: {
      body: {
        required: ['identity', 'bash'],
        properties: {
          identity: {
            type: 'string',
          },
          bash: {
            type: 'string',
          },
        },
      },
    },
  }, async (request, reply) => {
    try {
      const body = request.body;
      const stream = await client.exec(body.identity, body.bash);
      reply.send(stream);
    } catch (e) {
      reply.send({
        error: 1,
        msg: e,
      });
    }
  });
  /**
   * Delete a ssh client
   */
  fastify.post('/delete', {
    schema: {
      body: {
        required: ['identity'],
        properties: {
          identity: {
            type: 'string',
          },
        },
      },
    },
  }, (request, reply) => {
    const body = request.body;
    const result: boolean = client.delete(body.identity);
    reply.send(result ? {
      error: 0,
      msg: 'ok',
    } : {
      error: 1,
      msg: 'failed',
    });
  });
  /**
   * Set tunnels of ssh client
   */
  fastify.post('/tunnels', {
    schema: {
      body: {
        required: ['identity', 'tunnels'],
        properties: {
          identity: {
            type: 'string',
          },
          tunnels: {
            type: 'array',
            items: {
              type: 'array',
              maxItems: 4,
              minItems: 4,
              items: [
                { type: 'string' },
                { type: 'number' },
                { type: 'string' },
                { type: 'number' },
              ],
            },
          },
        },
      },
    },
  }, async (request, reply) => {
    try {
      const body = request.body;
      const result = await client.tunnel(body.identity, body.tunnels);
      temporary();
      reply.send(result ? {
        error: 0,
        msg: 'ok',
      } : {
        error: 1,
        msg: 'failed',
      });
    } catch (e) {
      reply.send({
        error: 1,
        msg: e,
      });
    }
  });
};

export { api };

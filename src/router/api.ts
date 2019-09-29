import { FastifyInstance } from 'fastify';
import { ClientService } from '../common/client.service';

const api = (fastify: FastifyInstance, client: ClientService) => {
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
      const result: boolean = await client.testing({
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
   * Create a tunnel
   */
  fastify.post('/bind', {
    schema: {
      body: {
        required: ['identity', 'scr_ip', 'src_port', 'dst_port'],
        properties: {
          identity: {
            type: 'string',
          },
          src_ip: {
            type: 'string',
          },
          src_port: {
            type: 'number',
          },
          dst_port: {
            type: 'number',
          },
        },
      },
    },
  }, (request, reply) => {
    reply.send({ error: 0 });
  });
  /**
   * Delete a tunnel
   */
  fastify.post('/unbind', {
    schema: {
      body: {
        required: ['identity', 'dst_port'],
        properties: {
          identity: {
            type: 'string',
          },
          dst_port: {
            type: 'number',
          },
        },
      },
    },
  }, (request, reply) => {
    reply.send({ error: 0 });
  });
  /**
   * Get tunnels of ssh client
   */
  fastify.post('/listen', {
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
    reply.send({ error: 0 });
  });
};

export { api };

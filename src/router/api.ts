import { FastifyInstance } from 'fastify';
import { ClientService } from '../common/client.service';

const api = (fastify: FastifyInstance, client: ClientService) => {
  /**
   * Lists all ssh client
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
    reply.send({ error: 0 });
  });
  /**
   * Get a ssh client
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
    reply.send({ error: 0 });
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
  }, (request, reply) => {
    const body = request.body;
    console.log(body);
    reply.send({ error: 0 });
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
  }, (request, reply) => {
    reply.send({ error: 0 });
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
    reply.send({ error: 0 });
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
  }, (request, reply) => {
    reply.send({ error: 0 });
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

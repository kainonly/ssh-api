import { FastifyInstance } from 'fastify';

const api = (fastify: FastifyInstance) => {
  fastify.get('/', (request, reply) => {
    reply.send({ version: 1 });
  });
};

export { api };

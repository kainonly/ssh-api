import { FastifyInstance } from 'fastify';

const api = (fastify: FastifyInstance) => {
  /**
   * Lists all ssh client
   */
  fastify.post('/lists', (request, reply) => {
    reply.send({ error: 0 });
  });
  /**
   * Get a ssh client
   */
  fastify.post('/get', (request, reply) => {
    reply.send({ error: 0 });
  });
  /**
   * Testing a ssh client
   */
  fastify.post('/testing', (request, reply) => {
    reply.send({ error: 0 });
  });
  /**
   * Put a ssh client
   */
  fastify.post('/put', (request, reply) => {
    reply.send({ error: 0 });
  });
  /**
   * Delete a ssh client
   */
  fastify.post('/delete', (request, reply) => {
    reply.send({ error: 0 });
  });
  /**
   * A ssh client to exec
   */
  fastify.post('/exec', (request, reply) => {
    reply.send({ error: 0 });
  });
  /**
   * Create a tunnel
   */
  fastify.post('/bind', (request, reply) => {
    reply.send({ error: 0 });
  });
  /**
   * Delete a tunnel
   */
  fastify.post('/unbind', (request, reply) => {
    reply.send({ error: 0 });
  });
  /**
   * Get tunnels of ssh client
   */
  fastify.post('/listen', (request, reply) => {
    reply.send({ error: 0 });
  });
};

export { api };

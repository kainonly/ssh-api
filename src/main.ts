import * as fastify from 'fastify';
import * as fastifyCompress from 'fastify-compress';
import { env } from 'process';

import { AppModule } from './app.module';

const server: fastify.FastifyInstance = fastify({
  logger: true,
});
server.register(fastifyCompress);
server.register(AppModule.footRoot);
server.listen(3000, '0.0.0.0', (err, address) => {
  if (err) {
    server.log.error(err);
    process.exit(1);
  }
  server.log.info(`server listening on ${address}`);
});

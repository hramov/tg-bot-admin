import { NestFactory } from '@nestjs/core';
import { config } from 'dotenv';
import { NextFunction } from 'express';
import {AppModule} from "./app.module";
import { ASYNC_STORAGE, LOGGER } from "./common/constants";
import {CustomLoggerService} from "./common/logger/custom-logger.service";
import {DocumentBuilder, SwaggerModule} from "@nestjs/swagger";
import {ValidationPipe} from "@nestjs/common";
import {Fetch} from "../../Infrastructure/Fetch/Fetch";
import helmet from 'helmet';
import { v4 } from 'uuid';
import { Logger } from "../../Infrastructure/Logger/Logger";
import { NestExpressApplication } from "@nestjs/platform-express";

config({
  path: '.env.local',
});

async function bootstrap() {
  const logger = new Logger('App')
  const app = await NestFactory.create<NestExpressApplication>(AppModule, {
    logger: logger,
  });

  app.setGlobalPrefix(process.env.APP_GLOBAL_PREFIX);
  app.enableCors({
    origin: '*',
    allowedHeaders: '*',
    methods: '*',
    exposedHeaders: '*'
  });
  app.use((req: any, res: any, next: NextFunction) => {
    const asyncStorage = app.get(ASYNC_STORAGE);
    const traceId = req.headers['x-request-id'] || v4().toString();
    const store = new Map<string, string>().set('traceId', traceId);
    asyncStorage.run(store, () => {
      next();
    });
  });

  app.useGlobalPipes(new ValidationPipe());

  app.useLogger(app.get<CustomLoggerService>(LOGGER));

  app.use(helmet());

  Fetch.init()

  const config = new DocumentBuilder()
      .setTitle('GVC Projects')
      .setDescription('GVC Projects MVC')
      .setVersion('1.0.0')
      .build();
  const document = SwaggerModule.createDocument(app, config);
  SwaggerModule.setup('swagger', app, document);

  await app.listen(process.env.APP_PORT);
}

bootstrap().catch((err: unknown) => console.error(err))

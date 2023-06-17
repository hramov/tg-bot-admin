import { NestFactory } from '@nestjs/core';
import { config } from 'dotenv';
import { NextFunction } from 'express';
import {AppModule} from "./app.module";
import {Logger} from "../../Infrastructure/Logger/Logger";
import {ASYNC_STORAGE} from "./common/constants";
import {CustomLoggerService} from "./common/logger/custom-logger.service";
import {Uuid} from "../../Shared/src/ValueObject/Objects/Uuid";
import {DocumentBuilder, SwaggerModule} from "@nestjs/swagger";
import {ValidationPipe} from "@nestjs/common";
import {Fetch} from "../../Infrastructure/Fetch/Fetch";
config({
  path: '.env.local',
});

async function bootstrap() {
  const logger = new Logger('Main');
  const app = await NestFactory.create(AppModule, {
    logger: logger,
  });

  app.setGlobalPrefix(process.env.APP_GLOBAL_PREFIX);
  app.enableCors();
  app.use((req: any, res: any, next: NextFunction) => {
    const asyncStorage = app.get(ASYNC_STORAGE);
    const traceId = req.headers['x-request-id'] || new Uuid().toString();
    const store = new Map<string, string>().set('traceId', traceId);
    asyncStorage.run(store, () => {
      next();
    });
  });

  app.useGlobalPipes(new ValidationPipe());

  app.useLogger(app.get<CustomLoggerService>('CustomLogger'));

  Fetch.init()

  const config = new DocumentBuilder()
      .setTitle('TBotAdmin')
      .setDescription('TBotAdmin API')
      .setVersion('1.0.0')
      .build();
  const document = SwaggerModule.createDocument(app, config);
  SwaggerModule.setup('swagger', app, document);

  await app.listen(process.env.APP_PORT);
}

bootstrap().catch((err: unknown) => console.error(err))

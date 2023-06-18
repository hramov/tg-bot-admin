import { Test, TestingModule } from '@nestjs/testing';
import { INestApplication } from '@nestjs/common';
import * as request from 'supertest';
import { AppModule } from '../src/API/v1/app.module';
import {CreateShopDto} from "../src/API/v1/modules/shop/dto/create-shop.dto";

describe('AppController (e2e)', () => {
  let app: INestApplication;

  beforeEach(async () => {
    const moduleFixture: TestingModule = await Test.createTestingModule({
      imports: [AppModule],
    }).compile();

    app = moduleFixture.createNestApplication();
    await app.init();
  });

  it('/api/shop (POST) - Success query', () => {
    const dto: CreateShopDto = {
      local_shop_name: 'Test Shop',
      owner_tg_name: '@therealhramov',
      bot_token: '5945818399:AAHk2WbGLvhD3PqbG0Vg3onzX_vdb_s9SNs',
    }

    return request(app.getHttpServer())
      .post('/api/shop')
        .set('Content-type', 'application/json')
        .send(dto)
      .expect(200)
  });
});

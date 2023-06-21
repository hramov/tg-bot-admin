import { Injectable } from '@nestjs/common';
import {ProductDto} from "./dto/product.dto";
import {Uuid} from "../../../../Shared/src/ValueObject/Objects/Uuid";
import {ProductSearchFilter} from "../../common/filters/product/search.filter";
import {InjectRepository} from "@nestjs/typeorm";
import {ProductEntity} from "../../common/persistent/entity/product.entity";
import {CategoryEntity} from "../../common/persistent/entity/category.entity";
import {ShopEntity} from "../../common/persistent/entity/shop.entity";
import {Repository} from "typeorm";
import {CannotFindCategoryError} from "./error/CannotFindCategory.error";
import {CannotFindShopError} from "./error/CannotFindShop.error";

@Injectable()
export class ProductService {
    constructor(
        @InjectRepository(ProductEntity) private readonly productRepository: Repository<ProductEntity>,
        @InjectRepository(CategoryEntity) private readonly categoryRepository: Repository<CategoryEntity>,
        @InjectRepository(ShopEntity) private readonly shopRepository: Repository<ShopEntity>
    ) {}

    get(filters: ProductSearchFilter) {
        return this.productRepository.find({relations: { category: true, shop: true } })
    }

    getById(productId: Uuid) {
        return this.productRepository.findOne({ where: { id: productId.toString()}, relations: { category: true, shop: true } })
    }

    async save(dto: ProductDto) {
        const product = this.productRepository.create();

        product.id = dto.id;
        product.title = dto.title;
        product.description = dto.description;
        product.images = dto.images;
        product.price = dto.price;
        product.currency = dto.currency;
        product.quantity = dto.quantity;
        product.customFields = dto.custom_fields;

        const category = await this.categoryRepository.findOneBy({ id : dto.category_id });
        if (!category) {
            return new CannotFindCategoryError();
        }

        const shop = await this.shopRepository.findOneBy({ id : dto.shop_id });
        if (!shop) {
            return new CannotFindShopError();
        }

        product.shop = shop;
        product.category = category;

        return this.productRepository.save(product);
    }

    delete(productId: Uuid) {
        return this.productRepository.delete(productId.toString())
    }
}

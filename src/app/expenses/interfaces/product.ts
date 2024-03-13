export class Product {
  id: string;
  code: string;
  name: string;
  description: string;
  image: string;
  price: number;
  category: string;
  quantity: number;
  inventoryStatus: string;
  rating: number;

  constructor() {
    this.id = '1000';
    this.code = 'f230fh0g3';
    this.name = 'Bamboo Watch';
    this.description = 'Product Description';
    this.image = 'bamboo-watch.jpg';
    this.price = 65;
    this.category = 'Accessories';
    this.quantity = 24;
    this.inventoryStatus = 'INSTOCK';
    this.rating = 5;
  }
}

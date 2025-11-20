// types.ts

interface User {
  id: number;
  name: string;
  email: string;
  role: 'admin' | 'user';
}

interface Product {
  id: number;
  name: string;
  price: number;
  quantity: number;
}

interface Order {
  id: number;
  userId: number;
  productId: number;
  quantity: number;
  total: number;
}

interface Address {
  street: string;
  city: string;
  state: string;
  zip: string;
}

interface UserAddress {
  userId: number;
  address: Address;
}

interface ProductVariant {
  id: number;
  productId: number;
  variantName: string;
  price: number;
  quantity: number;
}

interface ProductVariantOption {
  id: number;
  variantId: number;
  optionName: string;
  price: number;
  quantity: number;
}

interface ProductVariantOptionValue {
  id: number;
  optionId: number;
  value: string;
}

interface ProductVariantOptionValueMap {
  [key: string]: number;
}

interface ProductVariantOptionWithValues {
  id: number;
  variantId: number;
  optionName: string;
  values: ProductVariantOptionValue[];
}

interface ProductVariantWithOptions {
  id: number;
  productId: number;
  variantName: string;
  options: ProductVariantOptionWithValues[];
  price: number;
  quantity: number;
}

interface ProductVariation {
  id: number;
  productId: number;
  variantId: number;
  optionId: number;
  value: string;
  quantity: number;
}

interface ProductVariationMap {
  [key: string]: number;
}

interface ProductVariationWithValues {
  id: number;
  productId: number;
  variantId: number;
  values: string[];
  quantity: number;
}

interface ProductVariationWithProduct {
  id: number;
  productId: number;
  variantId: number;
  product: Product;
}

interface OrderItem {
  id: number;
  orderId: number;
  productId: number;
  quantity: number;
  price: number;
  total: number;
}

interface OrderSummary {
  id: number;
  total: number;
  orderItems: OrderItem[];
}
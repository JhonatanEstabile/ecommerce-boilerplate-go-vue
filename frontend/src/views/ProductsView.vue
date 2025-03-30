<script setup>
import DefaultLayout from '@/layouts/DefaultLayout.vue';
import ProductFormModal from '@/components/ProductFormModal.vue';
import api from '../api/api';
import { onMounted, ref } from 'vue';

const product = ref({
  name: '',
  description: '',
  price: 0.0,
  quantity: 0,
  active: true,
});

const products = ref([]);

const loadProducts = () => {
  api.get(
    '/products/',
    {
      withCredentials: true
    }
  ).then((resp) => {
    products.value = resp.data;
  }).catch((error) => {
    console.log(error);
  });
}

onMounted(() => {
  loadProducts();
});

const cleanformFields = () => {
  product.value.name = '';
  product.value.description = '';
  product.value.price = 0;
  product.value.quantity = 0;
  product.value.active = true;
}

const handleSubmit = (event) => {
  event.preventDefault()

  api.post('/products/', {
    name: product.value.name,
    description: product.value.description,
    price: product.value.price,
    stock_quantity: product.value.quantity,
    is_active: product.value.active,
  }).then((req) => {
    cleanformFields();
    loadProducts();
  }).catch((error) => {
    console.log(error);
  });
};

</script>

<template>
    <DefaultLayout>
      <div class="row buttons-margin">
        <div class="col-md-3">
          <ProductFormModal :product-form-data="product" :handle-product-submit="handleSubmit"/>
        </div>
      </div>
      <div class="row">
        <table class="table">
          <thead>
            <tr>
              <th scope="col">id</th>
              <th scope="col">Name</th>
              <th scope="col">Description</th>
              <th scope="col">Price</th>
              <th scope="col">Quantity</th>
              <th scope="col">Is active</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="product in products">
              <th scope="row">{{ product.id }}</th>
              <td>{{ product.name }}</td>
              <td>{{ product.description }}</td>
              <td>{{ product.price }}</td>
              <td>{{ product.stock_quantity }}</td>
              <td>{{ product.is_active }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </DefaultLayout>
</template>
<style>
.buttons-margin {
  margin-top: 25px;
  margin-bottom: 25px;
}
</style>
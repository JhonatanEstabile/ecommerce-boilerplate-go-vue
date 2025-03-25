<script setup>
import { toRefs } from 'vue';

const props = defineProps(['productFormData', 'handleProductSubmit']);
const { productFormData, handleProductSubmit } = toRefs(props);

</script>

<template>
    <div class="modal fade" id="formProductModal" data-bs-backdrop="static" data-bs-keyboard="false" tabindex="-1" aria-labelledby="formProductModalLabel" aria-hidden="true">
      <div class="modal-dialog modal-lg modal-dialog-centered">
        <div class="modal-content">
          <div class="modal-header">
            <h1 class="modal-title fs-5" id="formProductModalLabel">Modal title</h1>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <form v-on:submit="handleProductSubmit">
            <div class="modal-body">
            
              <div class="col-12">
                <label for="name" class="form-label">Product Name:</label>
                <input v-model="productFormData.name" type="text" class="form-control" id="name" placeholder="Product name here">
              </div>
              <div class="col-12">
                <label for="description" class="form-label">Product Description:</label>
                <textarea v-model="productFormData.description" class="form-control" id="description" rows="3"></textarea>
              </div>
            
              <div class="row">
                <div class="col-3">
                  <label for="price" class="form-label">Product price:</label>
                  <input v-model="productFormData.price" type="number" step=".01" class="form-control" id="price" placeholder="Product price here">
                </div>
              
                <div class="col-3">
                  <label for="quantity" class="form-label">Product Quantity:</label>
                  <input v-model="productFormData.quantity" type="number" class="form-control" id="quantity" placeholder="Product quantity here">
                </div>
              
                <div class="col-3 d-flex align-items-center">
                  <div class="form-check">
                    <input class="form-check-input" type="checkbox" v-model="productFormData.active" id="active">
                    <label class="form-check-label" for="active">
                      Is Active
                    </label>
                  </div>
                </div>
              </div>
            </div>
            <br/>
            <div class="modal-footer">
              <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Cancel</button>
              <button type="submit" class="btn btn-success" data-bs-dismiss="modal">Register</button>
            </div>
          </form>
        </div>
      </div>
    </div>

    <button type="button" class="btn btn-success" data-bs-toggle="modal" data-bs-target="#formProductModal">
      Cadastrar
    </button>
</template>

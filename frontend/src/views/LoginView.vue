<script setup>
import { ref } from 'vue';
import api from '../api/api';
import { useRouter } from 'vue-router';

const router = useRouter();

const email = ref('');
const password = ref('');

const handleLogin = () => {
  api.post(
    '/login',
    {
      email: email.value,
      password: password.value,
    },
    {
      withCredentials: true
    }
  ).then((resp) => {
    router.push('/admin/products');
  }).catch((error) => {
    console.log(error);
  });
}

</script>

<template>
  <section class="vh-100">
    <div class="container-fluid">
      <div class="row vh-100">
        <div class="col-sm-6 text-black">
          <div class="d-flex vh-100 align-items-center justify-content-center h-custom-2 px-5 ms-xl-4 mt-5 pt-5 pt-xl-0 mt-xl-n5">
            <form style="width: 23rem;">
              <h3 class="fw-normal mb-3 pb-3" style="letter-spacing: 1px;">Log in</h3>

              <div data-mdb-input-init class="form-outline mb-4">
                <input type="email" id="email" v-model="email" class="form-control form-control-lg" />
                <label class="form-label" for="email">Email address</label>
              </div>

              <div data-mdb-input-init class="form-outline mb-4">
                <input type="password" id="password" v-model="password" class="form-control form-control-lg" />
                <label class="form-label" for="password">Password</label>
              </div>

              <div class="pt-1 mb-4">
                <button @click="handleLogin" data-mdb-button-init data-mdb-ripple-init class="btn btn-info btn-lg btn-block" type="button">Login</button>
              </div>

              <p class="small mb-5 pb-lg-2"><a class="text-muted" href="#!">Forgot password?</a></p>
              <p>Don't have an account? <a href="#!" class="link-info">Register here</a></p>

            </form>

          </div>

        </div>
        <div class="col-sm-6 px-0 d-none d-sm-block">
          <img src="https://mdbcdn.b-cdn.net/img/Photos/new-templates/bootstrap-login-form/img3.webp"
            alt="Login image" class="w-100 vh-100" style="object-fit: cover; object-position: left;">
        </div>
      </div>
    </div>
  </section>
</template>

<style>
.bg-image-vertical {
  position: relative;
  overflow: hidden;
  background-repeat: no-repeat;
  background-position: right center;
  background-size: auto 100%;
}

.vh-100 {
  overflow: hidden;
}

@media (min-width: 1025px) {
  .h-custom-2 {
    height: 100vh;
  }
}
</style>
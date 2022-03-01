<template>
<div class="wrapper">
      <profile-about/>
<div class="blog_post">
        <img src="@/assets/images/logo/1000_signature.png" alt="">
       <label
        for="extension"
        class="mt-1"
      >Allowed JPG or PNG. Max size of 800kB
</label>
      <b-form-file
        id="extension"
        accept=".jpg, .png, .gif"
    />
    </div>
    </div>
</template>

<script>
import ProfileAbout from './ProfileAbout.vue'
import { BFormFile } from 'bootstrap-vue'

export default {
  components: {
    ProfileAbout,
    BFormFile,
  },
}
</script>
<style lang="scss" scoped>
@import '@core/scss/vue/pages/page-profile.scss';

.wrapper {
  height: 100%;
  width: 100%;
  height: 100vh;
}

.blog_post {
  position: relative;
  padding: 6rem 3rem 6rem 6rem;
  background: #fff;
  max-width: 650px;
  border-radius: 10px;
  top: 50%;
  left: 50%;
  transform: translate(-25%, -169%);
  box-shadow: 1px 1px 2rem rgba(0, 0, 0, 0.3);
}

@media only screen and (max-width: 600px) {
  .blog_post {
    transform: translate(-46%, -156%);
  }

  .p-2 {
  transform: translate(31%, -51%) !important;
  padding-bottom: 12px !important;
}

  img {
        transform: translate(-84%, -10%) !important;
        width: 50px;
        height: 30px;

  }

  .mx {
    transform: translate(-60%, -176%) !important;
    /* margin: -148px 0 0.5rem 0; */
    color: #999 !important;
    font-size: 1.25rem !important;
  }

  .mxl {
    transform: translate(220%, -377%) !important;
    /* margin: -148px 0 0.5rem 0; */
    color: #999 !important;
    font-size: 1.25rem !important;
  }

  .rules {
    font-size: 0.7rem !important;
    transform: translate(27%, -529%) !important;
    font-size: 0.6rem !important;
  }

  hr.solid {
    border-top: 0.5px solid #bbb !important;
    transform: translate(-10%, -2491%) !important;
    max-width: 200px !important;
  }
  img {
  height: 100px !important;
  width: 100px !important;
  position: relative !important;
  top: 50% !important;
  left: 50% !important;
  transform: translate(-178%, -50%) !important;
  border-radius: 10% !important;
  box-shadow: 1px 1px 2rem rgba(0, 0, 0, 0.3) !important;
}
}
</style>


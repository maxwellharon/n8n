<template>
  <b-row>
    <b-col md="6">
      <div class="demo-inline-spacing">
        <b-button
          v-ripple.400="'rgba(255, 255, 255, 0.15)'"
          block
          variant="primary"
          :to="{ path: '/pages/profile' }"
        >

          General
        </b-button>
        <b-button
          v-ripple.400="'rgba(113, 102, 240, 0.15)'"
          block
          variant="outline-primary"
          :to="{ path: '/pages/profile/profile-signature' }"
        >
          Signature
        </b-button>
      </div>
    </b-col>

    <!-- <b-col md="6">
      <div class="demo-inline-spacing">
        <b-button
          v-ripple.400="'rgba(113, 102, 240, 0.15)'"
          block
          variant="outline-primary"
        >
          Signature
        </b-button>
      </div>
    </b-col> -->
  </b-row>
</template>

<script>
import { BButton, BCol, BRow } from 'bootstrap-vue'
import Ripple from 'vue-ripple-directive'




export default {
  components: {
    BButton,
    BCol,
    BRow,
  },
  directives: {
    Ripple,
  },
   methods:{
   goToSignature(){
   this.$router.push('/profileSignature'); 
      }
  },
}
</script>
 
<style lang="scss" scoped>
@import '@core/scss/vue/pages/page-profile.scss';

@import url("https://fonts.googleapis.com/css?family=Roboto:300,400,700&display=swap");

.btn-block {
    // display: block;
    width: 61% !important;
     transform: translate(3%, -15%) !important;
}
@media only screen and (max-width: 600px) {
       // display: block;
       .btn-block {
    width: 800px !important;
     transform: translate(23%, -15%) !important;
       }
}
 </style>
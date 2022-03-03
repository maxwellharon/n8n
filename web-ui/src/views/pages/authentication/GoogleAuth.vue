<template>
<div>
 <div>
   <img src="@/assets/images/pages/auth-v1-top-bg.png" alt="">
 <b-col md="7" class="col">
        <b-card class="text-center">
          <b-tabs
            pills
            nav-class="mb-3"
          >
            <b-tab
              title="OTP"
              active
            >
                   <!-- large label  -->
    <b-form-group
      
    >
    <p>Enter Google Auth OTP to log in</p>
      <b-form-input
        id="input-lg"
        size="lg"
        placeholder="OTP"
      />
    </b-form-group>
              <b-button
                v-ripple.400="'rgba(113, 102, 240, 0.15)'"
                variant="outline-primary"
                :to="{ path: '/' }"
              >
                Submit
              </b-button>
            </b-tab>
            <b-tab title="Recovery Code">
                 <!-- large label  -->
                 <p>Enter Google Auth recovery code to log in</p>
    <b-form-group
    >
     <b-form-input
        id="input-lg"
        size="lg"
        placeholder="Recovery Code"
      />
    </b-form-group>
              <b-button
                v-ripple.400="'rgba(113, 102, 240, 0.15)'"
                variant="outline-primary"
              >
                Submit
              </b-button>
            </b-tab>
          </b-tabs>
        </b-card>
      </b-col>
   </div>
    <div class="imgg">
      <!-- <img src="@/assets/images/pages/auth-v1-bottom-bg.png" alt=""> -->
      </div>
      </div>
</template>

<script>
import {BFormInput, BRow, BCol, BFormGroup} from 'bootstrap-vue'
import {
  BCard, BTab, BTabs, BCardText, BCardTitle, BButton,
} from 'bootstrap-vue'
import VuexyLogo from '@core/layouts/components/Logo.vue'
import Ripple from 'vue-ripple-directive'
import { getHomeRouteForLoggedInUser } from '@/auth/utils'

export default {
  components: {
    BFormInput,
    BFormGroup,
    BRow,
    BCol,
    BCard,
    VuexyLogo,
    BTab,
    BTabs,
    BCardText,
    BCardTitle,
    BButton,
  },
   directives: {
    Ripple,
  },
}
</script>
 

<style lang="scss" scoped>
@import '@core/scss/vue/pages/page-auth.scss';
.col{
    transform: translate(39%, -1%) !important;
}
img{
  transform: translate(83%, 79%) !important;
}
.imgg img {
  transform: translate(227%, -66%) !important;
}
@media only screen and (max-width: 600px) {
    .col{
    transform: translate(20%, 54%) !important;
    }    

}
</style>

<template>
  <b-card
    v-if="data"
    class="card-congratulation-medal"
  >
    <h5>Congratulations 🎉 {{ data.name }}!</h5>
    <b-card-text class="font-small-3">
      You have won gold medal
    </b-card-text>
    <h3 class="mb-75 mt-2 pt-50">
      <b-link>${{ kFormatter(data.saleToday) }}</b-link>
    </h3>
    <b-button
      v-ripple.400="'rgba(255, 255, 255, 0.15)'"
      variant="primary"
    >
      View Sales
    </b-button>
    <b-img
      :src="require('@/assets/images/illustration/badge.svg')"
      class="congratulation-medal"
      alt="Medal Pic"
    />
  </b-card>
</template>

<script>
import {
  BCard, BCardText, BLink, BButton, BImg,
} from 'bootstrap-vue'
import Ripple from 'vue-ripple-directive'
import { kFormatter } from '@core/utils/filter'
import BCardActions from '@core/components/b-card-actions/BCardActions.vue'

export default {
  data() {
        return {
            analysis
        }
    },
  components: {
    BCard,
    BCardText,
    BLink,
    BImg,
    BButton,
    BCardActions,
  },
  directives: {
    Ripple,
  },
  props: {
    data: {
      type: Object,
      default: () => {},
    },
  },
  methods: {
    kFormatter,
  },
  computed:{
        grf(){ return this.info.src },
        class_css(){
            if(this.info){
                return 'shadow-right'
            } else {
                return 'shadow-left'
            }
        }
    },
  
}
</script>

<style lang="scss">
  
   .frame{
        min-width: 100%;
        height: 600px;
        border: 1px rgb(211, 210, 210) solid;
        border-radius: 8px;
        
    } 

    .shadow-right{
        -webkit-box-shadow: 5px 5px 25px -7px rgba(0,0,0,0.39); 
        box-shadow: 5px 5px 25px -7px rgba(0,0,0,0.39);
    }

    .shadow-left{
        -webkit-box-shadow: -5px 5px 25px -7px rgba(0,0,0,0.39); 
        box-shadow: -5px 5px 25px -7px rgba(0,0,0,0.39);
    }
</style>
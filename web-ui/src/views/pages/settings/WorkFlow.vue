<template>
<div>
  <b-card no-body>
    <b-card-header class="pb-50">
      <h5>
        FILTERS
      </h5>
    </b-card-header>
    <b-card-actions
          action-collapse
        >
     
      <b-row>
        <b-col
          cols="12"
          md="3"
          class="mb-md-0 mb-2"
        >
         
          <v-select
            :dir="$store.state.appConfig.isRTL ? 'rtl' : 'ltr'"
            :value="roleFilter"
            :options="roleOptions"
            class="w-100"
            :reduce="val => val.value"
            @input="(val) => $emit('update:roleFilter', val)"
            placeholder="Status ..."
          />
        </b-col>
     
      
       
       
         <b-col
          cols="12"
          md="3"
          class="mb-md-0 mb-2"
        >
          
          <b-button
      v-ripple.400="'rgba(113, 102, 240, 0.15)'"
      variant="outline-primary"
      class="w-100"
    >
    <!-- <feather-icon
                icon="FilterIcon"
                size="20"
                v-ripple.400="'rgba(255, 255, 255, 0.15)'"
              /> -->
              <font-awesome-icon icon="fa-solid fa-filter" />
              FILTER
    </b-button>
        </b-col>
      </b-row>
    </b-card-actions> 
  </b-card>
  <work-flow-list/>
  </div>
</template>

<script>
import BCardActions from '@core/components/b-card-actions/BCardActions.vue'

import {
  BCard, BCardHeader, BCardBody, BRow, BCol, BButton
} from 'bootstrap-vue'
import vSelect from 'vue-select'
import Ripple from 'vue-ripple-directive'
import FeatherIcon from '@core/components/feather-icon/FeatherIcon.vue'
import WorkFlowList from './WorkFlowList.vue'


export default {
  components: {
    BCardActions,
    BRow,
    BCol,
    BCard,
    BCardHeader,
    BCardBody,
    vSelect,
    BButton,
    FeatherIcon,
    WorkFlowList,
  },
  directives: {
    Ripple,
  },
  props: {
    roleFilter: {
      type: [String, null],
      default: null,
    },
    planFilter: {
      type: [String, null],
      default: null,
    },
    statusFilter: {
      type: [String, null],
      default: null,
    },
    roleOptions: {
      type: Array,
      required: true,
    },
    planOptions: {
      type: Array,
      required: true,
    },
    statusOptions: {
      type: Array,
      required: true,
    },
  },
}
</script>

<style lang="scss">
@import '@core/scss/vue/libs/vue-select.scss';
</style>

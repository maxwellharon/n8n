<template>
 <section id="dashboard-ecommerce">
      <b-row class="match-height">
      <b-col
        xl="3"
        sm="3"
      >
        <iframe
    src="http://172.16.10.204:3000/public/question/f165f73d-f902-4e11-839d-327319a83c29"
    frameborder="0"
    allowtransparency
></iframe>
      </b-col>
      <b-col
        xl="3"
        sm="3"
      >
        <iframe
    src="http://172.16.10.204:3000/public/question/d1aabc03-462d-454d-80d1-482555a341f2"
    frameborder="0"
    allowtransparency
></iframe>
      </b-col>
      <b-col
        xl="3"
        sm="3"
      >
        <iframe
    src="http://172.16.10.204:3000/public/question/969a3089-4245-41f0-8736-52adfe64780e"
    frameborder="0"
    allowtransparency
></iframe>
      </b-col>
        <b-col
        xl="3"
        md="6"
      >
       <iframe
    src="http://172.16.10.204:3000/public/question/43f4c208-5729-4c68-a5cc-1d0f05c6bfe2"
    frameborder="0"
    allowtransparency
></iframe>
      </b-col>
    </b-row>

     <b-row class="match-height">
      <b-col
        xl="12"
        md="6"
      >
        <iframe
    src="http://172.16.10.204:3000/public/question/571cc6db-1ef7-4e20-87db-6c0522fbb0d1"
    frameborder="0"
    height="600"
    allowtransparency
></iframe>
      </b-col>
      </b-row>

</section>
</template>

<script>
import { BRow, BCol } from 'bootstrap-vue'

import { getUserData } from '@/auth/utils'
import EcommerceMedal from './EcommerceMedal.vue'
import EcommerceStatistics from './EcommerceStatistics.vue'
import EcommerceRevenueReport from './EcommerceRevenueReport.vue'
import EcommerceOrderChart from './EcommerceOrderChart.vue'
import EcommerceProfitChart from './EcommerceProfitChart.vue'
import EcommerceEarningsChart from './EcommerceEarningsChart.vue'
import EcommerceCompanyTable from './EcommerceCompanyTable.vue'
import EcommerceMeetup from './EcommerceMeetup.vue'
import EcommerceBrowserStates from './EcommerceBrowserStates.vue'
import EcommerceGoalOverview from './EcommerceGoalOverview.vue'
import EcommerceTransactions from './EcommerceTransactions.vue'
import draggable from 'vuedraggable'
import Prism from 'vue-prism-component'
import 'prismjs'
import 'prismjs/themes/prism-tomorrow.css'

export default {
  components: {
    BRow,
    BCol,

    EcommerceMedal,
    EcommerceStatistics,
    EcommerceRevenueReport,
    EcommerceOrderChart,
    EcommerceProfitChart,
    EcommerceEarningsChart,
    EcommerceCompanyTable,
    EcommerceMeetup,
    EcommerceBrowserStates,
    EcommerceGoalOverview,
    EcommerceTransactions,
    Prism,
    draggable,
  },
  data() {
    return {
      data: {},
    }
  },
  created() {
    // data
    this.$http.get('/ecommerce/data')
      .then(response => {
        this.data = response.data

        // ? Your API will return name of logged in user or you might just directly get name of logged in user
        // ? This is just for demo purpose
        const userData = getUserData()
        this.data.congratulations.name = userData.fullName.split(' ')[0] || userData.username
      })
  },
}
</script>
<style lang="scss" scoped>
@import '@core/scss/vue/pages/dashboard-ecommerce.scss';
@import '@core/scss/vue/libs/chart-apex.scss';
</style>

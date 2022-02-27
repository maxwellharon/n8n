<template>
<b-input-group class="input-group-merge ml-1 w-50">
  <b-dropdown
      id="dropdown-1"
      v-ripple.400="'rgba(255, 255, 255, 0.15)'"
      text="Case File ID"
      variant="primary"
    >
      <b-dropdown-item>MemberNo</b-dropdown-item>
      <b-dropdown-item>Case File ID</b-dropdown-item>
      <b-dropdown-item>Names</b-dropdown-item>
       <b-dropdown-item>Identification</b-dropdown-item>
        <b-dropdown-item>LoanAccount</b-dropdown-item>
         <b-dropdown-item>Phones</b-dropdown-item>
          <b-dropdown-item>Email</b-dropdown-item>
    </b-dropdown>

              <b-input-group-prepend is-text>
                <!-- <feather-icon
                  icon="SearchIcon"
                  class="text-muted"
                /> -->
              </b-input-group-prepend>
              <b-form-input
                v-model="searchQuery"
                placeholder="Search..."
              />
            </b-input-group>
</template>

<script>
import {
  BFormInput, BLink, BImg, BAvatar, BInputGroup, BInputGroupPrepend, BInputGroupAppend, BFormGroup, BFormTextarea, BDropdown, BDropdownItem
} from 'bootstrap-vue'
import Ripple from 'vue-ripple-directive'

import { ref, watch } from '@vue/composition-api'
import VuePerfectScrollbar from 'vue-perfect-scrollbar'
import useAutoSuggest from '@core/components/app-auto-suggest/useAutoSuggest'
import { title } from '@core/utils/filter'
import router from '@/router'
import store from '@/store'
import searchAndBookmarkData from '../search-and-bookmark-data'

export default {
  components: {
    BFormInput,
    BLink,
    BImg,
    BAvatar,
    VuePerfectScrollbar,
    BInputGroup,
    BInputGroupPrepend,
    BInputGroupAppend,
    BFormGroup,
    BFormTextarea,
    BDropdown,
    BDropdownItem,
  },
  directives: {
    Ripple,
  },
  setup() {
    const showSearchBar = ref(false)

    const perfectScrollbarSettings = {
      maxScrollbarLength: 60,
    }

    const suggestionSelected = (grpName, suggestion) => {
      // If parameter is not provided => Use current selected
      if (!suggestion) {
        // If currentSelected value is -1 => No value/item is selected (Prevent Errors)
        /* eslint-disable no-use-before-define, no-param-reassign */
        if (currentSelected.value !== -1) {
          /* eslint-disable no-use-before-define, no-param-reassign */
          const [grpIndex, itemIndex] = currentSelected.value.split('.')
          grpName = Object.keys(filteredData.value)[grpIndex]
          suggestion = filteredData.value[grpName][itemIndex]
          /* eslint-enable */
        }
      }
      if (grpName === 'pages') router.push(suggestion.route).catch(() => {})
      // eslint-disable-next-line no-use-before-define
      resetsearchQuery()
      showSearchBar.value = false
    }

    const {
      searchQuery,
      resetsearchQuery,
      filteredData,
    } = useAutoSuggest({ data: searchAndBookmarkData, searchLimit: 4 })

    watch(searchQuery, val => {
      store.commit('app/TOGGLE_OVERLAY', Boolean(val))
    })

    const currentSelected = ref(-1)
    watch(filteredData, val => {
      if (!Object.values(val).some(obj => obj.length)) {
        currentSelected.value = -1
      } else {
        // Auto Select first item if it's not item-404
        let grpIndex = null

        // eslint-disable-next-line no-restricted-syntax
        for (const [index, grpSuggestions] of Object.values(val).entries()) {
          if (grpSuggestions.length) {
            grpIndex = index
            break
          }
        }

        if (grpIndex !== null) currentSelected.value = `${grpIndex}.0`
      }
    })

    const increaseIndex = (val = true) => {
      /* eslint-disable no-lonely-if, no-plusplus */

      // If there's no matching items
      if (!Object.values(filteredData.value).some(grpItems => grpItems.length)) return

      const [grpIndex, itemIndex] = currentSelected.value.split('.')

      const grpArr = Object.entries(filteredData.value)
      const activeGrpTotalItems = grpArr[grpIndex][1].length

      if (val) {
        // If active item is not of last item in grp
        if (activeGrpTotalItems - 1 > itemIndex) {
          currentSelected.value = `${grpIndex}.${Number(itemIndex) + 1}`

        // If active item grp is not last in grp list
        } else if (grpIndex < grpArr.length - 1) {
          for (let i = Number(grpIndex) + 1; i < grpArr.length; i++) {
            // If navigating group have items => Then move in that group
            if (grpArr[i][1].length > 0) {
              currentSelected.value = `${Number(i)}.0`
              break
            }
          }
        }
      } else {
        // If active item is not of first item in grp
        if (Number(itemIndex)) {
          currentSelected.value = `${grpIndex}.${Number(itemIndex) - 1}`

        // If active item grp  is not first in grp list
        } else if (Number(grpIndex)) {
          for (let i = Number(grpIndex) - 1; i >= 0; i--) {
            // If navigating group have items => Then move in that group
            if (grpArr[i][1].length > 0) {
              currentSelected.value = `${i}.${grpArr[i][1].length - 1}`
              break
            }
          }
        }
      }
      /* eslint-enable no-lonely-if, no-plusplus */
    }

    return {
      showSearchBar,
      perfectScrollbarSettings,
      searchAndBookmarkData,
      title,
      suggestionSelected,
      currentSelected,
      increaseIndex,

      // AutoSuggest
      searchQuery,
      resetsearchQuery,
      filteredData,
    }
  },
}
</script>

<style lang="scss" scoped>
@import '~@core/scss/base/bootstrap-extended/include';
@import '~@core/scss/base/components/variables-dark';

</style>

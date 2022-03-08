<template>

  <div>

    <user-list-add-new
      :is-add-new-user-sidebar-active.sync="isAddNewUserSidebarActive"
      :role-options="roleOptions"
      :plan-options="planOptions"
      @refetch-data="refetchData"
    />

    <!-- Filters -->
    <users-list-filters
      :role-filter.sync="roleFilter"
      :plan-filter.sync="planFilter"
      :status-filter.sync="statusFilter"
      :role-options="roleOptions"
      :plan-options="planOptions"
      :status-options="statusOptions"
    />
    <!-- <filters/> -->
      
    <!-- Table Container Card -->
    <b-card
      no-body
      class="mb-0"
    >

      <div class="m-2">

        <!-- Table Top -->
        <b-row>

          <!-- Per Page -->
          <b-col
            cols="12"
            md="6"
            class="d-flex align-items-center justify-content-start mb-1 mb-md-0"
          >
            <label>Show</label>
            <v-select
              v-model="perPage"
              :dir="$store.state.appConfig.isRTL ? 'rtl' : 'ltr'"
              :options="perPageOptions"
              :clearable="false"
              class="per-page-selector d-inline-block mx-50"
            />
            <label>entries</label>
          </b-col>

          <!-- Search -->
          <b-col
            cols="12"
            md="6"
          >
            <div class="d-flex align-items-center justify-content-end">
              <b-form-input
                v-model="searchQuery"
                class="d-inline-blockk mr-1"
                placeholder="Search..."
              />
              <b-button
                variant="primary"
                @click="isAddNewUserSidebarActive = true"
              >
                <span class="text-nowrap"><font-awesome-icon icon="fa-solid fa-circle-plus" /> Add User</span>
              </b-button>
              <b-button-group>
                  <!-- <font-awesome-icon icon="fa-solid fa-download" /> -->
                
               <b-dropdown
               class="tes"
                v-ripple.400="'rgba(113, 102, 240, 0.15)'"
                right
                variant="info"
                text="Export"
              >
         
                  <b-dropdown-item>CSV</b-dropdown-item>
                  <b-dropdown-item>XLSX</b-dropdown-item>
                  <b-dropdown-divider />
                </b-dropdown>
              </b-button-group>
            </div>
          </b-col>
        </b-row>

      </div>

      <b-table
        ref="refUserListTable"
        class="position-relative"
        :items="fetchUsers"
        responsive
        :fields="tableColumns"
        primary-key="id"
        :sort-by.sync="sortBy"
        show-empty
        empty-text="No matching records found"
        :sort-desc.sync="isSortDirDesc"
      >
  
        <!-- Column: Actions -->
        <template #cell(actions)>
          <div class="d-flex mb-2">
              <div
        value="left"
        plain
        >
        <span class="ml-50 mr-1">
                          <b-button
               v-b-tooltip.hover.v-dark
              title="Edit"           
              @click="isAddNewUserSidebarActive = true"
              variant="outline-info"
              class="btn-icon rounded-circle"
              size="sm"
            >

              <feather-icon icon="EditIcon" />
            </b-button>
            </span>
              </div>
              <div 
         value="left"
        plain>
        <span class="ml-50 mr-1">
                          <b-button
              v-b-tooltip.hover.v-dark
              title="Deactivate"            
              @click="warning"            
              variant="outline-warning"
              class="btn-icon rounded-circle"
              size="sm"
            >
              <feather-icon icon="XCircleIcon" />
            </b-button>
            </span>
              </div>
              <div  value="left"
        plain>
        <span class="ml-50 mr-1">
                          <b-button
              v-b-tooltip.hover.v-dark
      title="Reset Password"            
              @click="info"            
              variant="outline-primary"
              class="btn-icon rounded-circle"
              size="sm"
            >
              <feather-icon icon="RefreshCwIcon" />
            </b-button>
            </span>
              </div>
                <div 
         value="left"
        plain>
        <span class="ml-50 mr-1">
                          <b-button
                          v-b-tooltip.hover.v-dark
      title="Delete"
                          @click="error"
              variant="outline-danger"
              class="btn-icon rounded-circle"
              size="sm"
            >
              <feather-icon icon="Trash2Icon" />
            </b-button>
            </span>
              </div>
             <div 
         value="left"
        plain>
        <span class="ml-50 mr-1">
                          <b-button
              v-b-tooltip.hover.v-dark
      title="Login As"            
              @click="danger"            
              variant="outline-secondary"
              class="btn-icon rounded-circle"
              size="sm"
            >
              <feather-icon icon="LogInIcon" />
            </b-button>
            </span>
              </div>

          </div>
         
    
          <!-- <b-dropdown
            variant="link"
            no-caret
            :right="$store.state.appConfig.isRTL"
          >

            <template #button-content>
              <feather-icon
                icon="MoreVerticalIcon"
                size="16"
                class="align-middle text-body"
              />
            </template>
            <b-dropdown-item :to="{ name: 'apps-users-view', params: { id: data.item.id } }">
              <feather-icon icon="FileTextIcon" />
              <span class="align-middle ml-50">Details</span>
            </b-dropdown-item>

            <b-dropdown-item :to="{ name: 'apps-users-edit', params: { id: data.item.id } }">
              <feather-icon icon="EditIcon" />
              <span class="align-middle ml-50">Edit</span>
            </b-dropdown-item>

            <b-dropdown-item>
              <feather-icon icon="TrashIcon" />
              <span class="align-middle ml-50">Delete</span>
            </b-dropdown-item>
          </b-dropdown> -->
        </template>


        <!-- Column: User -->
        <template #cell(user)="data">
          <b-media vertical-align="center">
            <!-- <template #aside>
              <b-avatar
                size="32"
                :src="data.item.avatar"
                :text="avatarText(data.item.fullName)"
                :variant="`light-${resolveUserRoleVariant(data.item.role)}`"
                :to="{ name: 'apps-users-view', params: { id: data.item.id } }"
              />
            </template> -->
            <b-link
              :to="{ name: 'apps-users-view', params: { id: data.item.id } }"
              class="font-weight-bold d-block text-nowrap"
            >
              {{ data.item.fullName }}
            </b-link>
            <small class="text-muted">@{{ data.item.username }}</small>
          </b-media>
        </template>

        <!-- Column: Role -->
        <template #cell(role)="data">
          <div class="text-nowrap">
            <feather-icon
              :icon="resolveUserRoleIcon(data.item.role)"
              size="18"
              class="mr-50"
              :class="`text-${resolveUserRoleVariant(data.item.role)}`"
            />
            <span class="align-text-top text-capitalize">{{ data.item.role }}</span>
          </div>
        </template>

        <!-- Column: Status -->
        <template #cell(status)="data">
          <b-badge
            pill
            :variant="`light-${resolveUserStatusVariant(data.item.status)}`"
            class="text-capitalize"
          >
            {{ data.item.status }}
          </b-badge>
        </template>


      </b-table>
      <div class="mx-2 mb-2">
        <b-row>
      <!-- Pagination -->
          <b-col
            cols="12"
            sm="6"
          >

            <b-pagination
              v-model="currentPage"
              :total-rows="totalUsers"
              :per-page="perPage"
              first-number
              last-number
              class="mb-0 mt-1 mt-sm-0"
              prev-class="prev-item"
              next-class="next-item"
              value="default"
            >
              <template #prev-text>
                <feather-icon
                  icon="ChevronLeftIcon"
                  size="18"
                />
              </template>
              <template #next-text>
                <feather-icon
                  icon="ChevronRightIcon"
                  size="18"
                />
              </template>
            </b-pagination>

          </b-col>
          <b-col
            cols="12"
            sm="6"
            class="d-flex justify-content-end"
          >
            <span class="text-muted">Showing {{ dataMeta.from }} to {{ dataMeta.to }} of {{ dataMeta.of }} entries</span>
          </b-col>
    

        </b-row>
      </div>
       <div class="d-flex mb-2">
      </div>
    </b-card>
    
  </div>
</template>

<script>
import {
  BCard, BRow, BCol, BFormInput, BButton, BTable, BMedia, BAvatar, BLink,
  BBadge, BDropdown, BDropdownItem, BPagination,BFormRadio ,BNav, BNavItem,
} from 'bootstrap-vue'
import vSelect from 'vue-select'
import store from '@/store'
import { ref, onUnmounted } from '@vue/composition-api'
import { avatarText } from '@core/utils/filter'
import UsersListFilters from './UsersListFilters.vue'
import useUsersList from './useUsersList'
import userStoreModule from '../userStoreModule'
import UserListAddNew from './UserListAddNew.vue'
import FeatherIcon from '@core/components/feather-icon/FeatherIcon.vue'
import Ripple from 'vue-ripple-directive'
import ToastificationContent1 from '@core/components/toastification/ToastificationContent1.vue'
import Filters from './Filters.vue'




export default {
  components: {
    UsersListFilters,
    UserListAddNew,
    Filters,

    BCard,
    BFormRadio,
    BRow,
    BCol,
    BFormInput,
    BButton,
    BTable,
    BMedia,
    BNav,
    BNavItem,
    BAvatar,
    BLink,
    BBadge,
    BDropdown,
    BDropdownItem,
    BPagination,
    FeatherIcon,
    ToastificationContent1,

    vSelect,
  },
  // props: ['variant'],
  // computed: {
  //   style () {
  //     return 'variant: 'rgba(40, 199, 111, 0.15)''
  //   }
  // }
   directives: {
    Ripple,
  },
  data: () => ({
    selected: 'center',
  }),
  setup() {
    const USER_APP_STORE_MODULE_NAME = 'app-user'

    // Register module
    if (!store.hasModule(USER_APP_STORE_MODULE_NAME)) store.registerModule(USER_APP_STORE_MODULE_NAME, userStoreModule)

    // UnRegister on leave
    onUnmounted(() => {
      if (store.hasModule(USER_APP_STORE_MODULE_NAME)) store.unregisterModule(USER_APP_STORE_MODULE_NAME)
    })

    const isAddNewUserSidebarActive = ref(false)

    const roleOptions = [
      { label: 'Administrator', value: 'admin' },
      { label: 'Manager', value: 'author' },
      { label: 'Branch Manager', value: 'editor' },
      { label: 'Team Leader', value: 'maintainer' },
      { label: 'Account Manager', value: 'subscriber' },
       { label: 'Data Entry', value: 'subscriber' },
        { label: 'Agency Representative', value: 'subscriber' },
    ]

    const planOptions = [
      { label: 'Basic', value: 'basic' },
      { label: 'Company', value: 'company' },
      { label: 'Enterprise', value: 'enterprise' },
      { label: 'Team', value: 'team' },
    ]

    const statusOptions = [
      { label: 'Pending', value: 'pending' },
      { label: 'Active', value: 'active' },
      { label: 'Inactive', value: 'inactive' },
    ]

    const {
      fetchUsers,
      tableColumns,
      perPage,
      currentPage,
      totalUsers,
      dataMeta,
      perPageOptions,
      searchQuery,
      sortBy,
      isSortDirDesc,
      refUserListTable,
      refetchData,

      // UI
      resolveUserRoleVariant,
      resolveUserRoleIcon,
      resolveUserStatusVariant,

      // Extra Filters
      roleFilter,
      planFilter,
      statusFilter,
    } = useUsersList()

    return {

      // Sidebar
      isAddNewUserSidebarActive,

      fetchUsers,
      tableColumns,
      perPage,
      currentPage,
      totalUsers,
      dataMeta,
      perPageOptions,
      searchQuery,
      sortBy,
      isSortDirDesc,
      refUserListTable,
      refetchData,

      // Filter
      avatarText,

      // UI
      resolveUserRoleVariant,
      resolveUserRoleIcon,
      resolveUserStatusVariant,

      roleOptions,
      planOptions,
      statusOptions,

      // Extra Filters
      roleFilter,
      planFilter,
      statusFilter,
    }
  },
   methods: {

 showToast(variant) {
      this.$toast({
        component: ToastificationContent1,
        props: {
          title: 'Success!',
          icon: 'CheckCircleIcon',
          text: 'User Deactivated',
          variant: "outline-success",
        },
      })
    },
    showToast1(variant) {
      this.$toast({
        component: ToastificationContent1,
        props: {
          title: 'Success!',
          icon: 'CheckCircleIcon',
          text: 'Password Reset.Email sent to user with instructions.',
          variant: "outline-success",
        },
      })
    },
    // success
    success() {
      this.$swal({
        title: 'Good job!',
        text: 'You clicked the button!',
        icon: 'success',
        customClass: {
          confirmButton: 'btn btn-primary',
        },
        buttonsStyling: false,
      })
    },

    // error
    error() {
      this.$swal({
        title: 'Delete?',
        text: ' This event is not recoverable!',
        icon: 'error',
        showCancelButton: true,
        confirmButtonText: 'Delete',
        customClass: {
          confirmButton: 'btn btn-danger',
          cancelButton: 'btn btn-outline-secondary ml-1',
        },
        buttonsStyling: false,
      })
    },

    // warning
    warning() {
      this.$swal({
        title: 'Deactivate Selected?',
        text: ' You can Re-activate later again',
        icon: 'warning',
        showCancelButton: true,
        customClass: {
          confirmButton: 'btn btn-warning',
          cancelButton: 'btn btn-outline-secondary ml-1',
          confirmButtonText: 'Deactivate',
        },
        buttonsStyling: false,
         }).then(result => {
        if (result.value) {
          this.$emit(this.showToast())
        }
      })
    },

    // info
    info() {
      this.$swal({
        title: 'Reset Password?',
        text: 'Email will be sent to user with Instructions',
        icon: 'info',
        showCancelButton: true,
        customClass: {
          confirmButton: 'btn btn-info',
          cancelButton: 'btn btn-outline-secondary ml-1',
        },
          buttonsStyling: false,
         }).then(result => {
        if (result.value) {
          this.$emit(this.showToast1())
        }
      })
    },
  },
}
</script>

<style lang="scss" scoped>
.per-page-selector {
  width: 90px;
}
</style>

<style lang="scss" scoped>
@import '@core/scss/vue/libs/vue-select.scss';

.tes {
  padding-left: 40px;
}
.d-inline-blockk{
  width: 400px;
}
</style>

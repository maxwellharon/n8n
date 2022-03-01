<template>
<div class="wrapper">
    <profile-about/>
    <div class="blog_post">
       
       
          <!-- success -->
    <div>
        <h2>User Settings</h2>
        <hr>
        <ol><li><h4>Menu Layout </h4>
         <b-form-radio-group
            v-model="layoutType"
            name="layout-type"
            :options="layoutTypeOptions"
          />
      </li></ol>
    </div>
    </div>
    </div>
</template>

<script>
import {
  BRow, BCol, BCard, BCardText, BButton, BFormCheckbox, BLink, BFormRadioGroup, BFormGroup,
} from 'bootstrap-vue'
import Ripple from 'vue-ripple-directive'
import ProfileAbout from './ProfileAbout.vue'
import vSelect from 'vue-select'
import VuePerfectScrollbar from 'vue-perfect-scrollbar'
import useAppCustomizer from '@core/layouts/components/app-customizer/useAppCustomizer.js'

export default {
  components: {
    BRow,
    BCol,
    BCard,
    BCardText,
    BButton,
    BFormCheckbox,
    ProfileAbout,
    // BSV
    BLink,
    BFormRadioGroup,
    BFormCheckbox,
    BFormGroup,

    // 3rd party
    vSelect,
    VuePerfectScrollbar,

  },
  directives: {
    Ripple,
  },
    setup() {
    const {
      // Vertical Menu
      isVerticalMenuCollapsed,

      // Customizer
      isCustomizerOpen,

      // Skin
      skin,
      skinOptions,

      // Content Width
      contentWidth,
      contentWidthOptions,

      // RTL
      isRTL,

      // routerTransition
      routerTransition,
      routerTransitionOptions,

      // Layout Type
      layoutType,
      layoutTypeOptions,

      // NavMenu Hidden
      isNavMenuHidden,

      // Navbar
      navbarColors,
      navbarTypes,
      navbarBackgroundColor,
      navbarType,

      // Footer
      footerTypes,
      footerType,
    } = useAppCustomizer()

    if (layoutType.value === 'horizontal') {
      // Remove semi-dark skin option in horizontal layout
      const skinSemiDarkIndex = skinOptions.findIndex(s => s.value === 'semi-dark')
      delete skinOptions[skinSemiDarkIndex]

      // Remove menu hidden radio in horizontal layout => As we already have switch for it
      const menuHiddneIndex = navbarTypes.findIndex(t => t.value === 'hidden')
      delete navbarTypes[menuHiddneIndex]
    }

    // Perfect Scrollbar
    const perfectScrollbarSettings = {
      maxScrollbarLength: 60,
      wheelPropagation: false,
    }

    return {
      // Vertical Menu
      isVerticalMenuCollapsed,

      // Customizer
      isCustomizerOpen,

      // Skin
      skin,
      skinOptions,

      // Content Width
      contentWidth,
      contentWidthOptions,

      // RTL
      isRTL,

      // routerTransition
      routerTransition,
      routerTransitionOptions,

      // Layout Type
      layoutType,
      layoutTypeOptions,

      // NavMenu Hidden
      isNavMenuHidden,

      // Navbar
      navbarColors,
      navbarTypes,
      navbarBackgroundColor,
      navbarType,

      // Footer
      footerTypes,
      footerType,

      // Perfect Scrollbar
      perfectScrollbarSettings,
    }
  },
}
</script>
<style lang="scss">
@import '@core/scss/vue/libs/vue-select.scss';
</style>

<style lang="scss">
@import '@core/scss/vue/pages/page-profile.scss';
@import '~@core/scss/base/bootstrap-extended/include';
@import '~@core/scss/base/components/variables-dark';
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
  transform: translate(-22%, -159%);
  box-shadow: 1px 1px 2rem rgba(0, 0, 0, 0.3);
}

.customizer-section {
  padding: 1.5rem;
    border-bottom: 1px solid #ebe9f1;

  .dark-layout & {
    border-color: $theme-dark-border-color;
  }

  #skin-radio-group ::v-deep {
    .custom-control-inline {
      margin-right: 0.7rem;
    }
  }

  .form-group {
    margin-bottom: 1.5rem;;
    &:last-of-type {
    margin-bottom: 0;
    }
    ::v-deep legend {
      font-weight: 500;
    }
  }
}

.mark-active {
  box-shadow: 0 0 0 0.2rem rgba(38,143,255,.5);
}

.ps-customizer-area {
  height: calc(100% - 83px)
}
</style>
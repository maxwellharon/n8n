<template>
  <div class="wrapper">
    <profile-about/>
    <div class="blog_post">
      <!-- <div class="img_pod"> -->
      <!-- </div> -->
      <div class="container_copy">
        <img
          src="https://pbs.twimg.com/profile_images/890901007387025408/oztASP4n.jpg"
          alt="random image"
        />
         <b-button class="mx"
      variant="secondary" size="sm"
    >
      Choose ...
    </b-button>

        <!-- <button class="mx">Choose</button> -->
        <b-button class="mxl"
      v-ripple.400="'rgba(113, 102, 240, 0.15)'"
          size="sm"
          variant="outline-primary"
    >
      Update
    </b-button>
        <!-- <button class="mxl">Update</button> -->
        <p class="rules">Allowed JPG or PNG. Max size of 800kB</p>
        <dl>
          <dt>
            <div class="p-2" id="main"><b>NAMES:</b></div>
          </dt>
          <hr class="solid" />
          <dt>
            <div class="p-2" id="main"><b>EMAIL:</b></div>
          </dt>
          <hr class="solid" />
          <dt>
            <p class="p-2" id="main"><b>OFFICE PHONE:</b></p>
          </dt>
          <hr class="solid" />
          <dt>
            <div class="p-2" id="main"><b>ROLE:</b></div>
          </dt>
        </dl>
      </div>
      <!-- <a class="btn_primary" href='#'>Read More</a> -->
    </div>
  </div>
</template>
<script>
import { BButton} from 'bootstrap-vue'
import { BRow, BCol } from 'bootstrap-vue'
import Ripple from 'vue-ripple-directive'
import ProfileAbout from './ProfileAbout.vue'

export default {
  components: {
    BButton,
    ProfileAbout,
    BRow,
    BCol,
  },
  directives: {
    Ripple,
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped >
@import '@core/scss/vue/pages/page-profile.scss';


#main {
  margin-left: -200px !important;
}

* {
  padding: 0;
  margin: 0;
}

// html {
//   font-family: "Roboto", sans-serif !important;
//   font-size: 12px !important;
//   overflow: hidden;
// }

// body {
//   background: #1488cc;
//   background: -webkit-linear-gradient(to right, #2b32b2, #1488cc);
//   background: linear-gradient(to right, #2b32b2, #1488cc);
// }

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
  transform: translate(-25%, -99%);
  box-shadow: 1px 1px 2rem rgba(0, 0, 0, 0.3);
}

.img_pod {
  height: 110px;
  width: 110px;
  background: linear-gradient(90deg, #ff9966, #ff5e62);
  z-index: -1;
  box-shadow: 1px 1px 2rem rgba(0, 0, 0, 0.3);
  border-radius: 100%;
  position: absolute;
  left: -50px;
  top: 10px;
  transform: translateY(-50%);
}

img {
  height: 50px !important;
  width: 50px !important;
  position: relative !important;
  top: 50% !important;
  left: 50% !important;
  transform: translate(-591%, -28%) !important;
  border-radius: 10% !important;
  box-shadow: 1px 1px 2rem rgba(0, 0, 0, 0.3) !important;
}

.mx {
  transform: translate(8%, 190%) !important;
  height: 33px;
  width: 100px;
  margin: -148px 0 0.5rem 0;
}
.mxl {
  transform: translate(55%, -42%) !important;
  height: 33px;
  width: 70px;
  
}

.rules {
  font-size: 0.9rem !important;
  transform: translate(10%, -96%) !important;
}

// h1 {
//   margin: 0 0 1rem 0 !important;
//   font-size: 2.5rem !important;
//   letter-spacing: 0.5px !important;
//   color: #333 !important;
// }

/* p {
  margin: 0 0 4.5rem 0;
  font-size: 1.5rem;
  color: #333;
} */

/* Solid border */
hr.solid {
  border-top: 0.5px solid #bbb !important;
  transform: translate(-7%, -2800%) !important;
}

.p-2 {
  transform: translate(20%, -51%) !important;
  padding-bottom: 12px !important;
}

@media only screen and (max-width: 600px) {
  .blog_post {
    transform: translate(-29%, -76%);
  }

  .p-2 {
  transform: translate(31%, -51%) !important;
  padding-bottom: 12px !important;
}

  img {
    transform: translate(-187%, -50%) !important;
  }

  .mx {
    transform: translate(-57%, 57%) !important;
    /* margin: -148px 0 0.5rem 0; */
    color: #999 !important;
    // font-size: 1.25rem !important;
  }

  .mxl {
    transform: translate(220%, -377%) !important;
    /* margin: -148px 0 0.5rem 0; */
    // color: #999 !important;
    // font-size: 1.25rem !important;
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

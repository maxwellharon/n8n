import Vue from 'vue'
import FeatherIcon from '@core/components/feather-icon/FeatherIcon.vue'
import { library } from '@fortawesome/fontawesome-svg-core'
import { faCircleXmark, faLockOpen, faSpinner } from '@fortawesome/free-solid-svg-icons'
import { faCirclePlus } from '@fortawesome/free-solid-svg-icons'
import { faFilter } from '@fortawesome/free-solid-svg-icons'
import { faDownload } from '@fortawesome/free-solid-svg-icons'
import { faGear } from '@fortawesome/free-solid-svg-icons'
import { faSignature } from '@fortawesome/free-solid-svg-icons'
import { faUser } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { faTrash } from '@fortawesome/free-solid-svg-icons'
import { faRightFromBracket } from '@fortawesome/free-solid-svg-icons'
import { faArrowsRotate } from '@fortawesome/free-solid-svg-icons'
import { faPenToSquare } from '@fortawesome/free-solid-svg-icons'
import { faRightToBracket } from '@fortawesome/free-solid-svg-icons'
library.add(faSpinner)
library.add(faCirclePlus)
library.add(faFilter)
library.add(faDownload)
library.add(faSignature)
library.add(faGear)
library.add(faUser)
library.add(faTrash)
library.add(faRightFromBracket)
library.add(faRightToBracket)
library.add(faArrowsRotate)
library.add(faLockOpen)
library.add(faCircleXmark)
library.add(faPenToSquare)
Vue.component('font-awesome-icon', FontAwesomeIcon)



Vue.component(FeatherIcon.name, FeatherIcon)
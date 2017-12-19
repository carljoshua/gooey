import Browse from './templates/browse.vue'
import Structure from './templates/struct.vue'
import Query from './templates/query.vue'
import Splash from './templates/splash.vue'

export default [
    {path: '/gooey/browse/:tab', component: Browse},
    {path: '/gooey/structure/:tab', component: Structure},
    {path: '/gooey/query/:tab', component: Query},
    {path: '/gooey/', component: Splash}

]

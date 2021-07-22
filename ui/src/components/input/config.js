import http from './http-config.vue'
import sql from './sql.vue'
import static_ from './static-config.vue'
import * as type from './type'
export default {
    [type.Http]: http,
    [type.Sql]: sql,
    [type.Static]: static_,
}
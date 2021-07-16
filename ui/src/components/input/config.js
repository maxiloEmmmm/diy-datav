import http from './http-config.vue'
import sql from './sql-config.vue'
import * as type from './type'
export default {
    [type.Http]: http,
    [type.Sql]: sql
}
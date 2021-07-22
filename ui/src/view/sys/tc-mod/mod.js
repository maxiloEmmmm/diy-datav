import * as inputType from '@/components/input/type'
import http from './http-type-config.vue'
import sql from './sql-type-config.vue'
import static_ from './static-type-config.vue'

export default {
    [inputType.Http]: http,
    [inputType.Sql]: sql,
    [inputType.Static]: static_,
}

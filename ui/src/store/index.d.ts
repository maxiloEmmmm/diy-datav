import { ComponentCustomProperties } from 'vue'
import storeType from './type'
declare module '@vue/runtime-core' {
    interface ComponentCustomProperties {
        $storeType: storeType,
    }
}
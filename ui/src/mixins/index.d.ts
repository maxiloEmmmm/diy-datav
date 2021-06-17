import { ComponentCustomProperties } from 'vue'

declare module '@vue/runtime-core' {
    interface ComponentCustomProperties {
        // 初始化focus
        _initFocus: Function,
        // 主动执行focus
        _doFocus: Function,
        // 全局混合数据
        app_mixin: {
            // focus模块
            focus: {
                // 当前是否focus
                in: boolean,
            }
        }
    }
}
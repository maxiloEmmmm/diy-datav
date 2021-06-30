import { ComponentCustomProperties } from 'vue'

interface Help {
    // 唯一标识, typ内唯一
    key: string
    // 图标
    component?()
    // 点击回调函数
    cb(),
    disable?: boolean
}

declare module '@vue/runtime-core' {
    interface ComponentCustomProperties {
        // 初始化focus
        mixinInitFocus(),
        // 主动执行focus
        mixinDoFocus(),
        // 添加帮助
        mixinAddHelp(typ: string, helps: [Help]),
        // 设置帮助
        mixinSetHelp(typ: string, helps: [Help]),
        // 禁用帮助
        mixinDisableHelp(typ: string, key: string),
        // 启用帮助
        mixinActiveHelp(typ: string, key: string),
        // 移除帮助
        mixinRemoveHelp(typ: string, key: string),
        // 显示配置
        mixinConfigShow(),
        // 隐藏配置
        mixinConfigHidden(),
        // 设置配置块标识
        mixinSetConfigKey(key: string),
        // 设置配置的类型
        mixinSetConfigType(typ: string),
        // 设置配置的类型和配置
        mixinSetConfigTypeAndConfig(typ: string, config: string),
        // 设置配置的配置
        mixinSetConfigConfig(config: string),
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
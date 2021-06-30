export const AntVConfig = function() {
    return {
        // TODO: 支持多类型
        type: '',
        color: '',
        coordinate: {
            type: '',
            transpose: false
        },
        scale: {
            x: {
                field: "",
                alias: "",
                format: {
                    prefix: '',
                    suffix: ''
                }
            },
            y: {
                field: "",
                alias: "",
                format: {
                    prefix: '',
                    suffix: ''
                }
            }
        },
        cats: []
    }
}
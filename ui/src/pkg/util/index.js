const util = {
    uuid,
    debounce,
    throttle
}

export default {
    ...util,
    install(app) {
        app.config.globalProperties.$util = util
    }
}

function uuid() {
    function S4() {
        return (((1 + Math.random()) * 0x10000) | 0).toString(16).substring(1);
    }
    return (S4() + S4() + "-" + S4() + "-" + S4() + "-" + S4() + "-" + S4() + S4() + S4());
}

function debounce(cb, timeout) {
    let handler = null
    return function() {
        if(handler == null) {
            handler = setTimeout(() => {
                handler = null
            }, timeout)
            cb(...arguments)
        }
    }
}

function throttle(cb, timeout) {
    let handler = null
    return function() {
        if(handler == null) {
            handler = setTimeout(() => {
                handler = null
                cb(...arguments)
            }, timeout)
        }
    }
}


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

const has = function (
    obj,
    path,
    returnValue = false,
    endValue = undefined,
    failBack = null) {
    let pathInfo = path.split('.')
    let x = []
    let tmp = {}
    pathInfo.forEach(v => {
        if (/\[/.test(v)) {
            //find [a]b[c] | b[a]c | [a][b]c => a.b.c
            let arrayPathInfo = v.match(/(\[([^\[\]]+?)\]|[^\[\]]+)+?/g)
            if (arrayPathInfo !== null) {
                arrayPathInfo.forEach(q => {
                    x.push(q.replace('[', '').replace(']', ''))
                })
            }
        } else {
            x.push(v)
        }
    })
    let x_len = x.length
    for (let i = 0; i < x_len; i++) {
        let v = x[i]
        if (!['[object Array]', '[object Object]'].includes(Object.prototype.toString.call(obj))
            || !(v in obj)) {
            if (failBack !== null) {
                failBack(obj, x.slice(i))
            }

            if(!returnValue) {
                return false
            }
            return returnValue ? endValue : false
        }
        tmp = obj
        obj = obj[v]
    }
    if (failBack !== null) {
        failBack(tmp, [x[x_len - 1]])
    }
    return returnValue ? obj : true
}

const getType = v => {
    let str = Object.prototype.toString.call(v);
    return str.slice(8, str.length-1);
}

const isDefine = function (obj, path) {
    let value = has(obj, path, true)
    return value !== undefined
}

const typeEQ = function (o, o2) {
    return getType(o) === getType(o2)
}

const isObject = o => getType(o) === 'Object'
const isArray = o => getType(o) === 'Array'
const isString = o => getType(o) === 'String'
const get = function (obj, path, d = null) {
    let value = has(obj, path, true)
    return value === undefined ? d : value
}

const set = function (obj, path, d) {
    has(obj, path, false, false, function (obj, pathInfo) {
        let p_len = pathInfo.length
        for (let i = 0; i < p_len; i++) {
            let v = pathInfo[i]
            if (i + 1 < p_len) {
                let tmp = {}
                obj = obj[v] = tmp
            } else {
                obj[v] = d
            }
        }
    })
}


const util = {
    uuid,
    debounce,
    throttle,
    has,
    set,
    get,
    isString,
    isArray,
    isObject
}

export default {
    ...util,
    install(app) {
        app.config.globalProperties.$util = util
    }
}


function uuid() {
    function S4() {
        return (((1 + Math.random()) * 0x10000) | 0).toString(16).substring(1);
    }
    return (S4() + S4() + "-" + S4() + "-" + S4() + "-" + S4() + "-" + S4() + S4() + S4());
}

function debounce(cb, timeout) {
    let handler = null
    return function() {
        if(handler) {
            clearTimeout(handler)
            handler = null
        }

        handler = setTimeout(() => {
            handler = null
            cb.call(this, ...arguments)
        }, timeout)
    }
}

function throttle(cb, timeout) {
    let can = true
    let handler = null
    return function() {
        if(can) {
            can = false
            handler = setTimeout(() => {
                can = true
            }, timeout)
            cb.call(this, ...arguments)
        }
    }
}

function deepClone(obj) {
    let o = obj
    if (util.isObject(obj)) {
        o = {}
        Object.keys(obj).forEach(k => {
            o[k] = deepClone(obj[k])
        })
    }else if(util.isArray(obj)) {
        o = []
        obj.forEach(val => {
            o.push(deepClone(val))
        })
    }

    return o
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
const isNumber = o => getType(o) === 'Number'
const isBoolean = o => getType(o) === 'Boolean'
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

const merge = function(dst, src) {
    dst = deepClone(dst)
    if(!typeEQ(dst, src)) {
        return dst
    }

    if(util.isArray(dst)) {
        return src
    }

    if(!util.isObject(dst)) {
        return dst
    }

    Object.keys(src).forEach(k => {
        let val = src[k]
        let dstVal = get(dst, k)
        if(!has(dst, k) || !typeEQ(val, dstVal)) {
            dst[k] = val
        }else {
            if(isArray(val)) {
                dst[k] = val
            }else if(isObject(dstVal)) {
                dst[k] = merge(dstVal, val)
            }else {
                dst[k] = val
            }
        }
    })

    return dst
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
    isObject,
    isBoolean,
    isNumber,
    deepClone,
    merge
}

export default {
    ...util,
    install(app) {
        app.config.globalProperties.$util = util
    }
}
import util from 'pkg/util'

export const PositionType = function() {
    return {
        left: 0,
        top: 0,
        width: 10,
        height: 10
    }
}

export const PaddingParse = (t) => {
    let cfg = Padding()
    cfg.top = PaddingFilter.top(t.top)
    cfg.left = PaddingFilter.left(t.left)
    cfg.right = PaddingFilter.right(t.right)
    cfg.bottom = PaddingFilter.bottom(t.bottom)
    return cfg
}

export const Padding = () => {
    return {
        top: PaddingDefault.top(),
        bottom: PaddingDefault.bottom(),
        left: PaddingDefault.left(),
        right: PaddingDefault.right(),
    }
}

export const PaddingDefault = {
    top() {
        return 0
    },
    bottom() {
        return 0
    },
    left() {
        return 0
    },
    right() {
        return 0
    },
}

export const PaddingFilter = {
    top(t) {
        return !util.isNumber(t) || t < 0 ? PaddingDefault.top() : t
    },
    bottom(t) {
        return !util.isNumber(t) || t < 0 ? PaddingDefault.bottom() : t
    },
    left(t) {
        return !util.isNumber(t) || t < 0 ? PaddingDefault.left() : t
    },
    right(t) {
        return !util.isNumber(t) || t < 0 ? PaddingDefault.right() : t
    },
}
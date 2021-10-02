import util from 'pkg/util'

export const PositionType = function() {
    return {
        left: PositionTypeDefault.left(),
        top: PositionTypeDefault.top(),
        width: PositionTypeDefault.width(),
        height: PositionTypeDefault.height()
    }
}

export const PositionTypeParse = t => {
    let pt = PositionType()
    pt.top = PositionTypeFilter.top(t.top)
    pt.left = PositionTypeFilter.top(t.left)
    pt.width = PositionTypeFilter.top(t.width)
    pt.height = PositionTypeFilter.top(t.height)
    return pt
}

export const PositionTypeFilter =  {
    top(t) {
        return util.isNumber(t) && t >= 0 ? t : PositionTypeDefault.top()
    },
    left(t) {
        return util.isNumber(t) && t >= 0 ? t : PositionTypeDefault.left()
    },
    width(t) {
        return util.isNumber(t) && t >= 0 ? t : PositionTypeDefault.width()
    },
    height(t) {
        return util.isNumber(t) && t >= 0 ? t : PositionTypeDefault.height()
    }
}

export const PositionTypeDefault = {
    top() {
        return 0
    },
    left() {
        return 0
    },
    width() {
        return 10
    },
    height() {
        return 10
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
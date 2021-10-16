import {ViewBLockTypeCommon} from 'type'

export default {
    inject: ['pointerEventsNone', 'blockKey'],
    props: {
        common: {
            type: Object,
            default: ViewBLockTypeCommon
        },
    }
}
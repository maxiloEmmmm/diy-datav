import * as type from './type'
import bgMockUrl from '@/assets/bg_design.png'

export default function() {
    return {
        [type.ViewStore]() {
            mockUtil.mock('view', 'put', function(request) {
                return {code: 'ok', msg: '', body: request.body}
            })
        },
        [type.ViewUploadBG](file) {
            return http.post('view/bg/upload')
        }
    }
}
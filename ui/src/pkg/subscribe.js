export default function() {
    const dd = {}
    this.addHandle = (resource, handle) => {
        if(!dd[resource]) {
            dd[resource] = []
        }
        dd[resource].push(handle)
    }
    this.dispatch = (resource, payload) => {
        (dd[resource] || []).forEach(h => {
            h(payload)
        })
    }
}
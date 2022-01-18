export default {
    state: {
        isCollapse: true,
        message: "ojsiuj"
    },
    mutations: {
        collapseMenu (state) {
            state.isCollapse = !state.isCollapse
        }
    }
}
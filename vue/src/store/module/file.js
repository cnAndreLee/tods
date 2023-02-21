import fileService from "@/service/fileService"

const fileModule = {
    namespaced: true,
    state: {
        files:[],
        selectedFile:{suffix:null},          // player/index需要suffix，否则报错，待优化
    },
    getters:{
        // get2Files(state) { return state.files }
    },

    mutations: {
        SET_FILES(state,files) {
            state.files = files
        },
        SET_SelectedFile(state, pyload) {
            state.selectedFile = pyload
        }
    },

    actions: {
        getCategories(context) {
            return new Promise((resolve, reject) => {
                fileService.getCategory().then( (res) => {
                    if ( res.data.status == 2000 ) {
                        resolve(res)
                    } else {
                        reject(res)
                    }
                }).catch((res) => {
                    reject(res)
                })
            })
        },
        deleteFile(context, payload) {
            return new Promise((resolve, reject) => {
                fileService.deleteFile(payload.id).then((res)=>{
                    if ( res.data.status == 2000 ) {
                        resolve(res)
                    } else {
                        reject(res)
                    }
                }).catch((res) => {
                    reject(res)
                })
            })
        },
        getFiles(context, payload) {
            return new Promise((resolve, reject) => {
                fileService.getFile(payload.id).then((res)=>{
                    if ( res.data.status == 2000 ) {
                        context.commit('SET_FILES', res.data.data.files)
                        resolve(res)
                    } else {
                        reject(res)
                    }
                }).catch((res) => {
                    reject(res)
                })
            })
        },
        postCategory(context, payload) {
            return new Promise((resolve, reject) => {
                fileService.postCategory(payload).then((res)=>{
                    if ( res.data.status == 2000 ) {
                        resolve(res)
                    } else {
                        reject(res)
                    }
                }).catch((res) => {
                    reject(res)
                })
            })
        },
        deleteCategory(context, payload) {
            return new Promise((resolve, reject) => {
                fileService.deleteCategory(payload).then((res)=>{
                    if ( res.data.status == 2000 ) {
                        resolve(res)
                    } else {
                        reject(res)
                    }
                }).catch((res) => {
                    reject(res)
                })
            })
        },
        renameCategory(context, payload) {
            return new Promise((resolve, reject) => {
                fileService.renameCategory(payload).then((res)=>{
                    if ( res.data.status == 2000 ) {
                        resolve(res)
                    } else {
                        reject(res)
                    }
                }).catch((res) => {
                    reject(res)
                })
            })
        },
        changeCategoryPermissions(context, payload) {
            return new Promise((resolve, reject) => {
                fileService.changeCategoryPermissions(payload).then((res)=>{
                    if ( res.data.status == 2000 ) {
                        resolve(res)
                    } else {
                        reject(res)
                    }
                }).catch((res) => {
                    reject(res)
                })
            })
        },
    },
}

export default fileModule;
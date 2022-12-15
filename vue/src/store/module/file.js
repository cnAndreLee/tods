import fileService from "@/service/fileService"

const fileModule = {
    namespaced: true,
    state: {
        files:[],
        selectdSecondaryCategory:null,
        selectedFile:{suffix:null},
    },
    getters:{
        // get2Files(state) { return state.files }
    },

    mutations: {
        SET_FILES(state,files) {
            state.files = files
        },
        SET_selectdSecondaryCategory(state,pyload) {
            state.selectdSecondaryCategory = pyload.id
        },
        SET_SelectedFile(state, pyload) {
            state.selectedFile = pyload
        }
    },

    actions: {
        getFiles(context) {
            return new Promise((resolve, reject) => {
                fileService.getFile().then((res)=>{
                    context.commit('SET_FILES', res.data.data.files)
                })
            })
        }
    },
}

export default fileModule;
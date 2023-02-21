import storageService from "@/service/storageService";
import userService from "@/service/userService"

const userModule = {
    namespaced: true,
    state: {
        token: storageService.get(storageService.USER_TOKEN),
        userInfo: JSON.parse(storageService.get(storageService.USER_INFO)),

    },

    mutations: {
        SET_TOKEN(state, token) {
            // 更新本地缓存
            storageService.set(storageService.USER_TOKEN, token)
            // 更新state
            state.token = token;
        },
        SET_USERINFO(state, userInfo) {
            // 更新本地缓存
            // if (userInfo === "") {
            //     storageService.set(storageService.USER_INFO, "");
            // } else {
                storageService.set(storageService.USER_INFO, JSON.stringify(userInfo))
            // }
            
            state.userInfo = userInfo
        }
    },

    actions: {
        login(context, { Account, Key }) {
            return new Promise((resolve, reject) => {
                userService.login({ Account, Key }).then( (res)=> {
                    if ( res.data.status == 2000 ) {
                        // save token
                        context.commit('SET_TOKEN', res.data.data.token)
                        resolve(res)
                    } else {
                        reject(res)
                    }
                })
                .catch((err) => {
                    reject(err)
                })
            })
        },
        register(context, payload) {
            return new Promise((resolve, reject ) => {
                userService.register(payload).then( (res) => {
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
        logout({commit}) {
            // clear 
            localStorage.clear()
        },
        info(context) {
            return new Promise((resolve, reject) => {
                userService.info().then( (res) => {
                    if ( res.data.status == 2000 ) {
                        context.commit('SET_USERINFO', res.data.data.user)
                        resolve(res)
                    } else {
                        localStorage.clear()
                        reject(res)
                    }
                }).catch((err) => {
                    localStorage.clear()
                    reject(err)
                })
            })
        },
        getAllUsers(context) {
            return new Promise((resolve, reject) => {
                userService.getAllUsers().then( (res) => {
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
        getAllSchools(context) {
            return new Promise((resolve, reject) => {
                userService.getAllSchools().then( (res) => {
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
        deleteUser(context, payload) {
            return new Promise((resolve, reject ) => {
                userService.deleteUser(payload).then( (res) => {
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

export default userModule;
import storageService from "@/service/storageService";
import userService from "@/service/userService"

const userModule = {
    namespaced: true,
    state: {
        token: storageService.get(storageService.USER_TOKEN),
        userInfo: JSON.parse(storageService.get(storageService.USER_INFO))
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
            if (userInfo === "") {
                storageService.set(storageService.USER_INFO, "");
            } else {
                storageService.set(storageService.USER_INFO, JSON.stringify(userInfo))
            }
            // 更新state
            state.userInfo = userInfo
        }
    },

    actions: {
        login(context, { Account, Key }) {
            return new Promise((resolve, reject) => {
                userService.login({ Account, Key }).then( (res)=> {
                    if (res.data.status == 0) {
                        // save token
                        context.commit('SET_TOKEN', res.data.data.token)
                        resolve(res)
                    } else {
                        reject(res)
                    }
                    
                    // return userService.info() 
                })
                // .then( response => {
                //     console.log("second then")
                //     // save userifo
                //     // context.commit('SET_USERINFO', response.data.data.user)
                    
                //     // resolve(response)
                // })
                .catch((err) => {
                    console.log("catch err")
                    reject(err)
                })
            })
        },
        register(context, { Account, Key }) {
            return new Promise((resolve, reject ) => {
                userService.register({ Account, Key }).then( (res) => {
                    console.log("singup success")
                }).then( response => {
                    resolve(response)
                }).catch((err) => {
                    reject(err)
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
                    context.commit('SET_USERINFO', res.data.data.user)
                }).catch((err) => {
                    localStorage.clear()
                    reject(err)
                })
            })
        }
    },
}

export default userModule;
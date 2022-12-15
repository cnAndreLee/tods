import { createStore } from 'vuex'
import userModule from './module/user'
import fileModule from './module/file'

const configModule = {
  namespaced: true,
  state: () => ({

  }),
  mutations:{},
  actions:{},
  getters:{}
}


const store = createStore({

  // 严格模式，确保发布环境时strict模式关闭，以保证性能
  // strict: process.env.NODE_ENV !== 'production' ,
  
  modules: {
    userModule,
    fileModule,
    configModule,
  }
})

export default store
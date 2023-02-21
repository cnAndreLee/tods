import request from '../api/request'

const getCategory = () => {
    return request.get('api/v1/categories/')
}

// 根据分类id获取文件列表
const getFile = (id) => {
    return request.get('api/v1/file' + "?id=" + id)
}


// 根据文件id删除
const deleteFile = (id) => {
    return request.delete('api/v1/file' + "?id=" + id)
}

// 根据id删除分类或合集
const deleteCategory =(payload) => {
    return request.delete('api/v1/categories/', {data:payload})
}

const postCategory = (payload) => {
    return request.post('api/v1/categories/', payload)
}

// 重命名分类
const renameCategory = (payload) => {
    return request.put('api/v1/categories/', payload)
}
// 修改分类权限
const changeCategoryPermissions = (payload) => {
    return request.put('/api/v1/categories/permission', payload)
}
export default {
    getCategory,
    deleteFile,
    getFile,
    deleteCategory,
    postCategory,
    renameCategory,
    changeCategoryPermissions,
}
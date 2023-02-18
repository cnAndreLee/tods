import request from '../api/request'

const getCategory = () => {
    return request.get('api/v1/categories/')
}

// 根据文件id删除
const deleteFile = (id) => {
    return request.delete('api/v1/file' + "?id=" + id)
}


// 根据分类id获取文件列表
const getFile = (id) => {
    return request.get('api/v1/file' + "?id=" + id)
}


export default {
    getCategory,
    deleteFile,
    getFile
}
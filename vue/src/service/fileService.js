import request from '../api/request'

const getCategory = () => {
    return request.get('api/v1/category/')
}

const getSecondaryCategory = () => {
    return request.get('api/v1/secondarycategory/')
}

const getFile = () => {
    return request.get('api/v1/file')
}

export default {
    getCategory,
    getSecondaryCategory,
    getFile
}
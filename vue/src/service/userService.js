import request from '../api/request'

//user register
const register = (payload) => {
    return request.post('api/v1/user/register', payload)
}

const login = ({Account, Key}) => {
    return request.post('api/v1/user/login', {Account, Key})
}

const info = () =>{
    return request.get('api/v1/user/info');
}

// admin get all users
const getAllUsers = () => {
    return request.get('api/v1/user/users');
}
const getAllSchools = () => {
    return request.get('api/v1/user/school');
}

const deleteUser = (payload) => {
    return request.delete('api/v1/user' + "?account=" + payload)
}

export default {
    register,
    login,
    info,
    getAllUsers,
    deleteUser,
    getAllSchools,
}
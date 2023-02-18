import request from '../api/request'

//user register
const register = ({Account, Key}) => {
    return request.post('api/v1/user/register', {Account, Key})
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

export default {
    register,
    login,
    info,
    getAllUsers,
}
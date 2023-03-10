// 本地缓存服务

const PREFIX = 'tods_';

// user 模块
const USER_PREFIX = `${PREFIX}user_`;
const USER_TOKEN = `${USER_PREFIX}token`;
const USER_INFO = `${USER_PREFIX}info`;


const set = (key, data) => {
    localStorage.setItem(key, data);
}

const get = (key) => localStorage.getItem(key)

const remove =(key) => localStorage.removeItem(key)

export default {
    set,
    get,
    remove,
    USER_TOKEN,
    USER_INFO,
};
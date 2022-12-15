import requests from './request'

//loginApi
export const reqToken = (data) => requests({url:"/api/token", method:'post'});


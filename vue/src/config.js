const config = {
    // 生产环境下交给nginx做反向代理，对外只暴露一个origin
    backend: import.meta.env.MODE == "development" ? import.meta.env.VITE_TODS_SERVER : window.location.origin,
    kkfile : import.meta.env.MODE == "development" ? import.meta.env.VITE_KK_SERVER : window.location.origin,

    filespath: "/api/files/",
    supportFileType :[
        "audio/mpeg", //mp3
        "video/mp4", 
        "application/msword", 
        "application/vnd.openxmlformats-officedocument.wordprocessingml.document", 
        "application/vnd.ms-powerpoint",
        "application/vnd.openxmlformats-officedocument.presentationml.presentation", 
        "application/vnd.ms-excel", 
        "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
        "application/pdf",
    ],
    officeFileType :['docx', 'doc', 'pptx', 'ppt', 'xlsx', 'xls', 'pdf', 'mp3']
}


export default config

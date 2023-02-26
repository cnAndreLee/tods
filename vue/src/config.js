const config = {
    backend: "http://10.0.0.10:8000",
    kkfile: "http://10.0.0.10:8012/onlinePreview?url=",

    filespath: "/api/files/",
    supportFileType :[
        "video/mp4", 
        "application/msword", 
        "application/vnd.openxmlformats-officedocument.wordprocessingml.document", 
        "application/vnd.ms-powerpoint",
        "application/vnd.openxmlformats-officedocument.presentationml.presentation", 
        "application/vnd.ms-excel", 
        "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
        "application/pdf",
    ]
}



export default config

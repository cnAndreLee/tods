const config = {
    backend: import.meta.env.VITE_TODS_SERVER,
    kkfile: import.meta.env.VITE_KK_SERVER,

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
    ],
    officeFileType :['docx', 'doc', 'pptx', 'ppt', 'xlsx', 'xls', 'pdf']
}

export default config

import instance from "@/request/request";


export function Login(data: any){
    return instance({
        url:"api/v1/login",
        method:"post",
        data
    })
}

export  function getAdminInfo(){
    return instance({
        url:"api/v1/admin/info",
        method:"get"
    })
}

export  function  ChangeInfo(data :any){
    return instance({
        url:"api/v1/admin/change/info",
        method:"post",
        data
    })
}

export  function  ChangePassword(data:any){
    return instance({
        url:"api/v1/admin/change/password",
        method:"post",
        data
    })
}

export function Logout(){
    return instance({
        url:"/api/v1/admin/logout",
        method:"post",
    })
}

export  function  getAdminList(data:any){
     return instance({
         url:"api/v1/admin/index",
         method:"get",
         params:data
     })
}
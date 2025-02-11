import instance from "@/request/request";
import {UnwrapRef} from "vue";
import {RouteParamValue} from "vue-router";


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

export  function delAdmin(id:number){
    return instance({
        "url":"api/v1/admin/delete",
        "method":"delete",
        data: { id },
    })
}

export  function getAdminById(id: UnwrapRef<string | RouteParamValue[]>){
    return instance({
        "url":`api/v1/admin/detail`,
        "method":"get",
        params:{"id":id}
    })
}
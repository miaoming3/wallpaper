import instance from "@/request/request";


export function Login(data: any){
    return instance({
        url:"api/v1/login",
        method:"post",
        data
    })
}
import axios from "axios";
import {ElMessage} from "element-plus";
import router from "@/router";
import {routes} from "vue-router/vue-router-auto-routes";

const  instance =axios.create({
    baseURL: "http://localhost:8078/",
    timeout: 20000
})

instance.interceptors.request.use(
    config=>{
        if(localStorage.getItem("token")){
            config.headers.Authorization=localStorage.getItem("token")
            config.headers["uid"]=localStorage.getItem("uid")
        }
        return config
    },error=>{
        return Promise.reject(error)
    }
)

instance.interceptors.response.use(
    response=>{
        return  response.data

    },error=>{
       if (error.code ==="ERR_NETWORK"){
           ElMessage.error("网络错误")
           return 
       }
        if (error.response.data.code ==401 && error.status==401){
            ElMessage.error("授权过期,3s后请重新登录")
            setTimeout(function (){
                router.replace({"path":"/"})
            },3000)
            return
        }
        return Promise.reject(error)
    }
)

export default  instance